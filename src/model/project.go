package model

import (
	sq "github.com/Masterminds/squirrel"

	"github.com/adamjedlicka/webapp/src/shared/db"
)

type Project struct {
	id   int64
	name string
	code string
}

func NewProject() *Project {
	p := new(Project)
	p.id = 0

	return p
}

func (p *Project) FindByID(id int64) error {
	q := p.Select().Where("ID = ?", id)
	return p.QueryRow(q)
}

func (p *Project) Select() sq.SelectBuilder {
	return sq.Select("ID, Name, Code").From("Projects")
}

func (p *Project) QueryRow(q sq.SelectBuilder) error {
	row := q.RunWith(db.DB).QueryRow()

	err := row.Scan(&p.id, &p.name, &p.code)

	return err
}

func (p *Project) Save() error {
	if p.id == 0 {
		q := sq.Insert("Users").Columns("Name, Code").Values(p.name, p.code)

		res, err := q.RunWith(db.DB).Exec()
		if err != nil {
			return err
		}

		id, err := res.LastInsertId()
		if err != nil {
			return err
		}

		p.id = id

		return nil
	}

	q := sq.Update("Users").
		Set("Name", p.name).
		Set("Code", p.code)

	_, err := q.RunWith(db.DB).Exec()
	return err

}

func (p Project) ID() int64    { return p.id }
func (p Project) Name() string { return p.name }

func (p *Project) SetName(name string) { p.name = name }
