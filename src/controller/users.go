package controller

import (
	"net/http"

	"github.com/adamjedlicka/webapp/src/model"
	"github.com/adamjedlicka/webapp/src/view"
)

func UsersGET(w http.ResponseWriter, r *http.Request) {
	v := view.New(r)
	v.Name = "user/users"
	v.Vars["Users"] = model.GetUsers()

	v.Render(w)
}

func UsersNewGET(w http.ResponseWriter, r *http.Request) {
	v := view.New(r)
	v.Name = "user/new"
	v.Vars["Users"] = model.GetUsers()

	v.Render(w)
}

func UsersNewPOST(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}

	u := model.NewUser()
	u.SetUsername(r.FormValue("Username"))
	u.SetFirstName(r.FormValue("FirstName"))
	u.SetLastName(r.FormValue("LastName"))
	u.SetIsEmployee(r.FormValue("IsEmployee") == "on")

	err = u.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/users", http.StatusSeeOther)
}
