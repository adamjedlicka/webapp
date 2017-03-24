package models

import (
	"errors"
)

var sessions map[string]*Session

func init() {
	sessions = make(map[string]*Session)
}

type Session struct {
	id     int64
	name   string
	userID int64
}

func NewSession() *Session {
	s := new(Session)
	s.id = 0
	s.userID = 0

	return s
}

func (s *Session) FindByName(name string) error {
	tmp, ok := sessions[name]
	if !ok {
		return errors.New("No such session")
	}

	s.id = tmp.id
	s.name = tmp.name
	s.userID = tmp.userID

	return nil
}

func (s *Session) Save() error {
	sessions[s.name] = s

	return nil
}

func (s *Session) Delete() error {
	delete(sessions, s.name)

	return nil
}

func (s Session) Name() string  { return s.name }
func (s Session) UserID() int64 { return s.userID }

func (s *Session) SetName(name string)    { s.name = name }
func (s *Session) SetUserID(userID int64) { s.userID = userID }
