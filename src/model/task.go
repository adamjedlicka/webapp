package model

import (
	"database/sql"
	"log"

	"github.com/adamjedlicka/webapp/src/shared/db"
)

type Task struct {
	id          int64
	name        string
	description sql.NullString
	code        string
	startDate   Date
	planEndDate NullDate
	endDate     NullDate
	project     *Project
	user        *User
}

func NewTask() *Task {
	t := new(Task)
	t.id = 0

	return t
}

func (t *Task) FindByID(id int64) error {
	var projectID sql.NullInt64
	var userID sql.NullInt64

	db.QueryRow(`SELECT ID, Name, Description, Code, StartDate, PlanEndDate, EndDate, Project_ID, User_ID FROM Tasks WHERE ID = ?`, id).
		Scan(&t.id, &t.name, &t.description, &t.code, &t.startDate, &t.planEndDate, &t.endDate, &projectID, &userID)

	if projectID.Valid {
		t.project = NewProject()
		t.project.FindByID(projectID.Int64)
	}

	if userID.Valid {
		t.user = NewUser()
		t.user.FindByID(userID.Int64)
	}

	return nil
}

func (t *Task) Save() error {
	var projectID interface{} = nil
	if t.project != nil {
		projectID = t.project.id
	}

	var userID interface{} = nil
	if t.user != nil {
		userID = t.user.id
	}

	if t.id == 0 {
		res, err := db.Exec("INSERT INTO Tasks (Name, Description, Code, StartDate, PlanEndDate, EndDate, Project_ID, User_ID) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
			t.name, t.description, t.code, t.startDate.String(), t.planEndDate.Val(), t.endDate.Val(), projectID, userID)
		if err != nil {
			return err
		}

		id, err := res.LastInsertId()
		if err != nil {
			return err
		}

		t.id = id
	} else {
		_, err := db.Exec("UPDATE Tasks SET Name=?, Description=?, Code=?, PlanEndDate=?, EndDate=?, Project_ID=?, User_ID=? WHERE ID = ?",
			t.name, t.description, t.code, t.planEndDate.Val(), t.endDate.Val(), projectID, userID, t.id)
		if err != nil {
			return err
		}
	}

	return nil
}

func (t *Task) Delete() error {
	_, err := db.Exec("DELETE FROM Tasks WHERE ID = ?", t.id)
	if err != nil {
		return err
	}

	return nil
}

func (t Task) ID() int64           { return t.id }
func (t Task) Name() string        { return t.name }
func (t Task) Description() string { return t.description.String }
func (t Task) Code() string        { return t.code }
func (t Task) StartDate() Date     { return t.startDate }
func (t Task) PlanEndDate() Date   { return t.planEndDate.Date }
func (t Task) EndDate() Date       { return t.endDate.Date }
func (t Task) Project() *Project   { return t.project }
func (t Task) User() *User         { return t.user }

func (t *Task) SetName(name string)         { t.name = name }
func (t *Task) SetCode(code string)         { t.code = code }
func (t *Task) SetStartDate(startDate Date) { t.startDate = startDate }
func (t *Task) SetDescription(description string) {
	t.description.String = description
	t.description.Valid = description != ""
}
func (t *Task) SetPlanEndDate(planEndDate Date) {
	t.planEndDate.Date = planEndDate
	t.planEndDate.Valid = planEndDate.String() != ""
}
func (t *Task) SetEndDate(endDate Date) {
	t.endDate.Date = endDate
	t.endDate.Valid = endDate.String() != ""
}

func (t *Task) SetProjectID(id int64) error {
	p := NewProject()
	err := p.FindByID(id)
	if err != nil {
		return err
	}

	t.project = p

	return nil
}

func (t *Task) SetUserID(id int64) error {
	u := NewUser()
	err := u.FindByID(id)
	if err != nil {
		return err
	}

	t.user = u

	return nil
}

// ---------- helper functions ----------

func GetTasks() []*Task {
	tasks := make([]*Task, 0)

	res, err := db.Query("SELECT ID FROM Tasks ORDER BY ID")
	if err != nil {
		log.Fatal(err)
	}

	var id int64

	for res.Next() {
		err := res.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}

		t := NewTask()
		t.FindByID(id)

		tasks = append(tasks, t)
	}

	return tasks
}
