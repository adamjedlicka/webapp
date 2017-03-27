package server

import (
	"log"
	"net/http"
)

type Configuration struct {
	Address string
	Port    string
}

var conf *Configuration

// Run the server
func Run(router http.Handler) {
	log.Println("Listening on address: ", conf.Address, " and port: ", conf.Port)

	err := http.ListenAndServe(conf.Address+":"+conf.Port, router)
	if err != nil {
		panic(err)
	}
}

func Configure(c *Configuration) {
	conf = c
}
