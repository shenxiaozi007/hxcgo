package adminproto

import (
	"github.com/huangxinchun/hxcgo/admin/core/rpc"
	"time"
)

type Admin struct {
	ID        uint
	Name      string
	Email     string
	Mobile    string
	Password  string
	Avatar    string
	State     uint8
	GroupID   uint
	LoginAt   *time.Time
	LoginIP   int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func FindByID(id uint) (*Admin, error) {
	admin := &Admin{}
	err := rpc.Service("admin").Call("Admin.FindByID", id, admin)

	return admin, err
}

func FindByName(name string) (*Admin, error) {
	admin := &Admin{}
	err := rpc.Service("admin").Call("Admin.FindByName", name, admin)

	return admin, err
}

func FindByEmail(email string) (*Admin, error) {
	admin := &Admin{}
	err := rpc.Service("admin").Call("Admin.FindByEmail", email, admin)

	return admin, err
}

func FindByMobile(mobile string) (*Admin, error) {
	admin := &Admin{}
	err := rpc.Service("admin").Call("Admin.FindByMobile", mobile, admin)

	return admin, err
}

func (a *Admin) Create() error {
	return rpc.Service("admin").Call("Admin.Add", a, a)
}

func (a *Admin) Update() error {
	return rpc.Service("admin").Call("Admin.Update", a, a)
}
func (a *Admin) Save() error {
	if a.ID > 0 {
		return a.Update()
	}
	return a.Create()
}

func (a *Admin) UpdateLoginTimeAndIP() error {
	return rpc.Service("admin").Call("Admin.UpdateLoginTimeAndIP", a, a)
}

func Delete(id uint) error {
	return rpc.Service("admin").Call("Admin.Delete", id, nil)
}

type QueryResp struct {
	Admins    []*Admin
	Page      uint
	Limit     uint
	Count     uint
	TotalPage uint
}

func Query(req *QueryReq) (*QueryResp, error) {
	resp := &QueryResp{}
	err := rpc.Service("admin").Call("Admin.Query", req, resp)

	return resp, err
}

type AdminRoleResp struct {
	RoleIDs []uint
}

func RoleIDs(id uint) ([]uint, error) {
	resp := &AdminRoleResp{}
	err := rpc.Service("admin").Call("Admin.RoleIDs", id, resp)
	return resp.RoleIDs, err
}

func AssociateRole(req *AssociateReq) error {
	return rpc.Service("admin").Call("Admin.AssociateRole", req, nil)
}
