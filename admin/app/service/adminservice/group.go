package adminservice

import (
	e "admin/app/err"
	"admin/app/proto/admingroupproto"
)

type Group struct {
}

func NewGroup() *Group {
	return &Group{}
}

func (g *Group) FindByID(id uint) (*admingroupproto.Group, error) {
	if id < 1 {
		return nil, e.EInvalidID
	}

	resp, err := admingroupproto.FindByID(id)

	return resp, err
}

func (g *Group) Add(group *admingroupproto.Group) error {
	group.ID = 0
	if group.Name == "" {
		return e.EInvalidName
	}

	return group.Save()
}

func (g *Group) Update(group *admingroupproto.Group) error {
	if group.ID < 1 {
		return e.EInvalidID
	}
	if group.Name == "" {
		return e.EInvalidName
	}

	return group.Save()
}

func (g *Group) Query(req *admingroupproto.QueryReq) (*admingroupproto.QueryResp, error) {
	if req.Page < 1 {
		req.Page = 1
	}

	if req.Limit < 1 {
		req.Limit = 20
	}

	resp, err := admingroupproto.Query(req)

	return resp, err
}

func (g *Group) Delete(id uint) error {
	if id < 1 {
		return e.EInvalidID
	}
	return admingroupproto.Delete(id)
}

func (g *Group) Actives() (map[uint]*admingroupproto.Group, error) {
	req := &admingroupproto.QueryReq{
		State: 0,
		Page:  1,
		Limit: 200,
	}

	resp, err := g.Query(req)
	if err != nil {
		return nil, err
	}

	groups := map[uint]*admingroupproto.Group{}
	for _, group := range resp.Groups {
		groups[group.ID] = group
	}

	return groups, err
}
