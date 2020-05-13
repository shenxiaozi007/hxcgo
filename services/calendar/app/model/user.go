package model

import (
	"calendar/core/db"
	"time"
)

type User struct {
	ID           uint `gorm:"primary_key"`
	OpenID       string
	UnionID      string
	NickName     string
	Avatar       string
	State        uint8
	Gender       uint8
	Country      string
	Province     string
	City         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time `sql:"index"`
}

func (u *User) FindByID(id uint) error {
	return db.Client().First(u, id).Error
}

func (u *User) FindByOpenID(openID string) error {
	return db.Client().Where(&User{OpenID: openID}).First(u).Error

}

func (u *User) Save() error {
	return db.Client().Save(u).Error
}

func (u *User) Delete(id uint) bool {
	u.ID = id
	if db.Client().Delete(u).RowsAffected > 0 {
		return true
	}

	return false
}
