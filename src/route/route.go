package route

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/adamjedlicka/webapp/src/controller"
	"github.com/adamjedlicka/webapp/src/route/middleware"
)

type Configuration struct {
	LogRequests bool
}

var conf *Configuration

func New() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/", controller.IndexGET).Methods("GET")

	r.HandleFunc("/login", controller.LoginPOST).Methods("POST")
	r.HandleFunc("/logout", controller.LogoutGET).Methods("GET")

	r.Handle("/tasks", middleware.MustLogin(http.HandlerFunc(controller.TasksGET))).Methods("GET")
	r.Handle("/tasks/view/{id}", middleware.MustLogin(http.HandlerFunc(controller.TasksViewGET))).Methods("GET")

	// serve files from ./static/ directory without any special routing
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	if conf.LogRequests {
		return middleware.LogRequest(r)
	}

	return r
}

func Configure(c *Configuration) { conf = c }
