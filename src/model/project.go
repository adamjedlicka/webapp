package model

import (
	"database/sql"

	"github.com/adamjedlicka/webapp/src/shared/db"
)

type Project struct {
	ID          string         `db:"ID"`
	Name        string         `db:"Name"`
	Code        string         `db:"Code"`
	Description sql.NullString `db:"Description"`
	StartDate   Date           `db:"StartDate"`
	PlanEndDate NullDate       `db:"PlanEndDate"`
	EndDate     NullDate       `db:"EndDate"`
	UserID      string         `db:"User_ID"`
	FirmID      string         `db:"Firm_ID"`
}

func (p *Project) Fill() error {
	err := db.Get(p, "SELECT * FROM Projects WHERE ID = ?", p.ID)
	if err != nil {
		return err
	}

	return nil
}

func (p *Project) Save() error {
	_, err := db.NamedExec(`
	INSERT INTO Projects (ID, Name, Code, Description, StartDate, PlanEndDate, EndDate, User_ID, User_ID)
		VALUES (UUID(), :Name, Code:, :Description, :StartDate, :PlanEndDate, :EndDate, :User_ID, :User_ID)`, p)

	return err
}
