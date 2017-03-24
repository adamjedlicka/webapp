package handlers

import (
	"net/http"
)

var (
	templateDocuments = ParseTemplates("templates/documents.html")
)

func DocumentsGET(w http.ResponseWriter, r *http.Request) {
	templateDocuments.Execute(w, NewData(r))
}