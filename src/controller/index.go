package controller

import (
	"net/http"

	"github.com/adamjedlicka/webapp/src/view"
)

func IndexGET(w http.ResponseWriter, r *http.Request) {
	v := view.New(r, "index")
	v.L["Title"] = "Home"

	v.Render(w)
}
