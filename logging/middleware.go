package logging

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/jrrdcnnlly/core/id"
)

// Middleware configuration.
type middlewareConfig struct {
	logger *slog.Logger
}

// Middleware option.
type MiddlewareOption func(*middlewareConfig)

// Create a middleware from a specific logger.
// Defaults to the default logger.
func WithLogger(logger *slog.Logger) MiddlewareOption {
	return func(cfg *middlewareConfig) {
		cfg.logger = logger
	}
}

// Create a new request logging middleware.
func Middleware(options ...MiddlewareOption) func(http.Handler) http.Handler {
	// Init default config.
	cfg := &middlewareConfig{
		logger: slog.Default(),
	}
	// Apply options to config.
	for _, option := range options {
		option(cfg)
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Record time of middleware entry to calculate processing duration.
			start := time.Now()
			// Create a new logger from the parent logger.
			logger := cfg.logger.With(
				slog.Group(
					"req",
					slog.Uint64("id", id.Sequential.Next()),
					slog.String("method", r.Method),
					slog.String("path", r.URL.Path),
				),
			)
			// Log start of request
			logger.Debug("incoming request")
			// ResponseWriter wraps an http.ResponseWriter to capture the return status code.
			rw := newResponseWriter(w)

			// Pass the rquest scoped logger to the next handler in the request context.
			next.ServeHTTP(rw, Request(r, logger))

			// Calculate total request processing time.
			elapsed := time.Since(start)
			// Log end of request.
			logger.Info(
				"request completed",
				slog.Duration("elapsed", elapsed),
				slog.Group(
					"res",
					slog.Int("status", rw.statusCode),
				),
			)
		})
	}
}
