package db

import (
	"database/sql"
	"log"

	// Import MySQL driver
	_ "github.com/go-sql-driver/mysql"
)

type Configuration struct {
	Type     string
	Database string
	Username string
	Password string
}

var conf *Configuration

var db *sql.DB

func Connect() {
	log.Println("Initializing database...")

	var err error
	db, err = sql.Open(conf.Type, conf.Username+":@/"+conf.Database)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Connect and check the server version
	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	log.Println("Connected to:", version)
}

func Exec(q string, args ...interface{}) (sql.Result, error) {
	return db.Exec(q, args...)
}

func QueryRow(q string, args ...interface{}) *sql.Row {
	return db.QueryRow(q, args...)
}

func Query(q string, args ...interface{}) (*sql.Rows, error) {
	return db.Query(q, args...)
}

func Configure(c *Configuration) { conf = c }
