package controller

import (
	"html/template"
	"net/http"

	"github.com/adamjedlicka/webapp/src/model"
	"github.com/adamjedlicka/webapp/src/shared/session"
)

var layouts = []string{
	"templates/layout/base.html",
}

type Base struct {
	IsLogin bool
	User    *model.User
}

func NewBase(r *http.Request) *Base {
	b := new(Base)
	b.IsLogin = session.IsLogin(r)
	b.User, _ = model.GetUser(r)

	return b
}

func ParseTemplates(path ...string) *template.Template {
	tmpl, err := template.ParseFiles(append(layouts, path...)...)
	if err != nil {
		panic(err)
	}

	return tmpl
}
