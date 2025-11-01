package logging

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
	"time"
)

// TextHandler configuration.
type textHandlerConfig struct {
	w     io.Writer
	level slog.Level
}

// TextHandler option.
type TextHandlerOption func(cfg *textHandlerConfig)

// TextHandler implements slog.Handler.
type TextHandler struct {
	level  slog.Level
	attrs  []slog.Attr
	group  string
	writer io.Writer
}

// Create a hander with the specified level.
func WithLevel(level slog.Level) TextHandlerOption {
	return func(cfg *textHandlerConfig) {
		cfg.level = level
	}
}

// Create a handler with the specified writer.
func WithWriter(w io.Writer) TextHandlerOption {
	return func(cfg *textHandlerConfig) {
		cfg.w = w
	}
}

// Create a new TextHandler.
func NewTextHandler(options ...TextHandlerOption) *TextHandler {
	// Init default config.
	cfg := &textHandlerConfig{
		w:     os.Stdout,
		level: slog.LevelInfo,
	}
	// Apply options to config.
	for _, option := range options {
		option(cfg)
	}
	// Create new TextHandler from config.
	return &TextHandler{
		level:  cfg.level,
		attrs:  make([]slog.Attr, 0),
		writer: cfg.w,
	}
}

// Check if the handler should handle the specified level.
func (h *TextHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= h.level
}

// Create a new handler with the specified attributes.
func (h *TextHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &TextHandler{
		level:  h.level,
		attrs:  append(h.attrs, attrs...),
		group:  h.group,
		writer: h.writer,
	}
}

// Create a new handler with the specified group name.
func (h *TextHandler) WithGroup(name string) slog.Handler {
	return &TextHandler{
		level:  h.level,
		attrs:  h.attrs,
		group:  name,
		writer: h.writer,
	}
}

// Handle a log record.
func (h *TextHandler) Handle(ctx context.Context, r slog.Record) error {
	text, err := h.format(r)
	if err != nil {
		return err
	}
	h.writer.Write([]byte(text))
	return nil
}

// Format a log record.
func (h *TextHandler) format(r slog.Record) (string, error) {
	// Format primary log record properties,
	// Time, Level & Message.
	text := fmt.Sprintf(
		"%s %s %s\n",
		h.formatTime(r.Time),
		h.formatLevel(r.Level),
		r.Message,
	)

	// Attributes start indented.
	depth := 1

	// Format handler attributes.
	for _, attr := range h.attrs {
		text += h.formatAttr(attr, depth)
	}

	// If handler has a group name specified,
	// format the group name and indent the depth.
	if len(h.group) > 0 {
		text += h.formatGroup(h.group, depth)
		depth++
	}

	// Format log record attributes.
	r.Attrs(func(attr slog.Attr) bool {
		text += h.formatAttr(attr, depth)
		return true
	})
	return text, nil
}

// Format log record time.
func (h *TextHandler) formatTime(t time.Time) string {
	return fmt.Sprintf("\033[%dm[%s]\033[0m", ANSIForegroundBlue, t.Format(time.RFC3339))
}

// Format log record level.
func (h *TextHandler) formatLevel(level slog.Level) string {
	switch level {
	case slog.LevelInfo:
		return fmt.Sprintf("\033[%dm[%5s]\033[0m", ANSIForegroundGreen, level)
	case slog.LevelWarn:
		return fmt.Sprintf("\033[%dm[%5s]\033[0m", ANSIForegroundMagenta, level)
	case slog.LevelError:
		return fmt.Sprintf("\033[%dm[%5s]\033[0m", ANSIForegroundRed, level)
	default:
		return fmt.Sprintf("\033[%dm[%5s]\033[0m", ANSIForegroundCyan, level)
	}
}

// Format handler group name.
func (h *TextHandler) formatGroup(name string, depth int) string {
	return fmt.Sprintf(
		"%s\033[%dm%s:\033[0m\n",
		strings.Repeat("  ", depth),
		ANSIForegroundGray,
		name,
	)
}

// Format log record attribute.
func (h *TextHandler) formatAttr(attr slog.Attr, depth int) string {
	if attr.Value.Kind() == slog.KindGroup {
		text := fmt.Sprintf(
			"%s\033[%dm%s:\033[0m\n",
			strings.Repeat("  ", depth),
			ANSIForegroundGray,
			attr.Key,
		)
		for _, subAttr := range attr.Value.Group() {
			text += h.formatAttr(subAttr, depth+1)
		}
		return text
	} else {
		return fmt.Sprintf(
			"%s\033[%dm%s: %v\033[0m\n",
			strings.Repeat("  ", depth),
			ANSIForegroundGray,
			attr.Key,
			attr.Value.Any(),
		)
	}
}
