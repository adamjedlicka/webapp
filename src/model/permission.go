package model

import "github.com/adamjedlicka/webapp/src/shared/db"

type Permission struct {
	ID                string `db:"ID"`
	Name              string `db:"Name"`
	IsAdmin           bool   `db:"IsAdmin"`
	CanManageUsers    bool   `db:"CanManageUsers"`
	CanManageProjects bool   `db:"CanManageProjects"`
}

func (p *Permission) Fill() error {
	return db.Get(p, "SELECT * FROM Permissions WHERE ID = ?", p.ID)
}

func (p *Permission) FindByID(id string) error {
	return db.Get(p, "SELECT * FROM Permissions WHERE ID LIKE ?", id+"%")
}

func (p *Permission) FindByName(name string) error {
	return db.Get(p, "SELECT * FROM Permissions WHERE Name = ?", name)
}
