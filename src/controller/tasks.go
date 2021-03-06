package controller

import (
	"net/http"
	"time"

	"github.com/adamjedlicka/webapp/src/model"
	"github.com/adamjedlicka/webapp/src/shared/db"
	"github.com/adamjedlicka/webapp/src/view"
	"github.com/gorilla/mux"
)

func TasksGET(w http.ResponseWriter, r *http.Request) {
	v := view.New(r, "Tasks")
	v.AppendTemplates("tasks/tasks", "component/task-list")

	tasks := []model.Task{}
	err := db.Select(&tasks, "SELECT * FROM Tasks")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	v.Vars["Tasks"] = tasks

	v.Render(w)
}

func TasksViewGET(w http.ResponseWriter, r *http.Request) {
	v := view.New(r, "Tasks")
	v.AppendTemplates("tasks/view")

	task := model.Task{ID: model.UUID(mux.Vars(r)["id"])}
	err := task.Fill()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	users := []model.User{}
	db.Select(&users, "SELECT * FROM Users")

	projects := []model.Project{}
	db.Select(&projects, "SELECT * FROM Projects")

	v.Vars["Task"] = task
	v.Vars["Users"] = users
	v.Vars["Projects"] = projects
	v.Vars["Action"] = "view"

	v.Vars["readonly"] = "readonly"
	v.Vars["disabled"] = "disabled"

	v.Render(w)
}

func TasksEditGET(w http.ResponseWriter, r *http.Request) {
	v := view.New(r, "Tasks")
	v.AppendTemplates("tasks/view")

	task := model.Task{ID: model.UUID(mux.Vars(r)["id"])}
	err := task.Fill()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	users := []model.User{}
	db.Select(&users, "SELECT * FROM Users")

	projects := []model.Project{}
	db.Select(&projects, "SELECT * FROM Projects")

	v.Vars["Task"] = task
	v.Vars["Users"] = users
	v.Vars["Projects"] = projects
	v.Vars["Action"] = "edit"

	v.Render(w)
}

func TasksNewGET(w http.ResponseWriter, r *http.Request) {
	v := view.New(r, "Tasks")
	v.AppendTemplates("tasks/view")

	users := []model.User{}
	db.Select(&users, "SELECT * FROM Users")

	projects := []model.Project{}
	db.Select(&projects, "SELECT * FROM Projects")

	v.Vars["Task"] = model.Task{
		StartDate: model.Date(time.Now().String()[:10]),
	}
	v.Vars["Users"] = users
	v.Vars["Projects"] = projects
	v.Vars["Action"] = "new"

	v.Render(w)
}

func TasksDeleteGET(w http.ResponseWriter, r *http.Request) {
	id, ok := mux.Vars(r)["id"]
	if !ok {
		http.Error(w, "Bad task ID!", http.StatusBadRequest)
		return
	}

	task := model.Task{ID: model.UUID(id)}
	err := task.Delete()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/tasks", http.StatusSeeOther)
}

func TasksSavePOST(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	task := model.Task{}
	task.ID.Scan(r.FormValue("ID"))
	task.Name = r.FormValue("Name")
	task.Description.Scan(r.FormValue("Description"))
	task.StartDate.Scan(r.FormValue("StartDate"))
	task.PlanEndDate.Scan(r.FormValue("PlanEndDate"))
	task.EndDate.Scan(r.FormValue("EndDate"))
	task.MaintainerID.Scan(r.FormValue("MaintainerID"))
	task.WorkerID.Scan(r.FormValue("WorkerID"))
	task.ProjectID.Scan(r.FormValue("ProjectID"))

	err := task.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if task.ID.String() == "" {
		http.Redirect(w, r, "/tasks", http.StatusSeeOther)
	}
	http.Redirect(w, r, "/tasks/view/"+task.ID.String(), http.StatusSeeOther)
}
