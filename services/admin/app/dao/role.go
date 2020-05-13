package dao

import (
	e "github.com/huangxinchun/hxcgo/services/admin/app/err"
	"github.com/huangxinchun/hxcgo/services/admin/app/model"
	"github.com/huangxinchun/hxcgo/services/admin/app/proto/roleproto"
	"github.com/huangxinchun/hxcgo/services/admin/core/db"
)

type Role struct {
}

func NewRole() *Role {
	return &Role{}
}

func (r *Role) FindByID(id uint) (*model.Role, error) {
	if id < 1 {
		return nil, e.EParamInvalidID
	}

	role := &model.Role{}
	err := role.FindByID(id)
	if err != nil {
		return nil, e.EDatabase
	}
	return role, nil
}

func (r *Role) Update(req *roleproto.Role) (*model.Role, error) {
	if req.ID < 1 || req.Name == "" {
		return nil, e.EParam
	}
	role := &model.Role{
		ID:    req.ID,
		Name:  req.Name,
		State: req.State,
	}
	err := role.Save()

	return role, err
}

func (r *Role) Add(req *roleproto.Role) (*model.Role, error) {
	if req.Name == "" {
		return nil, e.EParamInvalidName
	}
	role := &model.Role{
		Name:  req.Name,
		State: req.State,
	}
	err := role.Save()

	return role, err
}

func (r *Role) Roles(conditions map[string]interface{}, page uint, limit uint) ([]*model.Role, error) {
	if page < 1 {
		page = 1
	}

	offset := (page - 1) * limit

	var roles []*model.Role
	err := db.Client().Where(conditions).Offset(offset).Limit(limit).Find(&roles).Error

	return roles, err
}

func (r *Role) Count(conditions map[string]interface{}) (uint, error) {
	var countRecords uint
	err := db.Client().Model(&model.Role{}).Where(conditions).Count(&countRecords).Error
	return countRecords, err
}

func (r *Role) Delete(id uint) (bool, error) {
	if id < 1 {
		return false, e.EParamInvalidID
	}

	role := &model.Role{}
	isSuccess := role.Delete(id)

	rolePrivilegeDAO := NewRolePrivilege()
	if isSuccess {
		rolePrivilegeDAO.DeleteByRoleID(id)
	}
	return isSuccess, nil
}
