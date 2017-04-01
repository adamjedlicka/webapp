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

	Maintainer *User
	Worker     *User
	Project    *Project
}

func (t *Task) Fill() error {
	err := db.Get(t, "SELECT * FROM Tasks WHERE ID = ?", t.ID)
	if err != nil {
		return err
	}

	t.Maintainer = &User{ID: t.MaintainerID}
	err = t.Maintainer.Fill()
	if err != nil {
		return err
	}

	t.Project = &Project{ID: t.ProjectID}
	err = t.Project.Fill()
	if err != nil {
		return err
	}

	return nil
}

func (t *Task) Save() error {
	_, err := db.NamedExec(`
	INSERT INTO Tasks (ID, Name, Description, StartDate, PlanEndDate, EndDate, User_ID_Maintainer, User_ID_Worker, Project_ID)
		VALUES (UUID(), Name, Description, :StartDate, :PlanEndDate, :EndDate, :User_ID_Maintainer, :User_ID_Worker, :Project_ID)`, t)

	return err
}
