package logging

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"strings"
	"time"
)

type TextHandler struct {
	level  slog.Level
	attrs  []slog.Attr
	group  string
	writer io.Writer
}

func NewTextHandler(w io.Writer, level slog.Level) *TextHandler {
	return &TextHandler{
		level:  level,
		attrs:  make([]slog.Attr, 0),
		writer: w,
	}
}

func (h *TextHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= h.level
}

func (h *TextHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &TextHandler{
		level:  h.level,
		attrs:  append(h.attrs, attrs...),
		group:  h.group,
		writer: h.writer,
	}
}

func (h *TextHandler) WithGroup(name string) slog.Handler {
	return &TextHandler{
		level:  h.level,
		attrs:  h.attrs,
		group:  name,
		writer: h.writer,
	}
}

func (h *TextHandler) Handle(ctx context.Context, r slog.Record) error {
	text, err := h.format(ctx, r)
	if err != nil {
		return err
	}
	h.writer.Write([]byte(text))
	return nil
}

func (h *TextHandler) format(ctx context.Context, r slog.Record) (string, error) {
	text := fmt.Sprintf(
		"%s %s %s\n",
		h.formatTime(r.Time),
		h.formatLevel(r.Level),
		r.Message,
	)

	for _, attr := range h.attrs {
		text += h.formatAttr(attr, 1)
	}

	r.Attrs(func(attr slog.Attr) bool {
		text += h.formatAttr(attr, 1)
		return true
	})
	return text, nil
}

func (h *TextHandler) formatTime(t time.Time) string {
	return fmt.Sprintf("\033[%dm[%s]\033[0m", ANSIForegroundBlue, t.Format(time.RFC3339))
}

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
