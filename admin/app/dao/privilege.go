package dao

import (
	e "github.com/huangxinchun/hxcgo/admin/app/err"
	"github.com/huangxinchun/hxcgo/admin/app/model"
	"github.com/huangxinchun/hxcgo/admin/app/proto/privilegeproto"
	"github.com/huangxinchun/hxcgo/admin/core/db"
	"fmt"
)

type Privilege struct {
}

func NewPrivilege() *Privilege {
	return &Privilege{}
}

func (p *Privilege) FindByID(id uint) (*model.Privilege, error) {
	if id < 1 {
		return nil, e.EParamInvalidID
	}

	privilege := &model.Privilege{}
	err := privilege.FindByID(id)
	if err != nil {
		return nil, e.EDatabase
	}
	return privilege, nil
}

func (p *Privilege) verifyReq(req *privilegeproto.Privilege) error {
	if req.Name == "" {
		return e.EParamInvalidName
	}
	req.Root = ""
	if req.PID > 0 {
		parent, err := p.FindByID(req.PID)
		if err != nil {
			return e.EParamInvalidPID
		}
		if parent.Root != "" {
			req.Root = fmt.Sprintf("%s%d", parent.Root, req.PID)
		} else {
			req.Root = fmt.Sprintf("%d/", req.PID)
		}
	}

	return nil
}

func (p *Privilege) copy(dst *model.Privilege, src *privilegeproto.Privilege) {
	dst.ID = src.ID
	dst.PID = src.PID
	dst.Name = src.Name
	dst.Root = src.Root
	dst.Icon = src.Icon
	dst.URIRule = src.URIRule
	dst.IsMenu = src.IsMenu
	dst.SortOrder = src.SortOrder
	dst.UpdatedAt = src.UpdatedAt
}

func (p *Privilege) Update(req *privilegeproto.Privilege) (*model.Privilege, error) {
	if req.ID < 1 || req.ID == req.PID {
		return nil, e.EParam
	}
	err := p.verifyReq(req)
	if err != nil {
		return nil, err
	}

	privilege, err := p.FindByID(req.ID)
	if err != nil {
		return nil, e.EParamInvalidID
	}

	p.copy(privilege, req)
	err = privilege.Save()

	return privilege, err
}

func (p *Privilege) Add(req *privilegeproto.Privilege) (*model.Privilege, error) {
	if req.ID != 0 {
		return nil, e.EParam
	}
	err := p.verifyReq(req)
	if err != nil {
		return nil, err
	}

	privilege := &model.Privilege{}
	p.copy(privilege, req)
	err = privilege.Save()

	return privilege, err
}

func (p *Privilege) Privileges() ([]*model.Privilege, error) {
	var privileges []*model.Privilege
	err := db.Client().Order("sort_order asc,id asc").Limit(2000).Find(&privileges).Error

	return privileges, err
}

func (p *Privilege) Delete(id uint) (bool, error) {
	if id < 1 {
		return false, e.EParamInvalidID
	}

	privilege, err := p.FindByID(id)
	if err != nil {
		return false, err
	}

	var root string
	if privilege.Root != "" {
		root = fmt.Sprintf("%s/%d", privilege.Root, privilege.ID)
	} else {
		root = fmt.Sprintf("%d/", privilege.ID)
	}

	success := privilege.Delete(id)
	var privileges []*model.Privilege
	db.Client().Where("root LIKE ?", fmt.Sprintf("%s%%", root)).Find(privileges)
	rolePrivilegeDAO := NewRolePrivilege()
	rolePrivilegeDAO.DeleteByPrivilegeID(id)
	for _,p := range privileges {
		p.Delete(p.ID)
		rolePrivilegeDAO.DeleteByPrivilegeID(p.ID)
	}
	return success, nil
}
