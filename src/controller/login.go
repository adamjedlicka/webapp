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

	u := model.User{}
	err := u.FindByUserName(username)
	if err != nil || !u.CheckPassword(password) {
		http.Error(w, "Bad username or password!", http.StatusUnauthorized)
		return
	}

	session, err := session.SessionStore.Get(r, session.SessionAuth)
	if err != nil {
		http.Error(w, "Corrupted session", http.StatusInternalServerError)
		return
	}

	session.Values["id"] = u.ID.String()
	session.Values["username"] = u.UserName
	session.Values["login"] = true

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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
