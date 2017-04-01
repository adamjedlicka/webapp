package controller

import (
	"net/http"

	"github.com/adamjedlicka/webapp/src/model"
	"github.com/adamjedlicka/webapp/src/shared/db"
	"github.com/adamjedlicka/webapp/src/shared/session"
	"github.com/adamjedlicka/webapp/src/view"
)

func IndexGET(w http.ResponseWriter, r *http.Request) {
	v := view.New(r, "index")
	v.AppendTemplates("task/list")

	if session.IsLogin(r) {
		userID, err := session.GetUserID(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		tasks := []model.Task{}
		err = db.Select(&tasks, "SELECT * FROM Tasks WHERE User_ID_Worker = ?", userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		v.Vars["Tasks"] = tasks
	}

	v.Render(w)
}
