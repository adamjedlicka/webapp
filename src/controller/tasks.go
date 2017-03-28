package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/adamjedlicka/webapp/src/model"
	"github.com/adamjedlicka/webapp/src/view"
	"github.com/gorilla/mux"
)

func TasksActionGET(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	taskID, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t := model.NewTask()
	err = t.FindByID(taskID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if mux.Vars(r)["action"] == "delete" {
		err := t.Delete()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		http.Redirect(w, r, "/tasks", http.StatusSeeOther)
		return
	}

	v := view.New(r, "task/view")
	v.L["Title"] = "Tasks"

	v.Vars["Task"] = t
	v.Vars["Action"] = mux.Vars(r)["action"]

	v.Vars["readonly"] = ""
	if v.Vars["Action"] == "view" {
		v.Vars["readonly"] = "readonly"
	}

	v.Render(w)
}

func TasksNewGET(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	v := view.New(r, "task/view")
	v.Vars["Action"] = "new"

	v.Render(w)
}

func TasksPOST(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	taskID, err := strconv.ParseInt(r.FormValue("ID"), 10, 64)
	if err != nil {
		taskID = 0
	}

	t := model.NewTask()
	err = t.FindByID(taskID)
	if err != nil {
		log.Println("Creating new task...")
	}

	t.SetCode(r.FormValue("Code"))
	t.SetName(r.FormValue("Name"))
	t.SetDescription(r.FormValue("Description"))
	t.SetStartDate(model.Date(r.FormValue("StartDate")))
	t.SetPlanEndDate(model.Date(r.FormValue("PlanEndDate")))
	t.SetEndDate(model.Date(r.FormValue("EndDate")))

	projectID, err := strconv.ParseInt(r.FormValue("ProjectID"), 10, 64)
	if err == nil {
		t.SetProjectID(projectID)
	}

	userID, err := strconv.ParseInt(r.FormValue("UserID"), 10, 64)
	if err == nil {
		t.SetUserID(userID)
	} else if r.FormValue("UserID") == "" {
		t.SetUserID(-1)
	}

	err = t.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/tasks/view/"+strconv.FormatInt(t.ID(), 10), http.StatusSeeOther)
}

func TasksGET(w http.ResponseWriter, r *http.Request) {
	v := view.New(r, "task/tasks")
	v.AppendTemplates("task/list")

	v.L["Title"] = "Tasks"

	q := model.SelectTasks()
	tasks, err := model.QueryTasks(q)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	v.Vars["Tasks"] = tasks

	v.Render(w)
}
