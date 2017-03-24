package common

import (
	"errors"
	"math/rand"
	"net/http"
	"time"

	"github.com/adamjedlicka/webapp/src/models"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// RandString generates random string of length n
func RandString(n int) string {
	var randBytes = make([]byte, n)
	rand.Read(randBytes)

	for i, b := range randBytes {
		randBytes[i] = letters[b%byte(len(letters))]
	}

	return string(randBytes)
}

func IsLogin(r *http.Request) bool {
	c, err := r.Cookie("auth")
	if err != nil {
		return false
	}

	s := models.NewSession()
	err = s.FindByName(c.Value)
	if err != nil {
		return false
	}

	return true
}

func GetUser(r *http.Request) (*models.User, error) {
	s, err := GetSession(r)
	if err != nil {
		return nil, err
	}

	u := models.NewUser()
	err = u.FindByID(s.UserID())
	if err != nil {
		return nil, errors.New("No user logged in!")
	}

	return u, nil
}

func GetSession(r *http.Request) (*models.Session, error) {
	c, err := r.Cookie("auth")
	if err != nil {
		return nil, errors.New("No user logged in!")
	}

	s := models.NewSession()
	err = s.FindByName(c.Value)
	if err != nil {
		return nil, errors.New("No user logged in!")
	}

	return s, nil
}
