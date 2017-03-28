package controller

import (
	"net/http"

	"github.com/adamjedlicka/webapp/src/model"
	"github.com/adamjedlicka/webapp/src/shared/session"
	"github.com/adamjedlicka/webapp/src/view"
)

func IndexGET(w http.ResponseWriter, r *http.Request) {
	v := view.New(r, "index")
	v.L["Title"] = "Home"

	if session.IsLogin(r) {
		id, err := session.GetUserID(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		q := model.SelectTasks().Where("User_ID = ?", id)
		tasks, err := model.QueryTasks(q)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		v.Vars["Tasks"] = tasks
		v.AppendTemplates("task/list")
	}

	v.Render(w)
}
