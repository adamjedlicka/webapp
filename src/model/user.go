package model

import (
	"errors"
	"net/http"
)

type User struct {
	id       int64
	username string
	password string
}

var users map[string]*User

func init() {
	users = make(map[string]*User)
	users["admin"] = &User{
		id:       1,
		username: "admin",
		password: "admin",
	}
}

func NewUser() *User {
	u := new(User)
	u.id = -1

	return u
}

func (u *User) FindByUsername(username string) error {
	tmp, ok := users[username]
	if !ok {
		return errors.New("User not found: " + username)
	}

	u.id = tmp.id
	u.username = tmp.username
	u.password = tmp.password

	return nil
}

func (u *User) FindByID(id int64) error {
	for _, tmp := range users {
		if tmp.id == id {
			u.id = tmp.id
			u.username = tmp.username
			u.password = tmp.password
			return nil
		}
	}

	return errors.New("No such user")
}

func (u User) ID() int64        { return u.id }
func (u User) Username() string { return u.username }
func (u User) Password() string { return u.password }

func GetUser(r *http.Request) (*User, error) {
	u := NewUser()
	err := u.FindByID(1)
	if err != nil {
		return nil, errors.New("No user logged in!")
	}

	return u, nil
}
