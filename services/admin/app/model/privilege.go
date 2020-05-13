package model

import (
	"github.com/huangxinchun/hxcgo/services/admin/core/db"
	"time"
)

type Privilege struct {
	ID         uint `gorm:"primary_key"`
	PID        uint `gorm:"column:pid"`
	Name       string
	Root       string `sql:"index"`
	Icon       string
	URIRule    string
	IsMenu     uint16
	SortOrder  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `sql:"index"`
}

func (p *Privilege) FindByID(id uint) error {
	return db.Client().First(p, id).Error
}

func (p *Privilege) Save() error {
	return db.Client().Save(p).Error
}

func (p *Privilege) Delete(id uint) bool {
	p.ID = id
	if db.Client().Delete(p).RowsAffected > 0 {
		return true
	}

	return false
}
