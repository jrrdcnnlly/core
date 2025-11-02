package auth

import (
	"log/slog"
	"net/http"

	"github.com/jrrdcnnlly/core/logging"
	"github.com/jrrdcnnlly/core/sessions"
)

func MSALMiddleware(client *MSALClient) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get logger from request context.
			logger := logging.FromContextOrDefault(r.Context())

			// Get session from request context.
			session, err := sessions.FromContext(r.Context())
			if err != nil {
				logger.Error(
					"failed to retrieve session",
					slog.Any("err", err),
				)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// Attempt to refresh access token.
			err = client.refreshToken(r.Context(), session)
			if err == nil {
				logger.Debug("access token refreshed")
				next.ServeHTTP(w, r)
				return
			}

			logger.Debug("session is not authenticated")

			// Redirect to microsoft login page.
			authCodeURL, err := client.authCodeURL(r.Context())
			if err != nil {
				logger.Error(
					"failed to generae auth code URL",
					slog.Any("err", err),
				)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, authCodeURL, http.StatusTemporaryRedirect)
		})
	}
}
