package dao

import (
	e "github.com/huangxinchun/hxcgo/admin/app/err"
	"github.com/huangxinchun/hxcgo/admin/app/model"
	"github.com/huangxinchun/hxcgo/admin/core/db"
)

type AdminRole struct {
	adminDAO *Admin
	roleDAO  *Role
}

func NewAdminRole() *AdminRole {
	return &AdminRole{
		adminDAO: NewAdmin(),
		roleDAO:  NewRole(),
	}
}

func (ar *AdminRole) FindFirst(conditions interface{}) (*model.AdminRole, error) {
	m := &model.AdminRole{}
	err := m.FindFirst(conditions)
	return m, err
}

func (ar *AdminRole) Associate(adminID uint, roleID uint) (bool, error) {
	if roleID < 1 || adminID < 1 {
		return false, e.EParam
	}

	m := &model.AdminRole{AdminID: adminID, RoleID: roleID}

	adminRole, err := ar.FindFirst(m)
	if err != nil {
		_, err = ar.roleDAO.FindByID(roleID)
		if err != nil {
			return false, e.EParam
		}

		_, err = ar.adminDAO.FindByID(adminID)
		if err != nil {
			return false, e.EParam
		}

		err = m.Save()
		if err != nil {
			return false, err
		}
		return true, nil
	}

	return adminRole.Delete(adminRole.ID), nil
}

func (ar *AdminRole) FindRoleIDsByAdminID(adminID uint) ([]uint, error) {
	if adminID < 1 {
		return nil, e.EParam
	}

	var adminRoles []*model.AdminRole

	/*var adminRolesTableName = gorm.DefaultTableNameHandler(db.Client(), "admin_roles")
	var rolesTableName = gorm.DefaultTableNameHandler(db.Client(), "roles")
	err := db.Client().Joins(
		fmt.Sprintf("JOIN %s ON %s.id=%s.role_id", rolesTableName, rolesTableName, adminRolesTableName)).Where(
		fmt.Sprintf("%s.admin_id = ? AND %s.state = ?", adminRolesTableName, rolesTableName), adminID, model.StateActive).Find(&adminRoles).Error*/
	err := db.Client().Where(&model.AdminRole{AdminID: adminID}).Find(&adminRoles).Error
	if err != nil {
		return nil, err
	}

	var ids []uint
	for _, v := range adminRoles {
		ids = append(ids, v.RoleID)
	}

	return ids, nil
}
