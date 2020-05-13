package ctrl

import (
	"github.com/huangxinchun/hxcgo/admin/app/dao"
	e "github.com/huangxinchun/hxcgo/admin/app/err"
	"github.com/huangxinchun/hxcgo/admin/app/model"
	"github.com/huangxinchun/hxcgo/admin/app/proto/adminproto"
	"math"
)

type Admin struct {
	adminDAO     *dao.Admin
	adminRoleDAO *dao.AdminRole
}

func NewAdmin() *Admin {
	return &Admin{
		adminDAO:     dao.NewAdmin(),
		adminRoleDAO: dao.NewAdminRole(),
	}
}

func (a *Admin) FindByName(name string, resp *adminproto.Admin) error {
	if name == "" {
		return e.EParamInvalidName
	}

	admin, err := a.adminDAO.FindByName(name)
	if err != nil {
		return err
	}
	a.copy(resp, admin)

	return nil
}

func (a *Admin) FindByID(id uint, resp *adminproto.Admin) error {
	if id < 1 {
		return e.EParamInvalidID
	}

	admin, err := a.adminDAO.FindByID(id)
	if err != nil {
		return err
	}

	a.copy(resp, admin)
	return nil
}

func (a *Admin) FindByEmail(email string, resp *adminproto.Admin) error {
	if email == "" {
		return e.EParamInvalidEmail
	}

	admin, err := a.adminDAO.FindByEmail(email)
	if err != nil {
		return err
	}

	a.copy(resp, admin)
	return nil
}

func (a *Admin) FindByMobile(mobile string, resp *adminproto.Admin) error {
	if mobile == "" {
		return e.EParamInvalidMobile
	}

	admin, err := a.adminDAO.FindByMobile(mobile)
	if err != nil {
		return err
	}

	a.copy(resp, admin)
	return nil
}

func (a *Admin) copy(dst *adminproto.Admin, src *model.Admin) {
	dst.ID = src.ID
	dst.Name = src.Name
	dst.Avatar = src.Avatar
	dst.Password = src.Password
	dst.State = src.State
	dst.Mobile = src.Mobile
	dst.Email = src.Email
	dst.GroupID = src.GroupID
	dst.CreatedAt = src.CreatedAt
	dst.UpdatedAt = src.UpdatedAt
	dst.DeletedAt = src.DeletedAt
	dst.LoginAt = src.LoginAt
	dst.LoginIP = src.LoginIP
}

func (a *Admin) Add(req *adminproto.Admin, resp *adminproto.Admin) error {
	if req.GroupID < 1 {
		return e.EParamInvalidGroupID
	}

	if len(req.Name) < 5 {
		return e.EParamInvalidName
	}

	admin, err := a.adminDAO.Add(req)
	if err != nil {
		return err
	}

	a.copy(resp, admin)
	return nil
}

func (a *Admin) Update(req *adminproto.Admin, resp *adminproto.Admin) error {
	if req.ID < 1 {
		return e.EParamInvalidID
	}
	if req.GroupID < 1 {
		return e.EParamInvalidGroupID
	}

	if len(req.Name) < 5 {
		return e.EParamInvalidName
	}

	admin, err := a.adminDAO.Update(req)
	if err != nil {
		return err
	}

	a.copy(resp, admin)
	return nil
}

func (a *Admin) UpdateLoginTimeAndIP(req *adminproto.Admin, resp *adminproto.Admin) error {
	if req.ID < 1 {
		return e.EParamInvalidID
	}

	admin, err := a.adminDAO.UpdateLoginTimeAndIP(req)
	if err != nil {
		return err
	}

	a.copy(resp, admin)
	return nil
}

func (a *Admin) Query(req *adminproto.QueryReq, resp *adminproto.QueryResp) error {
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
	if req.GroupID > 0 {
		conditions["group_id"] = req.GroupID
	}

	count, err := a.adminDAO.Count(conditions)
	if err != nil {
		return err
	}

	resp.Page = req.Page
	resp.Limit = req.Limit
	resp.Count = count
	if count == 0 {
		return nil
	}

	admins, err := a.adminDAO.Admins(conditions, req.Page, req.Limit)
	if err != nil {
		return err
	}

	resp.TotalPage = uint(math.Ceil(float64(resp.Count) / float64(resp.Limit)))

	for _, admin := range admins {
		adminResp := &adminproto.Admin{}
		a.copy(adminResp, admin)

		resp.Admins = append(resp.Admins, adminResp)
	}

	return nil
}

func (a *Admin) Delete(id uint, _ *struct{}) error {
	if id < 1 {
		return e.EParamInvalidID
	}

	success, err := a.adminDAO.Delete(id)
	if err != nil {
		return err
	}

	if !success {
		return e.EDeleteFailed
	}
	return nil
}

func (a *Admin) RoleIDs(id uint, resp *adminproto.AdminRoleResp) error {
	if id < 1 {
		return e.EParamInvalidID
	}

	roleIDs, err := a.adminRoleDAO.FindRoleIDsByAdminID(id)
	if err != nil {
		return err
	}

	resp.RoleIDs = append(resp.RoleIDs, roleIDs...)
	return nil
}

func (a *Admin) AssociateRole(req *adminproto.AssociateReq, _ *struct{}) error {
	if req.ID < 1 || req.RoleID < 1 {
		return e.EParam
	}

	success, err := a.adminRoleDAO.Associate(req.ID, req.RoleID)
	if err != nil {
		return err
	}
	if !success {
		return e.EAssociateRolePrivilegeFailed
	}
	return nil
}
