package controller

import (
	"log"
	"net/http"

	"github.com/adamjedlicka/webapp/src/model"
	"github.com/adamjedlicka/webapp/src/shared/session"
	"github.com/gorilla/sessions"
)

func LoginPOST(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	username := r.FormValue("username")
	password := r.FormValue("password")

	log.Println("Attemt to login :: username: ", username, ", password: ", password)

	u := model.NewUser()
	err := u.FindByUsername(username)
	if err != nil || password != u.Password() {
		http.Error(w, "Bad username or password!", http.StatusUnauthorized)
		return
	}

	session, err := session.SessionStore.Get(r, session.SessionAuth)
	if err != nil {
		log.Println(err)
	}

	session.Values["username"] = username
	session.Values["login"] = true

	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func LogoutGET(w http.ResponseWriter, r *http.Request) {
	session, err := session.SessionStore.Get(r, session.SessionAuth)
	if err != nil {
		log.Println(err)
	}

	session.Values["login"] = false

	sessions.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
