package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"os/signal"

	"fmt"

	"github.com/adamjedlicka/webapp/src/route"
	"github.com/adamjedlicka/webapp/src/shared/db"
	"github.com/adamjedlicka/webapp/src/shared/server"
	"github.com/adamjedlicka/webapp/src/shared/util"
)

var (
	flagHelp    bool
	flagRun     bool
	flagInstall bool
	flagConfig  bool
)

func init() {
	flag.BoolVar(&flagHelp, "help", false, "Prints the help message")
	flag.BoolVar(&flagInstall, "install", false, "Installs the applications")
	flag.BoolVar(&flagRun, "run", false, "Runs the application")
	flag.BoolVar(&flagConfig, "config", false, "Displays current configuration")
}

func main() {
	flag.Parse()

	util.LoadConfig("config"+string(os.PathSeparator)+"config.json", config)
	if flagConfig {
		fmt.Println(config)
		os.Exit(0)
	}

	// Handle interrupt signals from terminal (Ctrl-C)
	signals := make(chan os.Signal)
	signal.Notify(signals, os.Interrupt)
	go func() {
		<-signals
		log.Println("Stopping server...")
		os.Exit(0)
	}()

	// Double check install process
	if flagInstall {
		fmt.Println("This will delete all data in database and installs the application!")
		fmt.Print("Are you sure you want to proceed? [y/N] ")
		var res rune
		fmt.Scanf("%c", &res)
		if res != 'y' && res != 'Y' {
			os.Exit(0)
		}

		db.Install(config.Database)

		fmt.Println("Instalation successfull")
		os.Exit(0)
	}

	if !flagRun || flagHelp {
		flag.PrintDefaults()
		os.Exit(0)
	}

	err := db.Connect(config.Database)
	if err != nil {
		panic(err)
	}

	router := route.New()

	server.Run(router)
}

// *****************************************************************************
// Application Settings
// *****************************************************************************

// config the settings variable
var config = &configuration{}

type configuration struct {
	Database db.Config `json:"Database"`
}

// ParseJSON unmarshals bytes to structs
func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}
