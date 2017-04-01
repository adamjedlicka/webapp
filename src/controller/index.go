package controller

import (
	"net/http"

	"fmt"

	"github.com/adamjedlicka/webapp/src/model"
	"github.com/adamjedlicka/webapp/src/shared/db"
	"github.com/adamjedlicka/webapp/src/shared/session"
	"github.com/adamjedlicka/webapp/src/view"
)

func IndexGET(w http.ResponseWriter, r *http.Request) {
	v := view.New(r, "Home")
	v.AppendTemplates("index", "component/task-list")

	if session.IsLogin(r) {
		userID, err := session.GetUserID(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tasks := []model.Task{}
		err = db.Select(&tasks, "SELECT * FROM Tasks WHERE User_ID_Maintainer = ?", userID.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		v.Vars["Tasks"] = tasks
	} else {
		fmt.Println("Not logged in!")
	}

	v.Render(w)
}
