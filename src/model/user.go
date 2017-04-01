package model

import (
	"database/sql"
	"fmt"

	"github.com/adamjedlicka/webapp/src/shared/db"
)

type User struct {
	ID           UUID           `db:"ID"`
	UserName     string         `db:"UserName"`
	Password     string         `db:"Password"`
	FirstName    sql.NullString `db:"FirstName"`
	LastName     sql.NullString `db:"LastName"`
	PermissionID UUID           `db:"Permission_ID"`

	Permission *Permission
}

func (u *User) Fill() error {
	err := db.Get(u, "SELECT * FROM Users WHERE ID = ?", u.ID)
	if err != nil {
		return fmt.Errorf("Couldn't fill User with ID %s: %v", u.ID, err)
	}

	u.Permission = &Permission{ID: u.PermissionID}
	err = u.Permission.Fill()
	if err != nil {
		return fmt.Errorf("Couldn't fill User with ID %s: %v", u.ID, err)
	}

	return nil
}

func (u *User) FindByID(id UUID) error {
	return db.Get(u, "SELECT * FROM Users WHERE ID LIKE ?", id.String()+"%")
}

func (u *User) FindByUserName(username string) error {
	return db.Get(u, "SELECT * FROM Users WHERE UserName = ?", username)
}

func (u *User) Save() error {
	_, err := db.NamedExec(`
	INSERT INTO Users (ID, UserName, Password, FirstName, LastName, Permission_ID) 
		VALUES (UUID(), :UserName, :Password, :FirstName, :LastName, :Permission_ID)`, u)
	return err
}

func (u User) CheckPassword(password string) bool {
	return password == u.Password
}
