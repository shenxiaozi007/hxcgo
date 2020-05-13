package ctrl

import (
	"log"
	"math"
	"calendar/app/dao"
	e "calendar/app/err"
	"calendar/app/model"
	"calendar/app/proto/userproto"
)

type User struct {
	userDAO *dao.User
}

func NewUser() *User {
	return &User{
		userDAO: dao.NewUser(),
	}
}

func (a *User) FindByID(id uint, resp *userproto.User) error {
	if id < 1 {
		return e.EParamInvalidID
	}

	user, err := a.userDAO.FindByID(id)
	if err != nil {
		return err
	}

	a.copy(resp, user)
	return nil
}

func (a *User) FindByOpenID(openID string, resp *userproto.User) error {
	if openID == "" {
		return e.EParamInvalidOpenID
	}

	user, err := a.userDAO.FindByOpenID(openID)
	if err != nil {
		return err
	}

	a.copy(resp, user)
	return nil
}

func (a *User) copy(dst *userproto.User, src *model.User) {
	dst.ID = src.ID
	dst.OpenID = src.OpenID
	dst.UnionID = src.UnionID
	dst.NickName = src.NickName
	dst.Avatar = src.Avatar
	dst.Gender = src.Gender
	dst.State = src.State
	dst.Country = src.Country
	dst.Province = src.Province
	dst.City = src.City
	dst.CreatedAt = src.CreatedAt
	dst.UpdatedAt = src.UpdatedAt
	dst.DeletedAt = src.DeletedAt
}

func (a *User) Add(req *userproto.User, resp *userproto.User) error {
	user, err := a.userDAO.Add(req)
	if err != nil {
		return err
	}

	a.copy(resp, user)
	return nil
}

func (a *User) Update(req *userproto.User, resp *userproto.User) error {
	if req.ID < 1 {
		return e.EParamInvalidID
	}

	user, err := a.userDAO.Update(req)
	if err != nil {
		return err
	}

	a.copy(resp, user)
	return nil
}

func (a *User) Query(req *userproto.QueryReq, resp *userproto.QueryResp) error {
	if req.Page < 1 {
		return e.EParamInvalidPage
	}
	if req.Limit < 1 {
		return e.EParamInvalidLimit
	}

	conditions := map[string]interface{}{}
	if req.NickName != "" {
		conditions["nickName"] = req.NickName
	}
	if req.RegBeginTime != "" {
		conditions["regBeginTime"] = req.RegBeginTime
	}
	if req.RegEndTime != "" {
		conditions["regEndTime"] = req.RegEndTime
	}

	log.Printf("%#v", conditions)
	log.Printf("%#v", req)

	count, err := a.userDAO.Count(conditions)
	if err != nil {
		return err
	}

	resp.Page = req.Page
	resp.Limit = req.Limit
	resp.Count = count
	if count == 0 {
		return nil
	}

	users, err := a.userDAO.Users(conditions, req.Page, req.Limit)
	if err != nil {
		return err
	}

	resp.TotalPage = uint(math.Ceil(float64(resp.Count) / float64(resp.Limit)))

	for _, user := range users {
		userResp := &userproto.User{}
		a.copy(userResp, user)

		resp.Users = append(resp.Users, userResp)
	}

	return nil
}

func (a *User) Delete(id uint, _ *struct{}) error {
	if id < 1 {
		return e.EParamInvalidID
	}

	success, err := a.userDAO.Delete(id)
	if err != nil {
		return err
	}

	if !success {
		return e.EDeleteFailed
	}
	return nil
}
