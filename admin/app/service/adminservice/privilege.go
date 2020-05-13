package adminservice

import (
	e "github.com/huangxinchun/hxcgo/admin/app/err"
	"github.com/huangxinchun/hxcgo/admin/app/proto/adminprivilegeproto"
)

type Privilege struct {
}

func NewPrivilege() *Privilege {
	return &Privilege{}
}

func (p *Privilege) FindByID(id uint) (*adminprivilegeproto.Privilege, error) {
	if id < 1 {
		return nil, e.EInvalidID
	}

	resp, err := adminprivilegeproto.FindByID(id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *Privilege) verifyReq(req *adminprivilegeproto.Privilege) error {
	if req.Name == "" {
		return e.EInvalidName
	}

	if req.PID > 0 {
		_, err := p.FindByID(req.PID)
		if err != nil {
			return e.EInvalidPID
		}
	}

	return nil
}

func (p *Privilege) Add(privilege *adminprivilegeproto.Privilege) error {
	privilege.ID = 0
	err := p.verifyReq(privilege)
	if err != nil {
		return err
	}

	return privilege.Save()
}

func (p *Privilege) Update(privilege *adminprivilegeproto.Privilege) error {
	if privilege.ID < 1 {
		return e.EParam
	}
	err := p.verifyReq(privilege)
	if err != nil {
		return err
	}

	return privilege.Save()
}

func (p *Privilege) FindAll() ([]*adminprivilegeproto.Privilege, error) {
	resp, err := adminprivilegeproto.FindAll()
	return resp, err
}

func (p *Privilege) Delete(id uint) error {
	if id < 1 {
		return e.EInvalidID
	}

	return adminprivilegeproto.Delete(id)
}

func (p *Privilege) RoleIDs(id uint) ([]uint, error) {
	if id < 1 {
		return nil, e.EInvalidID
	}

	return adminprivilegeproto.RoleIDs(id)
}

func (p *Privilege) Menus() ([]*adminprivilegeproto.Privilege, error) {
	privileges, err := p.FindAll()
	if err != nil {
		return nil, err
	}

	var menus []*adminprivilegeproto.Privilege
	for _, v := range privileges {
		if v.IsMenu == 1 {
			menus = append(menus, v)
		}
	}

	return menus, nil
}
