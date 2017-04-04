package view

import (
	"html/template"
	"net/http"
	"time"

	"github.com/adamjedlicka/webapp/src/shared/session"
)

var (
	templateBase = "layout/base"
)

type View struct {
	Name      string
	Vars      map[string]interface{}
	L         map[string]string
	templates []string
}

func New(r *http.Request, name string) *View {
	v := new(View)
	v.Name = name
	v.Vars = make(map[string]interface{})
	v.L = make(map[string]string)
	v.templates = make([]string, 0)
	v.templates = append(v.templates, templateBase)

	v.Vars["Name"] = v.Name
	v.Vars["IsLogin"] = session.IsLogin(r)
	if v.Vars["IsLogin"].(bool) {
		u, err := session.GetUser(r)
		if err != nil {
			v.Vars["IsLogin"] = false
		}

		v.Vars["User"] = u
	}

	// sets var Date to the format yyyy-mm-dd
	v.Vars["Date"] = time.Now().String()[:10]

	return v
}

func (v *View) AppendTemplates(templates ...string) {
	v.templates = append(v.templates, templates...)
}

func (v *View) Render(w http.ResponseWriter) {
	templateList := make([]string, len(v.templates))

	for i, name := range v.templates {
		path := "template/" + name + ".html"

		templateList[i] = path
	}

	template, err := template.ParseFiles(templateList...)
	if err != nil {
		http.Error(w, "Template Parse Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	v.Vars["L"] = v.L
	template.Execute(w, v.Vars)
}
