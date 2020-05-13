package adminroleproto

type QueryReq struct {
	State int8
	Page  uint
	Limit uint
}

type AssociateReq struct {
	ID          uint
	PrivilegeID uint
}
