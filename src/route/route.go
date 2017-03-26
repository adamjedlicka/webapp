package route

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/adamjedlicka/webapp/src/controller"
	"github.com/adamjedlicka/webapp/src/route/middleware"
)

func New() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", controller.IndexGET).Methods("GET")

	r.Handle("/projects", middleware.MustLogin(http.HandlerFunc(controller.ProjectsGET))).Methods("GET")
	r.Handle("/tasks", middleware.MustLogin(http.HandlerFunc(controller.TasksGET))).Methods("GET")
	r.Handle("/documents", middleware.MustLogin(http.HandlerFunc(controller.DocumentsGET))).Methods("GET")

	r.HandleFunc("/login", controller.LoginPOST).Methods("POST")
	r.HandleFunc("/logout", controller.LogoutGET).Methods("GET")

	// serve files from ./static/ directory without any special routing
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	return r
}
