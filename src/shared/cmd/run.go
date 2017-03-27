package cmd

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/adamjedlicka/webapp/src/route"
	"github.com/adamjedlicka/webapp/src/shared/config"
	"github.com/adamjedlicka/webapp/src/shared/db"
	"github.com/adamjedlicka/webapp/src/shared/server"
)

func RunInit() {
	// Handle interrupt signals from terminal (Ctrl-C)
	signals := make(chan os.Signal)
	signal.Notify(signals, os.Interrupt)
	go func() {
		<-signals
		log.Println("Stopping server...")
		os.Exit(0)
	}()

	config.LoadConfig("config/config.json")

	db.Configure(&config.Config.Database)
	db.Connect()

	route.Configure(&config.Config.Router)
	r := route.New()

	server.Configure(&config.Config.Server)
	server.Run(r)
}

func RunHelp() {
	fmt.Println("Server command help")
}
