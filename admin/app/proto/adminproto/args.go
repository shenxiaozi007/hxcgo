package adminproto

type QueryReq struct {
	GroupID uint
	State   int8
	Page    uint
	Limit   uint
}

type AssociateReq struct {
	ID     uint
	RoleID uint
}
