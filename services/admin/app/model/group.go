package model

import (
	"github.com/huangxinchun/hxcgo/services/admin/core/db"
	"time"
)

type Group struct {
	ID        uint `gorm:"primary_key"`
	Name      string
	State     uint8
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (g *Group) FindByID(id uint) error {
	return db.Client().First(g, id).Error
}

func (g *Group) Save() error {
	return db.Client().Save(g).Error
}

func (g *Group) Delete(id uint) bool {
	g.ID = id
	if db.Client().Delete(g).RowsAffected > 0 {
		return true
	}

	return false
}
