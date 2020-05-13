package admingroupproto

import (
	"github.com/huangxinchun/hxcgo/admin/core/rpc"
	"log"
	"time"
)

type Group struct {
	ID        uint
	Name      string
	State     uint8
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func FindByID(id uint) (*Group, error) {
	group := &Group{}
	err := rpc.Service("admin").Call("Group.FindByID", id, group)
	return group, err
}

func (r *Group) Create() error {
	return rpc.Service("admin").Call("Group.Add", r, r)
}

func (r *Group) Update() error {
	return rpc.Service("admin").Call("Group.Update", r, r)
}

func (r *Group) Save() error {
	if r.ID > 0 {
		return r.Update()
	}

	return r.Create()
}

type QueryResp struct {
	Groups    []*Group
	Page      uint
	Limit     uint
	Count     uint
	TotalPage uint
}

func Query(req *QueryReq) (*QueryResp, error) {
	resp := &QueryResp{}
	err := rpc.Service("admin").Call("Group.Query", req, resp)
	log.Println(err)
	return resp, err
}

func Delete(id uint) error {
	return rpc.Service("admin").Call("Group.Delete", id, nil)
}
