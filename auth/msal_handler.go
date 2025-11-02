package auth

import (
	"log/slog"
	"net/http"

	"github.com/jrrdcnnlly/core/logging"
	"github.com/jrrdcnnlly/core/sessions"
)

func MSALHandler(client *MSALClient) http.Handler {
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

		// Attempt to acquire access token.
		code := r.URL.Query().Get("code")
		err = client.acquireToken(r.Context(), code, session)
		if err != nil {
			logger.Error(
				"failed to retrieve token",
				slog.Any("err", err),
			)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	})
}
