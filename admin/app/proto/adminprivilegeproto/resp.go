package adminprivilegeproto

import (
	"admin/core/rpc"
	"time"
)

type Privilege struct {
	ID         uint
	PID        uint
	Name       string
	Icon       string
	URIRule    string
	IsMenu     uint16
	SortOrder  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}

func FindByID(id uint) (*Privilege, error) {
	privilege := &Privilege{}
	err := rpc.Service("admin").Call("Privilege.FindByID", id, privilege)
	return privilege, err
}

func (p *Privilege) Create() error {
	return rpc.Service("admin").Call("Privilege.Add", p, p)
}

func (p *Privilege) Update() error {
	return rpc.Service("admin").Call("Privilege.Update", p, p)
}

func (p *Privilege) Save() error {
	if p.ID > 0 {
		return p.Update()
	}

	return p.Create()
}

type QueryResp struct {
	Privileges []*Privilege
}

func FindAll() ([]*Privilege, error) {
	resp := &QueryResp{}
	err := rpc.Service("admin").Call("Privilege.Query", &struct{}{}, resp)
	return resp.Privileges, err
}

func Delete(id uint) error {
	return rpc.Service("admin").Call("Privilege.Delete", id, nil)
}

type PrivilegeRoleResp struct {
	RoleIDs []uint
}

func RoleIDs(id uint) ([]uint, error) {
	resp := &PrivilegeRoleResp{}
	err := rpc.Service("admin").Call("Privilege.RoleIDs", id, resp)
	return resp.RoleIDs, err
}
