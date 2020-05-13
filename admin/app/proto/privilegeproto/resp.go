package privilegeproto

import "time"

type Privilege struct {
	ID         uint
	PID        uint
	Name       string
	Root       string
	Icon       string
	URIRule    string
	IsMenu     uint16
	SortOrder  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}

type QueryResp struct {
	Privileges []*Privilege
}

type PrivilegeRoleResp struct {
	RoleIDs []uint
}
