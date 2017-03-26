package middleware

import (
	"net/http"

	"github.com/adamjedlicka/webapp/src/shared/session"
)

// MustLogin check if any user is logged in.
// If not shows error page and cancels entry to the requested page
func MustLogin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !session.IsLogin(r) {
			http.Error(w, "Bad authentification!", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}
