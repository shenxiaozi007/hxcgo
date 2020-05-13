package groupproto

import "time"

type Group struct {
	ID        uint
	Name      string
	State     uint8
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type QueryResp struct {
	Groups    []*Group
	Page      uint
	Limit     uint
	Count     uint
	TotalPage uint
}
