package session

import (
	"errors"
	"net/http"

	"github.com/adamjedlicka/webapp/src/model"
	"github.com/gorilla/sessions"
)

const (
	// SessionAuth is session name for standart authentication session store
	SessionAuth = "auth"
)

var (
	// authKey is used to authentificate cookies received from user.
	// Static authKey allows accepting old cookies even after server restart.
	// Generating new authKey every start is more secure but deletes old sessions of all users.
	authKey = []byte("zO6vVDycza42VAWBDq1n9OwLLFqPmNYX")

	// SessionStore can be used to get/or create new session based on its session name
	SessionStore = sessions.NewCookieStore(authKey)
)

// IsLogin is helper function that checks if user is logged in
func IsLogin(r *http.Request) bool {
	session, err := SessionStore.Get(r, SessionAuth)
	if err != nil {
		return false
	}

	val, ok := session.Values["login"].(bool)
	if !ok {
		return false
	}

	return val
}

func GetUser(r *http.Request) (model.User, error) {
	u := model.User{}

	session, err := SessionStore.Get(r, SessionAuth)
	if err != nil {
		return u, errors.New("No user logged in")
	}

	id, ok := session.Values["id"].(string)
	if !ok {
		return u, errors.New("No user logged in")
	}

	err = u.FindByID(id)
	if err != nil {
		return u, errors.New("No such User in database")
	}

	return u, nil
}

func GetUserID(r *http.Request) (string, error) {
	session, err := SessionStore.Get(r, SessionAuth)
	if err != nil {
		return "", errors.New("No user logged in")
	}

	id, ok := session.Values["id"].(string)
	if !ok {
		return "", errors.New("No user logged in")
	}

	return id, nil
}
