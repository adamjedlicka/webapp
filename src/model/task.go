package model

import (
	"database/sql"

	"github.com/adamjedlicka/webapp/src/shared/db"
)

type Task struct {
	ID           UUID           `db:"ID"`
	Name         string         `db:"Name"`
	Description  sql.NullString `db:"Description"`
	StartDate    Date           `db:"StartDate"`
	PlanEndDate  NullDate       `db:"PlanEndDate"`
	EndDate      NullDate       `db:"EndDate"`
	MaintainerID UUID           `db:"User_ID_Maintainer"`
	WorkerID     NullUUID       `db:"User_ID_Worker"`
	ProjectID    UUID           `db:"Project_ID"`

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
	if t.ID.String() == "" {
		_, err := db.NamedExec(`
		INSERT INTO Tasks (ID, Name, Description, StartDate, PlanEndDate, EndDate, User_ID_Maintainer, User_ID_Worker, Project_ID)
			VALUES (UUID(), :Name, :Description, :StartDate, :PlanEndDate, :EndDate, :User_ID_Maintainer, :User_ID_Worker, :Project_ID)`, t)

		return err
	}

	_, err := db.NamedExec(`UPDATE Tasks SET
		Name = :Name,
		Description = :Description,
		StartDate = :StartDate,
		PlanEndDate = :PlanEndDate,
		EndDate = :EndDate,
		User_ID_Maintainer = :User_ID_Maintainer,
		User_ID_Worker = :User_ID_Worker,
		Project_ID = :Project_ID
		WHERE ID = :ID`, t)

	return err
}

func (t *Task) Delete() error {
	_, err := db.NamedExec("DELETE FROM Tasks WHERE ID = :ID", t)
	return err
}
