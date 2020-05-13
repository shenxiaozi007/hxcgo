package dao

import (
	e "calendar/app/err"
	"calendar/app/model"
	"calendar/app/proto/userproto"
	"calendar/core/db"

	"github.com/jinzhu/gorm"
)

type User struct {
}

func NewUser() *User {
	return &User{}
}

func (u *User) FindByID(id uint) (*model.User, error) {
	if id < 1 {
		return nil, e.EParamInvalidID
	}

	user := &model.User{}
	err := user.FindByID(id)
	if err != nil {
		return nil, e.DBError(err)
	}

	return user, nil
}

func (u *User) FindByOpenID(openID string) (*model.User, error) {
	if openID == "" {
		return nil, e.EParamInvalidOpenID
	}
	user := &model.User{}
	err := user.FindByOpenID(openID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, e.ENotFound
		}

		return nil, e.EDatabase
	}
	return user, nil
}

func (u *User) copy(dst *model.User, src *userproto.User) {
	dst.ID = src.ID
	if src.OpenID != "" {
		dst.OpenID = src.OpenID
	}
	if src.UnionID != "" {
		dst.UnionID = src.UnionID
	}

	if src.NickName != "" {
		dst.NickName = src.NickName
	}
	if src.Avatar != "" {
		dst.Avatar = src.Avatar
	}

	dst.State = src.State
	dst.Gender = src.Gender

	if src.Country != "" {
		dst.Country = src.Country
	}

	if src.Province != "" {
		dst.Province = src.Province
	}

	if src.City != "" {
		dst.City = src.City
	}

}

func (u *User) Add(req *userproto.User) (*model.User, error) {
	if req.ID != 0 || req.OpenID == "" {
		return nil, e.EParam
	}

	user := &model.User{}
	u.copy(user, req)

	err := user.Save()
	return user, err
}

func (u *User) Update(req *userproto.User) (*model.User, error) {
	if req.ID < 1 {
		return nil, e.EParamInvalidID
	}

	user, err := u.FindByID(req.ID)
	if err != nil {
		return nil, e.EParamInvalidID
	}

	u.copy(user, req)

	err = user.Save()
	return user, err
}

func (u *User) DBConditions(conditions map[string]interface{}) *gorm.DB {
	client := db.Client()
	if name, ok := conditions["name"]; ok {
		client = client.Where("name = ?", name)
	}
	if nickName, ok := conditions["nickName"]; ok {
		client = client.Where("nickName = ?", nickName)
	}
	if regBeginTime, ok := conditions["regBeginTime"]; ok {
		client = client.Where("created_at >= ?", regBeginTime)
	}
	if regEndTime, ok := conditions["regEndTime"]; ok {
		client = client.Where("created_at <= ?", regEndTime)
	}

	return client
}

func (u *User) Users(conditions map[string]interface{}, page uint, limit uint) ([]*model.User, error) {
	if page < 1 {
		page = 1
	}

	offset := (page - 1) * limit

	var users []*model.User
	err := u.DBConditions(conditions).Offset(offset).Limit(limit).Find(&users).Error

	return users, err
}

func (u *User) Count(conditions map[string]interface{}) (uint, error) {
	var countRecords uint
	err := u.DBConditions(conditions).Model(&model.User{}).Count(&countRecords).Error
	return countRecords, err
}

func (u *User) Delete(id uint) (bool, error) {
	if id < 1 {
		return false, e.EParamInvalidID
	}

	user := &model.User{}
	return user.Delete(id), nil
}
