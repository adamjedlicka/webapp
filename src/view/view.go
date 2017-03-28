package view

import (
	"html/template"
	"net/http"

	"github.com/adamjedlicka/webapp/src/shared/session"
)

var (
	templateBase = "layout/base"
)

type View struct {
	Name string
	Vars map[string]interface{}
}

func New(r *http.Request, name string) *View {
	v := new(View)
	v.Name = name
	v.Vars = make(map[string]interface{})

	v.Vars["IsLogin"] = session.IsLogin(r)
	if v.Vars["IsLogin"].(bool) {
		u, err := session.GetUser(r)
		if err != nil {
			v.Vars["IsLogin"] = false
		}

		v.Vars["User"] = u
	}

	return v
}

func (v *View) Render(w http.ResponseWriter) {
	var templateList []string
	templateList = append(templateList, templateBase)
	templateList = append(templateList, v.Name)

	for i, name := range templateList {
		path := "template/" + name + ".html"

		templateList[i] = path
	}

	template, err := template.ParseFiles(templateList...)
	if err != nil {
		http.Error(w, "Template Parse Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	template.Execute(w, v.Vars)
}
