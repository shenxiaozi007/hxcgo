package adminservice

import (
	"admin/app/proto/adminproto"
	"admin/app/service"
	"admin/core/encrypt"
	"admin/core/reg"
	"log"
	"time"

	e "admin/app/err"
)

type Admin struct {
}

func NewAdmin() *Admin {
	return &Admin{}
}

func (a *Admin) FindByID(id uint) (*adminproto.Admin, error) {
	if id < 1 {
		return nil, e.EInvalidID
	}

	resp, err := adminproto.FindByID(id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *Admin) FindByName(name string) (*adminproto.Admin, error) {
	if name == "" {
		return nil, e.EInvalidName
	}

	resp, err := adminproto.FindByName(name)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *Admin) FindByEmail(email string) (*adminproto.Admin, error) {
	if email == "" {
		return nil, e.EEmail
	}

	resp, err := adminproto.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *Admin) FindByMobile(mobile string) (*adminproto.Admin, error) {
	if mobile == "" {
		return nil, e.EMobile
	}

	resp, err := adminproto.FindByMobile(mobile)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *Admin) FindByAccount(account string) (*adminproto.Admin, error) {
	if reg.Email(account) {
		return a.FindByEmail(account)
	} else if reg.Mobile(account) {
		return a.FindByMobile(account)
	}

	return a.FindByName(account)
}

func (a *Admin) FakePassword() string {
	return "##########"
}

func (a *Admin) verifyReq(req *adminproto.Admin) error {
	if req.GroupID < 1 {
		return e.EGroup
	}
	if len(req.Name) < 5 {
		return e.EAdminName
	}
	if len(req.Password) < 8 {
		return e.EPassword
	}

	//检查角色是否合法
	groupService := &Group{}
	group, err := groupService.FindByID(req.GroupID)
	if err != nil || group.ID != req.GroupID || group.State != service.StateActive {
		return e.EGroupNotExists
	}

	admin, err := adminproto.FindByName(req.Name)
	if err == nil && admin.ID != 0 && admin.ID != req.ID {
		return e.EAdminNameExists
	}

	if req.Mobile != "" {
		admin, err = adminproto.FindByMobile(req.Mobile)
		log.Println(err, admin)
		if err == nil && admin.ID != 0 && admin.ID != req.ID {
			return e.EMobileExists
		}
	}
	if req.Email != "" {
		admin, err = adminproto.FindByEmail(req.Email)
		if err == nil && admin.ID != 0 && admin.ID != req.ID {
			return e.EEmailExists
		}
	}

	return nil
}

func (a *Admin) Add(admin *adminproto.Admin) error {
	admin.ID = 0
	err := a.verifyReq(admin)
	if err != nil {
		return err
	}

	admin.Password, err = encrypt.Password(admin.Password)
	if err != nil {
		return e.EPassword
	}

	return admin.Save()
}

func (a *Admin) Update(admin *adminproto.Admin) error {
	if admin.ID < 1 {
		return e.EParam
	}
	err := a.verifyReq(admin)
	if err != nil {
		return err
	}

	var hashPassword string
	if admin.Password != a.FakePassword() {
		hashPassword, err = encrypt.Password(admin.Password)
		if err != nil {
			return e.EPassword
		}
	}
	admin.Password = hashPassword

	return admin.Save()
}

func (a *Admin) UpdateLoginTimeAndIP(admin *adminproto.Admin) error {
	if admin.ID < 1 {
		return e.EParam
	}

	if admin.LoginAt == nil {
		now := time.Now()
		admin.LoginAt = &now
	}

	return admin.UpdateLoginTimeAndIP()
}

func (a *Admin) Query(req *adminproto.QueryReq) (*adminproto.QueryResp, error) {
	if req.Page < 1 {
		req.Page = 1
	}

	if req.Limit < 1 {
		req.Limit = 20
	}

	resp, err := adminproto.Query(req)

	return resp, err
}

func (a *Admin) Delete(id uint) error {
	if id < 1 {
		return e.EInvalidID
	}

	return adminproto.Delete(id)
}

func (a *Admin) RoleIDs(id uint) ([]uint, error) {
	if id < 1 {
		return nil, e.EInvalidID
	}

	ids, err := adminproto.RoleIDs(id)
	if err != nil {
		return nil, err
	}

	roleService := &Role{}
	activeRoleIDs, err := roleService.Actives()
	if err != nil {
		return nil, err
	}

	var ret []uint
	for _, id := range ids {
		_, ok := activeRoleIDs[id]
		if ok {
			ret = append(ret, id)
		}
	}
	return ret, nil
}

func (a *Admin) AssociateRole(id uint, roleID uint) error {
	if id < 1 || roleID < 1 {
		return e.EParam
	}

	req := &adminproto.AssociateReq{
		ID:     id,
		RoleID: roleID,
	}
	return adminproto.AssociateRole(req)
}
