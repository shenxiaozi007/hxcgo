package ctrl

import (
	"github.com/huangxinchun/hxcgo/admin/app/dao"
	e "github.com/huangxinchun/hxcgo/admin/app/err"
	"github.com/huangxinchun/hxcgo/admin/app/model"
	"github.com/huangxinchun/hxcgo/admin/app/proto/groupproto"
	"math"
)

type Group struct {
	groupDAO *dao.Group
}

func NewGroup() *Group {
	return &Group{
		groupDAO: dao.NewGroup(),
	}
}

func (g *Group) FindByID(id uint, resp *groupproto.Group) error {
	if id < 1 {
		return e.EParamInvalidID
	}

	group, err := g.groupDAO.FindByID(id)
	if err != nil {
		return err
	}

	g.copy(resp, group)
	return nil
}

func (g *Group) Add(req *groupproto.Group, resp *groupproto.Group) error {
	if req.Name == "" {
		return e.EParamInvalidName
	}

	group, err := g.groupDAO.Add(req)
	if err != nil {
		return err
	}

	g.copy(resp, group)
	return nil
}

func (g *Group) Update(req *groupproto.Group, resp *groupproto.Group) error {
	if req.Name == "" {
		return e.EParamInvalidName
	}

	if req.ID < 1 {
		return e.EParamInvalidID
	}

	group, err := g.groupDAO.Update(req)
	if err != nil {
		return err
	}

	g.copy(resp, group)
	return nil
}

func (g *Group) copy(dst *groupproto.Group, src *model.Group) {
	dst.ID = src.ID
	dst.Name = src.Name
	dst.State = src.State
	dst.CreatedAt = src.CreatedAt
	dst.UpdatedAt = src.UpdatedAt
	dst.DeletedAt = src.DeletedAt
}

func (g *Group) Query(req *groupproto.QueryReq, resp *groupproto.QueryResp) error {
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

	count, err := g.groupDAO.Count(conditions)
	if err != nil {
		return err
	}

	resp.Page = req.Page
	resp.Limit = req.Limit
	resp.Count = count
	if count == 0 {
		return nil
	}

	groups, err := g.groupDAO.Groups(conditions, req.Page, req.Limit)
	if err != nil {
		return err
	}

	resp.TotalPage = uint(math.Ceil(float64(resp.Count) / float64(resp.Limit)))

	for _, group := range groups {
		groupResp := &groupproto.Group{}
		g.copy(groupResp, group)

		resp.Groups = append(resp.Groups, groupResp)
	}

	return nil
}

func (g *Group) Delete(id uint, _ *struct{}) error {
	if id < 1 {
		return e.EParamInvalidID
	}

	success, err := g.groupDAO.Delete(id)
	if err != nil {
		return err
	}

	if !success {
		return e.EDeleteFailed
	}
	return nil
}
