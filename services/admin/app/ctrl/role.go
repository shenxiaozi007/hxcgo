package ctrl

import (
	"github.com/huangxinchun/hxcgo/services/admin/app/dao"
	e "github.com/huangxinchun/hxcgo/services/admin/app/err"
	"github.com/huangxinchun/hxcgo/services/admin/app/model"
	"github.com/huangxinchun/hxcgo/services/admin/app/proto/roleproto"
	"math"
)

type Role struct {
	roleDAO          *dao.Role
	rolePrivilegeDAO *dao.RolePrivilege
}

func NewRole() *Role {
	return &Role{
		roleDAO:          dao.NewRole(),
		rolePrivilegeDAO: dao.NewRolePrivilege(),
	}
}

func (r *Role) FindByID(id uint, resp *roleproto.Role) error {
	if id < 1 {
		return e.EParamInvalidID
	}

	role, err := r.roleDAO.FindByID(id)
	if err != nil {
		return err
	}

	r.copy(resp, role)
	return nil
}

func (r *Role) Add(req *roleproto.Role, resp *roleproto.Role) error {
	if req.Name == "" {
		return e.EParamInvalidName
	}

	role, err := r.roleDAO.Add(req)
	if err != nil {
		return err
	}

	r.copy(resp, role)
	return nil
}

func (r *Role) Update(req *roleproto.Role, resp *roleproto.Role) error {
	if req.Name == "" {
		return e.EParamInvalidName
	}

	if req.ID < 1 {
		return e.EParamInvalidID
	}

	role, err := r.roleDAO.Update(req)
	if err != nil {
		return err
	}

	r.copy(resp, role)
	return nil
}

func (r *Role) copy(dst *roleproto.Role, src *model.Role) {
	dst.ID = src.ID
	dst.Name = src.Name
	dst.State = src.State
	dst.CreatedAt = src.CreatedAt
	dst.UpdatedAt = src.UpdatedAt
	dst.DeletedAt = src.DeletedAt
}

func (r *Role) Query(req *roleproto.QueryReq, resp *roleproto.QueryResp) error {
	if req.Page < 1 {
		return e.EParamInvalidPage
	}
	if req.Limit < 1 {
		return e.EParamInvalidLimit
	}

	conditions := map[string]interface{}{}
	if req.State >= 0 {
		conditions["state"] = uint16(req.State)
	}

	count, err := r.roleDAO.Count(conditions)
	if err != nil {
		return err
	}

	resp.Page = req.Page
	resp.Limit = req.Limit
	resp.Count = count
	if count == 0 {
		return nil
	}

	roles, err := r.roleDAO.Roles(conditions, req.Page, req.Limit)
	if err != nil {
		return err
	}

	resp.TotalPage = uint(math.Ceil(float64(resp.Count) / float64(resp.Limit)))

	for _, role := range roles {
		roleResp := &roleproto.Role{}
		r.copy(roleResp, role)

		resp.Roles = append(resp.Roles, roleResp)
	}

	return nil
}

func (r *Role) Delete(id uint, _ *struct{}) error {
	if id < 1 {
		return e.EParamInvalidID
	}

	success, err := r.roleDAO.Delete(id)
	if err != nil {
		return err
	}

	if !success {
		return e.EDeleteFailed
	}
	return nil
}

func (r *Role) AssociatePrivilege(req roleproto.AssociateReq, _ *struct{}) error {
	if req.ID < 1 || req.PrivilegeID < 1 {
		return e.EParam
	}

	success, err := r.rolePrivilegeDAO.Associate(req.ID, req.PrivilegeID)
	if err != nil {
		return err
	}
	if !success {
		return e.EAssociateRolePrivilegeFailed
	}
	return nil
}

func (r *Role) PrivilegeIDs(id uint, resp *roleproto.RolePrivilegeResp) error {
	if id < 1 {
		return e.EParamInvalidID
	}

	privilegeIDs, err := r.rolePrivilegeDAO.FindPrivilegeIDsByRoleID(id)
	if err != nil {
		return err
	}

	resp.PrivilegeIDs = append(resp.PrivilegeIDs, privilegeIDs...)
	return nil
}
