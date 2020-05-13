package userproto

import "time"

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

type QueryResp struct {
	Users     []*User
	Page      uint
	Limit     uint
	Count     uint
	TotalPage uint
}
