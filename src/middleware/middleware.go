package middleware

import (
	"net/http"

	"github.com/adamjedlicka/webapp/src/models"
)

func MustLogin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("auth")
		if err != nil {
			http.Error(w, "Bad authentification!", http.StatusBadRequest)
			return
		}

		s := models.NewSession()
		err = s.FindByName(c.Value)
		if err != nil {
			http.Error(w, "Bad authentification!", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}
