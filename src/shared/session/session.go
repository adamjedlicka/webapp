package session

import (
	"net/http"

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
