package server

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	flagAddress string
	flagPort    string
)

func init() {
	flag.StringVar(&flagAddress, "flag", "localhost", "Sets the address of the web server")
	flag.StringVar(&flagPort, "port", "8080", "Sets the port of the webserver")
}

func Run() {
	r := mux.NewRouter()

	initRoutes(r)

	log.Println("Listening on address: ", flagAddress, " and port: ", flagPort)

	err := http.ListenAndServe(flagAddress+":"+flagPort, r)
	if err != nil {
		panic(err)
	}
}
