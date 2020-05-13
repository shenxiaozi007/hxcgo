package adminservice

import (
	e "github.com/huangxinchun/hxcgo/admin/app/err"
	"github.com/huangxinchun/hxcgo/admin/app/proto/adminroleproto"
)

type Role struct {
}

func NewRole() *Role {
	return &Role{}
}

func (r *Role) FindByID(id uint) (*adminroleproto.Role, error) {
	if id < 1 {
		return nil, e.EInvalidID
	}

	resp, err := adminroleproto.FindByID(id)

	return resp, err
}

func (r *Role) Add(role *adminroleproto.Role) error {
	if role.Name == "" {
		return e.EInvalidName
	}

	return role.Save()
}

func (r *Role) Update(role *adminroleproto.Role) error {
	if role.ID < 1 {
		return e.EInvalidID
	}
	if role.Name == "" {
		return e.EInvalidName
	}

	return role.Save()
}

func (r *Role) Query(req *adminroleproto.QueryReq) (*adminroleproto.QueryResp, error) {
	if req.Page < 1 {
		req.Page = 1
	}

	if req.Limit < 1 {
		req.Limit = 20
	}

	resp, err := adminroleproto.Query(req)

	return resp, err
}

func (r *Role) Delete(id uint) error {
	if id < 1 {
		return e.EInvalidID
	}

	return adminroleproto.Delete(id)
}

func (r *Role) Actives() (map[uint]*adminroleproto.Role, error) {
	req := &adminroleproto.QueryReq{
		State: 0,
		Page:  1,
		Limit: 200,
	}

	resp, err := r.Query(req)
	if err != nil {
		return nil, err
	}

	roles := map[uint]*adminroleproto.Role{}
	for _, role := range resp.Roles {
		roles[role.ID] = role
	}

	return roles, err
}

func (r *Role) AssociatePrivilege(id uint, privilegeID uint) error {
	if id < 1 || privilegeID < 1 {
		return e.EParam
	}

	req := &adminroleproto.AssociateReq{
		ID:          id,
		PrivilegeID: privilegeID,
	}
	return adminroleproto.AssociatePrivilege(req)
}

func (r *Role) PrivilegeIDs(id uint) ([]uint, error) {
	if id < 1 {
		return nil, e.EInvalidID
	}

	return adminroleproto.PrivilegeIDs(id)
}

func (r *Role) PrivilegeIDMap(id uint) (map[uint]bool, error) {
	if id < 1 {
		return nil, e.EInvalidID
	}

	privilegeIDs, err := r.PrivilegeIDs(id)
	if err != nil {
		return nil, err
	}
	idMap := map[uint]bool{}
	for _, v := range privilegeIDs {
		idMap[v] = true
	}

	return idMap, err
}
