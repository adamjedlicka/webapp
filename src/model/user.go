package model

import (
	"errors"
	"net/http"

	"github.com/adamjedlicka/webapp/src/shared/db"
)

type User struct {
	id        int64
	firstName string
	lastName  string
	username  string
	password  string
}

func NewUser() *User {
	u := new(User)
	u.id = -1

	return u
}

func (u *User) FindByUsername(username string) error {
	err := db.DB.QueryRow("SELECT ID, FirstName, LastName, Username, Password FROM Users WHERE Username = ?", username).
		Scan(&u.id, &u.firstName, &u.lastName, &u.username, &u.password)

	return err
}

func (u *User) FindByID(id int64) error {
	err := db.DB.QueryRow("SELECT ID, FirstName, LastName, Username, Password FROM Users WHERE ID = ?", id).
		Scan(&u.id, &u.firstName, &u.lastName, &u.username, &u.password)

	return err
}

func (u User) ID() int64         { return u.id }
func (u User) FirstName() string { return u.firstName }
func (u User) LastName() string  { return u.lastName }
func (u User) Username() string  { return u.username }

func (u User) CheckPassword(password string) bool { return password == u.password }

func GetUser(r *http.Request) (*User, error) {
	u := NewUser()
	err := u.FindByID(1)
	if err != nil {
		return nil, errors.New("No user logged in!")
	}

	return u, nil
}
