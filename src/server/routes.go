package server

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/adamjedlicka/webapp/src/handlers"
	"github.com/adamjedlicka/webapp/src/middleware"
)

func initRoutes(r *mux.Router) {
	r.HandleFunc("/", handlers.IndexGET).Methods("GET")

	r.Handle("/projects", middleware.MustLogin(http.HandlerFunc(handlers.ProjectsGET))).Methods("GET")
	r.Handle("/tasks", middleware.MustLogin(http.HandlerFunc(handlers.TasksGET))).Methods("GET")
	r.Handle("/documents", middleware.MustLogin(http.HandlerFunc(handlers.DocumentsGET))).Methods("GET")

	r.HandleFunc("/login", handlers.LoginPOST).Methods("POST")
	r.HandleFunc("/logout", handlers.LogoutGET).Methods("GET")

	// serve files from ./static/ directory without any special routing
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
}
