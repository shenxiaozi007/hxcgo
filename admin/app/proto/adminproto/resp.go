package adminproto

import "time"

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

type QueryResp struct {
	Admins    []*Admin
	Page      uint
	Limit     uint
	Count     uint
	TotalPage uint
}

type AdminRoleResp struct {
	RoleIDs []uint
}
