package model

import (
	"errors"
	"net/http"

	"github.com/adamjedlicka/webapp/src/shared/db"
	"github.com/adamjedlicka/webapp/src/shared/session"
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
	err := db.QueryRow("SELECT ID, FirstName, LastName, Username, Password FROM Users WHERE Username = ?", username).
		Scan(&u.id, &u.firstName, &u.lastName, &u.username, &u.password)

	return err
}

func (u *User) FindByID(id int64) error {
	err := db.QueryRow("SELECT ID, FirstName, LastName, Username, Password FROM Users WHERE ID = ?", id).
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

	s, err := session.SessionStore.Get(r, session.SessionAuth)
	if err != nil {
		return nil, errors.New("No user logged in!")
	}

	id, ok := s.Values["id"].(int64)
	if !ok {
		return nil, errors.New("Wrong ID in session")
	}

	err = u.FindByID(id)
	if err != nil {
		return nil, errors.New("No such user in databse!")
	}

	return u, nil
}
