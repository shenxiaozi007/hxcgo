package model

import "github.com/huangxinchun/hxcgo/admin/core/db"

type RolePrivilege struct {
	ID          uint `gorm:"primary_key"`
	RoleID      uint
	PrivilegeID uint
}

func (rp *RolePrivilege) FindFirst(condition interface{}) error {
	return db.Client().Where(condition).First(rp).Error
}

func (rp *RolePrivilege) Save() error {
	return db.Client().Save(rp).Error
}

func (rp *RolePrivilege) Delete(id uint) bool {
	rp.ID = id
	if db.Client().Delete(rp).RowsAffected > 0 {
		return true
	}

	return false
}
