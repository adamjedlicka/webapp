package main

import (
	"flag"
	"os"
	"os/signal"

	"log"

	"github.com/adamjedlicka/webapp/src/server"
)

var (
	flagHelp bool
)

func init() {
	flag.BoolVar(&flagHelp, "help", false, "Prints the help message")
}

func main() {
	flag.Parse()
	if flagHelp {
		flag.PrintDefaults()
		os.Exit(0)
	}

	signals := make(chan os.Signal)
	signal.Notify(signals, os.Interrupt)
	go func() {
		<-signals
		log.Println("Stopping server...")
		os.Exit(0)
	}()

	server.Run()
}
