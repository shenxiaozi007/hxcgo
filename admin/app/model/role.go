package model

import (
	"github.com/huangxinchun/hxcgo/admin/core/db"
	"time"
)

type Role struct {
	ID        uint `gorm:"primary_key"`
	Name      string
	State     uint8
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (r *Role) FindByID(id uint) error {
	return db.Client().First(r, id).Error
}

func (r *Role) Save() error {
	return db.Client().Save(r).Error
}

func (r *Role) Delete(id uint) bool {
	r.ID = id
	if db.Client().Delete(r).RowsAffected > 0 {
		return true
	}

	return false
}
