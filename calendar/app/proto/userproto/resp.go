package userproto

import (
	"calendar/core/rpc"
	"time"
)

type User struct {
	ID           uint
	OpenID       string
	UnionID      string
	NickName     string
	Avatar       string
	State        uint8
	Gender       uint8
	Country      string
	Province     string
	City         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}

func FindByOpenID(openID string) (*User, error) {
	user := &User{}
	err := rpc.Service("calendar").Call("User.FindByOpenID", openID, user)
	return user, err
}

func (u *User) Create() error {
	return rpc.Service("calendar").Call("User.Add", u, u)
}

func (u *User) Update() error {
	return rpc.Service("calendar").Call("User.Update", u, u)
}
func (u *User) Save() error {
	if u.ID > 0 {
		return u.Update()
	}
	return u.Create()
}
