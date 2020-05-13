package dao

import (
	e "github.com/huangxinchun/hxcgo/admin/app/err"
	"github.com/huangxinchun/hxcgo/admin/app/model"
	"github.com/huangxinchun/hxcgo/admin/app/proto/groupproto"
	"github.com/huangxinchun/hxcgo/admin/core/db"
)

type Group struct {
}

func NewGroup() *Group {
	return &Group{}
}

func (g *Group) FindByID(id uint) (*model.Group, error) {
	if id < 1 {
		return nil, e.EParamInvalidID
	}

	group := &model.Group{}
	err := group.FindByID(id)
	if err != nil {
		return nil, e.EDatabase
	}
	return group, nil
}

func (g *Group) Update(req *groupproto.Group) (*model.Group, error) {
	if req.ID < 1 || req.Name == "" {
		return nil, e.EParam
	}
	group := &model.Group{
		ID:    req.ID,
		Name:  req.Name,
		State: req.State,
	}
	err := group.Save()

	return group, err
}

func (g *Group) Add(req *groupproto.Group) (*model.Group, error) {
	if req.Name == "" {
		return nil, e.EParamInvalidName
	}
	group := &model.Group{
		Name:  req.Name,
		State: req.State,
	}
	err := group.Save()

	return group, err
}

func (g *Group) Groups(conditions map[string]interface{}, page uint, limit uint) ([]*model.Group, error) {
	if page < 1 {
		page = 1
	}

	offset := (page - 1) * limit

	var groups []*model.Group
	err := db.Client().Where(conditions).Offset(offset).Limit(limit).Find(&groups).Error

	return groups, err
}

func (g *Group) Count(conditions map[string]interface{}) (uint, error) {
	var countRecords uint
	err := db.Client().Model(&model.Group{}).Where(conditions).Count(&countRecords).Error
	return countRecords, err
}

func (g *Group) Delete(id uint) (bool, error) {
	if id < 1 {
		return false, e.EParamInvalidID
	}

	group := &model.Group{}
	return group.Delete(id), nil
}
