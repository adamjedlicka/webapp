package controller

import (
	"net/http"

	"github.com/adamjedlicka/webapp/src/view"
)

func ProjectsGET(w http.ResponseWriter, r *http.Request) {
	v := view.New(r)
	v.Name = "project/projects"

	v.Render(w)
}
