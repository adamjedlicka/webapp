package models

import (
	"testing"
)

func TestFindByUsername(t *testing.T) {
	u := NewUser()
	err := u.FindByUsername("admin")
	if err != nil {
		t.Error(err)
	}

	if u.id != users["admin"].id {
		t.Error("ID mismatch")
	}

	if u.username != users["admin"].username {
		t.Error("Username mismatch")
	}
}

func TestFindByID(t *testing.T) {
	u := NewUser()
	err := u.FindByID(1)
	if err != nil {
		t.Error(err)
	}

	if u.id != users["admin"].id {
		t.Error("ID mismatch")
	}

	if u.username != users["admin"].username {
		t.Error("Username mismatch")
	}
}
