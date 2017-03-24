package handlers

import (
	"log"
	"net/http"

	"github.com/adamjedlicka/webapp/src/common"
	"github.com/adamjedlicka/webapp/src/models"
)

func LoginPOST(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	username := r.FormValue("username")
	password := r.FormValue("password")

	log.Println("Attemt to login :: username: ", username, ", password: ", password)

	u := models.NewUser()
	err := u.FindByUsername(username)
	if err != nil || password != u.Password() {
		http.Error(w, "Bad username or password!", http.StatusUnauthorized)
		return
	}

	authCookie, err := r.Cookie("auth")
	if err != nil {
		log.Println("Creating new auth cookie: ", authCookie.Value)
		authCookie = &http.Cookie{
			Name:  "auth",
			Value: common.RandString(32),
		}
	}

	s := models.NewSession()
	err = s.FindByName(authCookie.Value)
	if err != nil {
		s.SetName(authCookie.Value)
		s.SetUserID(u.ID())
	}

	err = s.Save()
	if err != nil {
		log.Fatal(err)
	}

	http.SetCookie(w, authCookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func LogoutGET(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("auth")
	if err != nil {
		log.Println(err)
		return
	}

	s := models.NewSession()
	err = s.FindByName(c.Value)
	if err != nil {
		log.Println(err)
		return
	}

	err = s.Delete()
	if err != nil {
		log.Println(err)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
