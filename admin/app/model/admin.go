package model

import (
	"time"

	"github.com/huangxinchun/hxcgo/admin/core/db"
)

type Admin struct {
	ID        uint `gorm:"primary_key"`
	Name      string
	Email     string
	Mobile    string
	Password  string
	Avatar    string
	State     uint8
	GroupID   uint
	LoginAt   *time.Time
	LoginIP   int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (a *Admin) FindByID(id uint) error {
	return db.Client().First(a, id).Error
}

func (a *Admin) FindByName(name string) error {
	return db.Client().Where(&Admin{Name: name}).First(a).Error
}

func (a *Admin) FindByEmail(email string) error {
	return db.Client().Where(&Admin{Email: email}).First(a).Error
}

func (a *Admin) FindByMobile(mobile string) error {
	return db.Client().Where(&Admin{Mobile: mobile}).First(a).Error
}

func (a *Admin) Save() error {
	return db.Client().Save(a).Error
}

func (a *Admin) Delete(id uint) bool {
	a.ID = id
	if db.Client().Delete(a).RowsAffected > 0 {
		return true
	}

	return false
}
