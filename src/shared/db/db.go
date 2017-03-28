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

var DB *sql.DB

func Connect() {
	log.Println("Initializing database...")

	var err error
	DB, err = sql.Open(conf.Type, conf.Username+":@/"+conf.Database)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Connect and check the server version
	var version string
	DB.QueryRow("SELECT VERSION()").Scan(&version)
	log.Println("Connected to:", version)
}

func Exec(q string, args ...interface{}) (sql.Result, error) {
	return DB.Exec(q, args...)
}

func QueryRow(q string, args ...interface{}) *sql.Row {
	return DB.QueryRow(q, args...)
}

func Query(q string, args ...interface{}) (*sql.Rows, error) {
	return DB.Query(q, args...)
}

func Configure(c *Configuration) { conf = c }
