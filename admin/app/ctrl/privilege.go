package ctrl

import (
	"github.com/huangxinchun/hxcgo/admin/app/dao"
	e "github.com/huangxinchun/hxcgo/admin/app/err"
	"github.com/huangxinchun/hxcgo/admin/app/model"
	"github.com/huangxinchun/hxcgo/admin/app/proto/privilegeproto"
)

type Privilege struct {
	privilegeDAO     *dao.Privilege
	rolePrivilegeDAO *dao.RolePrivilege
}

func NewPrivilege() *Privilege {
	return &Privilege{
		privilegeDAO:     dao.NewPrivilege(),
		rolePrivilegeDAO: dao.NewRolePrivilege(),
	}
}

func (p *Privilege) FindByID(id uint, resp *privilegeproto.Privilege) error {
	if id < 1 {
		return e.EParamInvalidID
	}

	privilege, err := p.privilegeDAO.FindByID(id)
	if err != nil {
		return err
	}

	p.copy(resp, privilege)
	return nil
}

func (p *Privilege) Add(req *privilegeproto.Privilege, resp *privilegeproto.Privilege) error {
	if req.Name == "" {
		return e.EParamInvalidName
	}

	privilege, err := p.privilegeDAO.Add(req)
	if err != nil {
		return err
	}

	p.copy(resp, privilege)
	return nil
}

func (p *Privilege) Update(req *privilegeproto.Privilege, resp *privilegeproto.Privilege) error {
	if req.Name == "" {
		return e.EParamInvalidName
	}

	if req.ID < 1 {
		return e.EParamInvalidID
	}

	if req.ID == req.PID {
		return e.EParamInvalidPID
	}

	privilege, err := p.privilegeDAO.Update(req)
	if err != nil {
		return err
	}

	p.copy(resp, privilege)
	return nil
}

func (p *Privilege) copy(dst *privilegeproto.Privilege, src *model.Privilege) {
	dst.ID = src.ID
	dst.Name = src.Name
	dst.PID = src.PID
	dst.Icon = src.Icon
	dst.URIRule = src.URIRule
	dst.IsMenu = src.IsMenu
	dst.SortOrder = src.SortOrder
	dst.CreatedAt = src.CreatedAt
	dst.UpdatedAt = src.UpdatedAt
	dst.DeletedAt = src.DeletedAt
}

func (p *Privilege) Query(_ *struct{}, resp *privilegeproto.QueryResp) error {
	privileges, err := p.privilegeDAO.Privileges()
	if err != nil {
		return err
	}

	for _, privilege := range privileges {
		privilegeResp := &privilegeproto.Privilege{}
		p.copy(privilegeResp, privilege)

		resp.Privileges = append(resp.Privileges, privilegeResp)
	}

	return nil
}

func (p *Privilege) Delete(id uint, _ *struct{}) error {
	if id < 1 {
		return e.EParamInvalidID
	}

	success, err := p.privilegeDAO.Delete(id)
	if err != nil {
		return err
	}

	if !success {
		return e.EDeleteFailed
	}
	return nil
}

func (p *Privilege) RoleIDs(id uint, resp *privilegeproto.PrivilegeRoleResp) error {
	if id < 1 {
		return e.EParamInvalidID
	}

	roleIDs, err := p.rolePrivilegeDAO.FindRoleIDsByPrivilegeID(id)
	if err != nil {
		return err
	}

	resp.RoleIDs = append(resp.RoleIDs, roleIDs...)
	return nil
}
