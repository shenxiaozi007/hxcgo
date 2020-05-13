package dao

import (
	e "github.com/huangxinchun/hxcgo/admin/app/err"
	"github.com/huangxinchun/hxcgo/admin/app/model"
	"github.com/huangxinchun/hxcgo/admin/core/db"

	"github.com/jinzhu/gorm"
)

type RolePrivilege struct {
	roleDAO      *Role
	privilegeDAO *Privilege
}

func NewRolePrivilege() *RolePrivilege {
	return &RolePrivilege{
		roleDAO:      NewRole(),
		privilegeDAO: NewPrivilege(),
	}
}

func (rp *RolePrivilege) FindFirst(conditions interface{}) (*model.RolePrivilege, error) {
	m := &model.RolePrivilege{}
	err := m.FindFirst(conditions)
	if err != nil {
		return nil, e.EDatabase
	}
	return m, nil
}

func (rp *RolePrivilege) Associate(roleID uint, privilegeID uint) (bool, error) {
	if roleID < 1 || privilegeID < 1 {
		return false, e.EParam
	}

	m := &model.RolePrivilege{RoleID: roleID, PrivilegeID: privilegeID}

	rolePrivilege, err := rp.FindFirst(m)
	if err != nil {
		_, err = rp.roleDAO.FindByID(roleID)
		if err != nil {
			return false, e.EParam
		}

		_, err = rp.privilegeDAO.FindByID(privilegeID)
		if err != nil {
			return false, e.EParam
		}

		err = m.Save()
		if err != nil {
			return false, err
		}
		return true, nil
	}

	return rolePrivilege.Delete(rolePrivilege.ID), nil
}

func (rp *RolePrivilege) FindRoleIDsByPrivilegeID(privilegeID uint) ([]uint, error) {
	if privilegeID < 1 {
		return nil, e.EParam
	}

	var rolePrivileges []*model.RolePrivilege
	err := db.Client().Where(&model.RolePrivilege{PrivilegeID: privilegeID}).Find(&rolePrivileges).Error
	if err != nil {
		return nil, err
	}

	var ids []uint
	for _, v := range rolePrivileges {
		ids = append(ids, v.RoleID)
	}

	return ids, nil
}

func (rp *RolePrivilege) FindPrivilegeIDsByRoleID(roleID uint) ([]uint, error) {
	if roleID < 1 {
		return nil, e.EParam
	}

	var rolePrivileges []*model.RolePrivilege
	err := db.Client().Where(&model.RolePrivilege{RoleID: roleID}).Find(&rolePrivileges).Error

	if err != nil {
		return nil, err
	}

	var ids []uint
	for _, v := range rolePrivileges {
		ids = append(ids, v.PrivilegeID)
	}

	return ids, nil
}

func (rp *RolePrivilege) DeleteByRoleID(roleID uint) (bool, error) {
	err := db.Client().Where("role_id = ?",roleID).Delete(&model.RolePrivilege{}).Error

	if err == nil || err == gorm.ErrRecordNotFound {
		return true, nil
	}
	return false, err
}

func (rp *RolePrivilege) DeleteByPrivilegeID(privilegeID uint) (bool, error) {
	err := db.Client().Where("privilege_id = ?",privilegeID).Delete(&model.RolePrivilege{}).Error

	if err == nil || err == gorm.ErrRecordNotFound {
		return true, nil
	}
	return false, err
}
