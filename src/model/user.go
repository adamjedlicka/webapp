package model

import (
	"database/sql"
	"log"

	"github.com/adamjedlicka/webapp/src/shared/db"
)

type User struct {
	id         int64
	firstName  sql.NullString
	lastName   sql.NullString
	username   string
	password   string
	isEmployee bool
	email      sql.NullString
	telNumber  sql.NullString
}

func NewUser() *User {
	u := new(User)
	u.id = -1

	return u
}

func (u *User) FindByUsername(username string) error {
	err := db.QueryRow(`SELECT ID, FirstName, LastName, Username, Password, IsEmployee, EMail, TelNumber
	                      FROM Users WHERE Username = ?`, username).
		Scan(&u.id, &u.firstName, &u.lastName, &u.username, &u.password, &u.isEmployee, &u.email, &u.telNumber)

	return err
}

func (u *User) FindByID(id int64) error {
	err := db.QueryRow(`SELECT ID, FirstName, LastName, Username, Password, IsEmployee, EMail, TelNumber
	                      FROM Users WHERE ID = ?`, id).
		Scan(&u.id, &u.firstName, &u.lastName, &u.username, &u.password, &u.isEmployee, &u.email, &u.telNumber)

	return err
}

func (u User) ID() int64         { return u.id }
func (u User) FirstName() string { return u.firstName.String }
func (u User) LastName() string  { return u.lastName.String }
func (u User) Username() string  { return u.username }
func (u User) IsEmployee() bool  { return u.isEmployee }

func (u *User) SetUsername(username string)   { u.username = username }
func (u *User) SetIsEmployee(isEmployee bool) { u.isEmployee = isEmployee }
func (u *User) SetFirstName(firstName string) {
	u.firstName.String = firstName
	u.firstName.Valid = firstName != ""
}
func (u *User) SetLastName(lastName string) {
	u.lastName.String = lastName
	u.lastName.Valid = lastName != ""
}

func (u User) CheckPassword(password string) bool { return password == u.password }

func (u *User) Save() error {
	if u.id == -1 {
		res, err := db.Exec("INSERT INTO Users (FirstName, LastName, Username, IsEmployee, Password) VALUES (?, ?, ?, ?, ?)",
			u.firstName.String, u.lastName.String, u.username, u.isEmployee, u.password)
		if err != nil {
			return err
		}

		id, err := res.LastInsertId()
		if err != nil {
			return err
		}

		u.id = id
	}

	return nil
}

func GetUsers() []*User {
	users := make([]*User, 0)

	res, err := db.Query("SELECT ID, FirstName, LastName, Username, IsEmployee FROM Users ORDER BY ID")
	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {
		u := NewUser()
		res.Scan(&u.id, &u.firstName, &u.lastName, &u.username, &u.isEmployee)
		users = append(users, u)
	}

	return users
}
