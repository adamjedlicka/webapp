package model

import (
	"log"

	"database/sql"

	"github.com/adamjedlicka/webapp/src/shared/db"
)

type Project struct {
	id          int64
	name        string
	description sql.NullString
	code        string
	userID      int64
}

func NewProject() *Project {
	p := new(Project)
	p.id = 0

	return p
}

func (p *Project) FindByID(id int64) error {
	db.QueryRow("SELECT ID, Name, Description, Code FROM Projects WHERE ID = ?", id).
		Scan(&p.id, &p.name, &p.description, &p.code)

	return nil
}

func (p *Project) Save() error {
	if p.id == 0 {
		res, err := db.Exec("INSERT INTO Projects (Name, Description) VALUES (?, ?)", p.name, p.description)
		if err != nil {
			return err
		}

		id, err := res.LastInsertId()
		if err != nil {
			return err
		}

		p.id = id
	}

	return nil
}

func (p Project) ID() int64           { return p.id }
func (p Project) Name() string        { return p.name }
func (p Project) Description() string { return p.description.String }

func (p *Project) SetName(name string)               { p.name = name }
func (p *Project) SetDescription(description string) { p.description.String = description }

func GetProjects() []*Project {
	projects := make([]*Project, 0)

	res, err := db.Query("SELECT ID, Name, Description FROM Projects ORDER BY ID")
	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {
		p := NewProject()
		res.Scan(&p.id, &p.name, &p.description)
		projects = append(projects, p)
	}

	return projects
}
