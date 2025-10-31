package logging

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
)

// Used as a unique context key.
type loggerKey struct{}

// Retrieve a logger from a context.
// If the context does not contain a logger an error is returned.
func FromContext(ctx context.Context) (*slog.Logger, error) {
	value := ctx.Value(loggerKey{})
	if value == nil {
		return nil, errors.New("no logger in context")
	}
	if logger, ok := value.(*slog.Logger); ok {
		return logger, nil
	}
	return nil, errors.New("no logger in context")
}

// Create a new request with a new context that contains the specified logger.
func withLogger(r *http.Request, logger *slog.Logger) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), loggerKey{}, logger))
}
