package model

import (
	"database/sql"

	"github.com/adamjedlicka/webapp/src/shared/db"
)

type Firm struct {
	ID          string         `db:"ID"`
	Name        string         `db:"Name"`
	Description sql.NullString `db:"Description"`
	Email       sql.NullString `db:"Email"`
	TelNumber   sql.NullInt64  `db:"TelNumber"`
}

func (f *Firm) Fill() error {
	return db.Get(f, "SELECT * FROM Firms WHERE ID = ?", f.ID)
}

func (f *Firm) Save() error {
	_, err := db.NamedExec(`
	INSERT INTO Firms (ID, Name, Description, Email, TelNumber)
		VALUES (UUID(), :Name, :Description, :Email, :TelNumber)`, f)

	return err
}
