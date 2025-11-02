package sessions

import (
	"log/slog"
	"net/http"

	"github.com/jrrdcnnlly/core/logging"
)

// Session cookie name.
const sessionCookie string = "session_id"

// Create a session handling middleware backed by the specified store.
func Middleware(store SessionStore) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get logger from request context.
			logger := logging.FromContextOrDefault(r.Context())

			// Get or create session.
			session, err := store.FromRequest(r)
			if err != nil {
				logger.Error(
					"failed to load session",
					slog.Any("err", err),
				)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// Create new logger with session information.
			logger = logger.With(
				slog.Group(
					"session",
					slog.String("id", session.id),
				),
			)

			logger.Debug("loaded session")

			// Pass the new logger to the next handler in the request context.
			next.ServeHTTP(w, logging.Request(r, logger))

			// Save session back to store.
			err = store.Update(session)
			if err != nil {
				logger.Error(
					"failed to save session",
					slog.Any("err", err),
				)
			}
		})
	}
}
