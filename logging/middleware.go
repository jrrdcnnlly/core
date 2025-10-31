package logging

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/jrrdcnnlly/core/id"
)

func Middleware(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			// Create a new logger from the parent logger.
			logger := logger.With(
				slog.Group(
					"req",
					slog.Uint64("id", id.Sequential.Next()),
					slog.String("method", r.Method),
					slog.String("path", r.URL.Path),
				),
			)

			logger.Debug("incoming request")

			rw := newResponseWriter(w)

			next.ServeHTTP(rw, withLogger(r, logger))

			elapsed := time.Since(start)
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
