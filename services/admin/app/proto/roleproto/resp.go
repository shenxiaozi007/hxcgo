package roleproto

import "time"

type Role struct {
	ID        uint
	Name      string
	State     uint8
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type QueryResp struct {
	Roles     []*Role
	Page      uint
	Limit     uint
	Count     uint
	TotalPage uint
}

type RolePrivilegeResp struct {
	PrivilegeIDs []uint
}
