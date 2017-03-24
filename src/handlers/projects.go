package handlers

import (
	"net/http"
)

var (
	templateProjects = ParseTemplates("templates/projects.html")
)

func ProjectsGET(w http.ResponseWriter, r *http.Request) {
	templateProjects.Execute(w, NewData(r))
}
