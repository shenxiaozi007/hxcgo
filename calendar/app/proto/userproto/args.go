package userproto

type QueryReq struct {
	NickName     string
	RegBeginTime string
	RegEndTime   string
	Page         uint
	Limit        uint
}
