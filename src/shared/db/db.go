package db

import (
	"log"

	// Import MySQL driver
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Configuration struct {
	Type     string
	Database string
	Username string
	Password string
}

var conf *Configuration

var database *sqlx.DB

func Connect() {
	log.Println("Initializing database...")

	var err error
	database, err = sqlx.Open(conf.Type, conf.Username+":@/"+conf.Database)
	if err != nil {
		log.Fatal(err)
	}

	err = database.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Connect and check the server version
	var version string
	database.QueryRow("SELECT VERSION()").Scan(&version)
	log.Println("Connected to:", version)
}

func Get(dest interface{}, query string, args ...interface{}) error {
	return database.Get(dest, query, args...)
}

func Select(dest interface{}, query string, args ...interface{}) error {
	return database.Select(dest, query, args...)
}

func NamedExec(query string, arg interface{}) (sql.Result, error) {
	return database.NamedExec(query, arg)
}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	return database.Exec(query, args...)
}

func Configure(c *Configuration) { conf = c }
