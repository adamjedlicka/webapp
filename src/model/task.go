package model

import (
	"database/sql"

	"github.com/adamjedlicka/webapp/src/shared/db"
)

type Task struct {
	ID           string         `db:"ID"`
	Name         string         `db:"Name"`
	Description  sql.NullString `db:"Description"`
	StartDate    Date           `db:"StartDate"`
	PlanEndDate  NullDate       `db:"PlanEndDate"`
	EndDate      NullDate       `db:"EndDate"`
	MaintainerID string         `db:"User_ID_Maintainer"`
	WorkerID     sql.NullString `db:"User_ID_Worker"`
	ProjectID    string         `db:"Project_ID"`
}

func (t *Task) Fill() error {
	return db.Get("SELECT * FROM Tasks WHERE ID = ?", t.ID)
}

func (t *Task) Save() error {
	_, err := db.NamedExec(`
	INSERT INTO Tasks (ID, Name, Description, StartDate, PlanEndDate, EndDate, User_ID_Maintainer, User_ID_Worker, Project_ID)
		VALUES (UUID(), Name, Description, :StartDate, :PlanEndDate, :EndDate, :User_ID_Maintainer, :User_ID_Worker, :Project_ID)`, t)

	return err
}
