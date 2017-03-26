package controller

import (
	"net/http"
)

var (
	templateIndex = ParseTemplates("templates/index.html")
)

func IndexGET(w http.ResponseWriter, r *http.Request) {
	templateIndex.Execute(w, NewBase(r))
}
