package handlers

import (
	"html/template"
	"net/http"

	"github.com/adamjedlicka/webapp/src/common"
	"github.com/adamjedlicka/webapp/src/models"
)

var layouts = []string{
	"templates/layout/base.html",
}

type Data struct {
	IsLogin bool
	User    *models.User
}

func NewData(r *http.Request) *Data {
	d := new(Data)
	d.IsLogin = common.IsLogin(r)
	d.User, _ = common.GetUser(r)

	return d
}

func ParseTemplates(path ...string) *template.Template {
	tmpl, err := template.ParseFiles(append(layouts, path...)...)
	if err != nil {
		panic(err)
	}

	return tmpl
}
