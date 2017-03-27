package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"

	"io/ioutil"

	"github.com/adamjedlicka/webapp/src/shared/config"
)

func InstallInit() {
	fmt.Println("This will delete all data in database and installs the application!")
	fmt.Print("Are you sure you want to proceed? [y/N] ")
	var res rune
	fmt.Scanf("%c", &res)
	if res != 'y' && res != 'Y' {
		os.Exit(0)
	}

	config.LoadConfig("config/config.json")

	installDB()
}

func InstallHelp() {
	fmt.Println("Install command help")
}

func installDB() {
	log.Println("Installing database...")

	db, err := sql.Open(config.Config.Database.Type, config.Config.Database.Username+":@/")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("DROP DATABASE IF EXISTS `" + config.Config.Database.Database + "`")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE DATABASE `" + config.Config.Database.Database + "`")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("USE `" + config.Config.Database.Database + "`")
	if err != nil {
		panic(err)
	}

	{ // DDL SQL statements
		data, err := ioutil.ReadFile("config/ddl.sql")
		if err != nil {
			panic(err)
		}

		stmts := strings.Split(string(data), ";")
		stmts = stmts[:len(stmts)-1]

		for _, stmt := range stmts {
			_, err := db.Exec(stmt)
			if err != nil {
				log.Println("ERR :: ", err)
				log.Println(stmt)
			}
		}
	}

	{ // DML SQL statements
		data, err := ioutil.ReadFile("config/dml.sql")
		if err != nil {
			panic(err)
		}

		stmts := strings.Split(string(data), ";")
		stmts = stmts[:len(stmts)-1]

		for _, stmt := range stmts {
			_, err := db.Exec(stmt)
			if err != nil {
				log.Println("ERR :: ", err)
				log.Println(stmt)
			}
		}
	}

	log.Println("Database installation successfull!")
}
