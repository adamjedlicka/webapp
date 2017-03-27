package controller

import (
	"net/http"

	"github.com/adamjedlicka/webapp/src/view"
)

func IndexGET(w http.ResponseWriter, r *http.Request) {
	v := view.New(r)
	v.Name = "index"

	v.Render(w)
}
