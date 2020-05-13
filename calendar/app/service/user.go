package service

import (
	"calendar/app/cachekey"
	e "calendar/app/err"
	"calendar/app/proto/userproto"
	"calendar/core/cache"
	"time"
)

type User struct {
}

func NewUser() *User {
	return &User{}
}

func (u *User) FindByOpenID(openID string) (*userproto.User, error) {
	if openID == "" {
		return nil, e.EParam
	}

	//尝试从缓存中获取数据
	user := &userproto.User{}
	cKey := cachekey.UserByOpenID(openID)
	if cache.Get(cKey, user) == nil {
		return user, nil
	}

	//从user服务获取数据
	user, err := userproto.FindByOpenID(openID)
	if err != nil {
		return nil, err
	}
	cache.Set(cKey, user, time.Minute)

	return user, nil
}

func (u *User) Add(user *userproto.User) error {
	if user.ID != 0 || user.OpenID == "" {
		return e.EParam
	}

	user.State = 0
	err := user.Save()
	if err != nil {
		cache.Set(cachekey.UserByOpenID(user.OpenID), user, time.Minute)
	}
	return err
}

func (u *User) Update(user *userproto.User) error {
	if user.OpenID == "" {
		return e.EParam
	}

	err := user.Save()
	if err != nil {
		cache.Set(cachekey.UserByOpenID(user.OpenID), user, time.Minute)
	}
	return err
}
