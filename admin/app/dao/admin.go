package dao

import (
	"github.com/huangxinchun/hxcgo/admin/app/cachekey"
	e "github.com/huangxinchun/hxcgo/admin/app/err"
	"github.com/huangxinchun/hxcgo/admin/app/model"
	"github.com/huangxinchun/hxcgo/admin/app/proto/adminproto"
	"github.com/huangxinchun/hxcgo/admin/core/db"
	"github.com/huangxinchun/hxcgo/admin/core/redis"
	"time"
)

type Admin struct {
	groupDAO *Group
}

func NewAdmin() *Admin {
	return &Admin{
		groupDAO: NewGroup(),
	}
}

func (a *Admin) FindByID(id uint) (*model.Admin, error) {
	if id < 1 {
		return nil, e.EParamInvalidID
	}

	admin := &model.Admin{}
	err := admin.FindByID(id)
	if err != nil {
		return nil, e.EDatabase
	}

	return admin, nil
}

func (a *Admin) FindByName(name string) (*model.Admin, error) {
	if name == "" {
		return nil, e.EParamInvalidName
	}

	admin := &model.Admin{}
	err := admin.FindByName(name)
	if err != nil {
		return nil, e.EDatabase
	}
	return admin, nil
}

func (a *Admin) FindByEmail(email string) (*model.Admin, error) {
	if email == "" {
		return nil, e.EParamInvalidEmail
	}
	admin := &model.Admin{}
	err := admin.FindByEmail(email)
	if err != nil {
		return nil, e.EDatabase
	}
	return admin, nil
}

func (a *Admin) FindByMobile(mobile string) (*model.Admin, error) {
	if mobile == "" {
		return nil, e.EParamInvalidMobile
	}
	admin := &model.Admin{}
	err := admin.FindByMobile(mobile)
	if err != nil {
		return nil, e.EDatabase
	}
	return admin, nil
}

func (a *Admin) verifyReq(req *adminproto.Admin) error {
	if req.GroupID < 1 {
		return e.EParamInvalidGroupID
	}

	if len(req.Name) < 5 {
		return e.EParamInvalidName
	}

	//检查角色是否合法
	group, err := a.groupDAO.FindByID(req.GroupID)
	if err != nil || group.ID != req.GroupID || group.State != model.StateActive {
		return e.EParamInvalidGroupID
	}

	admin, err := a.FindByName(req.Name)
	if err == nil && admin.ID != req.ID {
		return e.EParamInvalidName
	}

	if req.Mobile != "" {
		admin, err = a.FindByMobile(req.Mobile)
		if err == nil && admin.ID != req.ID {
			return e.EParamInvalidMobile
		}

		//手机唯一
		mobileKey := cachekey.AdminUniqueMobile(req.Mobile)
		err = redis.Client().SetNX(mobileKey, 1, 2*time.Second).Err()
		if err != nil {
			return e.EParamInvalidMobile
		}
		defer redis.Client().Expire(mobileKey, 0)
	}
	if req.Email != "" {
		admin, err = a.FindByEmail(req.Email)
		if err == nil && admin.ID != req.ID {
			return e.EParamInvalidEmail
		}

		//邮箱唯一
		emailKey := cachekey.AdminUniqueEmail(req.Email)
		err = redis.Client().SetNX(emailKey, 1, 2*time.Second).Err()
		if err != nil {
			return e.EParamInvalidEmail
		}
		defer redis.Client().Expire(emailKey, 0)
	}

	return nil
}

func (a *Admin) copy(dst *model.Admin, src *adminproto.Admin) {
	dst.ID = src.ID
	dst.Name = src.Name
	dst.State = src.State
	dst.Mobile = src.Mobile
	dst.Email = src.Email
	dst.GroupID = src.GroupID
	dst.UpdatedAt = src.UpdatedAt
	dst.LoginAt = src.LoginAt
	dst.LoginIP = src.LoginIP

	if len(src.Password) > 0 {
		dst.Password = src.Password
	}
	if len(src.Avatar) > 0 {
		dst.Avatar = src.Avatar
	}
}

func (a *Admin) Add(req *adminproto.Admin) (*model.Admin, error) {
	if req.ID != 0 {
		return nil, e.EParam
	}
	if len(req.Password) < 32 {
		return nil, e.EParamInvalidPassword
	}

	err := a.verifyReq(req)
	if err != nil {
		return nil, err
	}

	admin := &model.Admin{}
	a.copy(admin, req)

	err = admin.Save()
	return admin, err
}

func (a *Admin) Update(req *adminproto.Admin) (*model.Admin, error) {
	if req.ID < 1 {
		return nil, e.EParamInvalidID
	}

	admin, err := a.FindByID(req.ID)
	if err != nil {
		return nil, e.EParamInvalidID
	}
	err = a.verifyReq(req)
	if err != nil {
		return nil, err
	}
	a.copy(admin, req)

	err = admin.Save()
	return admin, err
}

func (a *Admin) UpdateLoginTimeAndIP(req *adminproto.Admin) (*model.Admin, error) {
	if req.ID < 1 {
		return nil, e.EParamInvalidID
	}

	admin, err := a.FindByID(req.ID)
	if err != nil {
		return nil, e.EParamInvalidID
	}

	admin.LoginAt = req.LoginAt
	admin.LoginIP = req.LoginIP

	err = admin.Save()
	return admin, err
}

func (a *Admin) Admins(conditions map[string]interface{}, page uint, limit uint) ([]*model.Admin, error) {
	if page < 1 {
		page = 1
	}

	offset := (page - 1) * limit

	var admins []*model.Admin
	err := db.Client().Where(conditions).Order("id desc").Offset(offset).Limit(limit).Find(&admins).Error

	return admins, err
}

func (a *Admin) Count(conditions map[string]interface{}) (uint, error) {
	var countRecords uint
	err := db.Client().Model(&model.Admin{}).Where(conditions).Count(&countRecords).Error
	return countRecords, err
}

func (a *Admin) Delete(id uint) (bool, error) {
	if id < 1 {
		return false, e.EParamInvalidID
	}

	admin := &model.Admin{}
	return admin.Delete(id), nil
}
