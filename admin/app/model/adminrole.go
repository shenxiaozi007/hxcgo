package model

import "github.com/huangxinchun/hxcgo/admin/core/db"

type AdminRole struct {
	ID      uint `gorm:"primary_key"`
	AdminID uint
	RoleID  uint
}

func (ar *AdminRole) FindFirst(condition interface{}) error {
	return db.Client().Where(condition).First(ar).Error
}

func (ar *AdminRole) Save() error {
	return db.Client().Save(ar).Error
}

func (ar *AdminRole) Delete(id uint) bool {
	ar.ID = id
	if db.Client().Delete(ar).RowsAffected > 0 {
		return true
	}

	return false
}
