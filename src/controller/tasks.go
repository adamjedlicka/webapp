package controller

import (
	"net/http"

	"github.com/adamjedlicka/webapp/src/model"
	"github.com/adamjedlicka/webapp/src/shared/db"
	"github.com/adamjedlicka/webapp/src/view"
	"github.com/gorilla/mux"
)

func TasksGET(w http.ResponseWriter, r *http.Request) {
	v := view.New(r, "Tasks")
	v.AppendTemplates("tasks/tasks", "component/task-list")

	tasks := []model.Task{}
	db.Select(&tasks, "SELECT * FROM Tasks")
	v.Vars["Tasks"] = tasks

	v.Render(w)
}

func TasksViewGET(w http.ResponseWriter, r *http.Request) {
	v := view.New(r, "Tasks - view")
	v.AppendTemplates("tasks/view")

	task := model.Task{ID: mux.Vars(r)["id"]}
	err := task.Fill()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	v.Vars["Task"] = task
	v.Vars["Action"] = "view"

	v.Render(w)
}
