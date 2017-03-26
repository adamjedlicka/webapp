package server

import (
	"flag"
	"log"
	"net/http"
)

var (
	flagAddress string
	flagPort    string
)

func init() {
	flag.StringVar(&flagAddress, "address", "localhost", "Sets the address of the web server")
	flag.StringVar(&flagPort, "port", "8080", "Sets the port of the webserver")
}

// Run the server
func Run(router http.Handler) {
	log.Println("Listening on address: ", flagAddress, " and port: ", flagPort)

	err := http.ListenAndServe(flagAddress+":"+flagPort, router)
	if err != nil {
		panic(err)
	}
}
