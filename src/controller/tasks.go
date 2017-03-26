package controller

import (
	"net/http"
)

var (
	templateTasks = ParseTemplates("templates/tasks.html")
)

func TasksGET(w http.ResponseWriter, r *http.Request) {
	templateTasks.Execute(w, NewBase(r))
}
