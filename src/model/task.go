package model

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"

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
	q := t.Select().Where("ID = ?", id)
	return t.QueryRow(q)
}

func (t *Task) Select() sq.SelectBuilder {
	return sq.Select("*").From("Tasks")
}

func (t *Task) QueryRow(q sq.SelectBuilder) error {
	row := q.RunWith(db.DB).QueryRow()
	return t.scan(row)
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
		q := sq.Insert("Tasks").
			Columns("Name", "Description", "Code", "StartDate", "PlanEndDate", "EndDate", "Project_ID", "User_ID").
			Values(t.name, t.description, t.code, t.startDate.String(), t.planEndDate.Val(), t.endDate.Val(), projectID, userID)

		res, err := q.RunWith(db.DB).Exec()
		if err != nil {
			return err
		}

		id, err := res.LastInsertId()
		if err != nil {
			return err
		}

		t.id = id
	} else {
		q := sq.Update("Tasks").
			Set("Code", t.code).
			Set("Name", t.name).
			Set("Description", t.description).
			Set("StartDate", t.startDate.String()).
			Set("PlanEndDate", t.planEndDate.Val()).
			Set("EndDate", t.endDate.Val()).
			Set("Project_ID", projectID).
			Set("User_ID", userID).
			Where("ID = ?", t.id)

		_, err := q.RunWith(db.DB).Exec()
		if err != nil {
			return err
		}
	}

	return nil
}

func (t *Task) Delete() error {
	q := sq.Delete("Tasks").Where("ID = ?", t.id)
	_, err := q.RunWith(db.DB).Exec()
	return err
}

func (t *Task) scan(row sq.RowScanner) error {
	var projectID, userID sql.NullInt64

	err := row.Scan(&t.id, &t.name, &t.description, &t.code, &t.startDate, &t.planEndDate, &t.endDate, &projectID, &userID)
	if err != nil {
		return err
	}

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

// -----------------------------------------------------------------------------
// --- GET & SET
// -----------------------------------------------------------------------------

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

// -----------------------------------------------------------------------------
// --- helper functions
// -----------------------------------------------------------------------------

func SelectTasks() sq.SelectBuilder {
	return sq.Select("*").From("Tasks")
}

func QueryTasks(q sq.SelectBuilder) ([]*Task, error) {
	tasks := make([]*Task, 0)

	rows, err := q.RunWith(db.DB).Query()
	if err != nil {
		return nil, err
	}

	var t *Task

	for rows.Next() {
		t = NewTask()
		t.scan(rows)

		tasks = append(tasks, t)
	}

	return tasks, nil
}
