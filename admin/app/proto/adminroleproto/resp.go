package adminroleproto

import (
	"admin/core/rpc"
	"time"
)

type Role struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Name      string
	State     uint8
}

func FindByID(id uint) (*Role, error) {
	role := &Role{}
	err := rpc.Service("admin").Call("Role.FindByID", id, role)
	return role, err
}

func (r *Role) Create() error {
	return rpc.Service("admin").Call("Role.Add", r, r)
}

func (r *Role) Update() error {
	return rpc.Service("admin").Call("Role.Update", r, r)
}

func (r *Role) Save() error {
	if r.ID > 0 {
		return r.Update()
	}

	return r.Create()
}

type QueryResp struct {
	Roles     []*Role
	Page      uint
	Limit     uint
	Count     uint
	TotalPage uint
}

func Query(req *QueryReq) (*QueryResp, error) {
	resp := &QueryResp{}
	err := rpc.Service("admin").Call("Role.Query", req, resp)
	return resp, err
}

func Delete(id uint) error {
	return rpc.Service("admin").Call("Role.Delete", id, nil)
}

func AssociatePrivilege(req *AssociateReq) error {
	return rpc.Service("admin").Call("Role.AssociatePrivilege", req, nil)
}

type RolePrivilegeResp struct {
	PrivilegeIDs []uint
}

func PrivilegeIDs(id uint) ([]uint, error) {
	resp := &RolePrivilegeResp{}
	err := rpc.Service("admin").Call("Role.PrivilegeIDs", id, resp)
	return resp.PrivilegeIDs, err
}
