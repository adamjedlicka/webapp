package main

import (
	"flag"
	"fmt"

	"github.com/adamjedlicka/webapp/src/shared/cmd"
)

func main() {
	flag.Parse()

	switch flag.Arg(0) {
	case "run":
		cmd.RunInit()
	case "install":
		cmd.InstallInit()

	case "help":
		switch flag.Arg(1) {
		case "run":
			cmd.RunHelp()
		case "install":
			cmd.InstallHelp()
		}

	default:
		fmt.Println("ISSZP is information system for administration of employees and projects")
		fmt.Println()
		fmt.Println("Usage:")
		fmt.Println()
		fmt.Println("\t", "isszp command [arguments]")
		fmt.Println()
		fmt.Println("The commands are:")
		fmt.Println()
		fmt.Println("\t", "run", "\t\t", "run the webserver")
		fmt.Println("\t", "install", "\t", "install the database and webserver")
		fmt.Println()
		fmt.Println("Use \"isszp help [command]\" for more information")
		fmt.Println()
	}
}
