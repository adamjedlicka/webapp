package controller

import (
	"log"
	"net/http"

	"github.com/adamjedlicka/webapp/src/model"
	"github.com/adamjedlicka/webapp/src/view"
)

func ProjectsGET(w http.ResponseWriter, r *http.Request) {
	v := view.New(r)
	v.Name = "project/projects"
	v.Vars["Projects"] = model.GetProjects()

	v.Render(w)
}

func ProjectsNewGET(w http.ResponseWriter, r *http.Request) {
	v := view.New(r)
	v.Name = "project/new"

	v.Render(w)
}

func ProjectsNewPOST(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	p := model.NewProject()
	p.SetName(r.FormValue("name"))
	p.SetDescription(r.FormValue("description"))
	err = p.Save()
	if err != nil {
		log.Fatal(err)
	}

	http.Redirect(w, r, "/projects", http.StatusSeeOther)
}
