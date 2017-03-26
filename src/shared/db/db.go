package db

import (
	"database/sql"
	"log"

	// Import MySQL driver
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

type Config struct {
	Type     string
	Database string
	Username string
	Password string
}

func Connect(c Config) error {
	log.Println("Initializing database...")

	var err error
	DB, err = sql.Open(c.Type, c.Username+":@/"+c.Database)
	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		return err
	}

	// Connect and check the server version
	var version string
	DB.QueryRow("SELECT VERSION()").Scan(&version)
	log.Println("Connected to:", version)

	return nil
}

func Install(c Config) {
	log.Println("Installing database...")

	var err error
	DB, err = sql.Open(c.Type, c.Username+":@/")
	if err != nil {
		panic(err)
	}

	_, err = DB.Exec("DROP DATABASE IF EXISTS `" + c.Database + "`")
	if err != nil {
		panic(err)
	}

	_, err = DB.Exec("CREATE DATABASE `" + c.Database + "`")
	if err != nil {
		panic(err)
	}

	_, err = DB.Exec("USE `" + c.Database + "`")
	if err != nil {
		panic(err)
	}

	cmds := []string{
		`CREATE TABLE Users (
			ID INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
			FirstName VARCHAR(30),
			LastName VARCHAR(30),
			Username VARCHAR(30),
			Password VARCHAR(30))`,

		`INSERT INTO Users (FirstName, LastName, Username, Password)
			VALUES ("Franta", "Sádlo", "admin", "admin")`,
	}

	for _, v := range cmds {
		_, err = DB.Exec(v)
		if err != nil {
			panic(err)
		}
	}
}
