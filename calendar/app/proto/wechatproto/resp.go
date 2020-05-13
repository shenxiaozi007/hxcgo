package wechatproto

import "calendar/core/rpc"

type Session struct {
	OpenID  string
	Key     string
	UnionID string
	ErrCode int
	ErrMsg  string
}

func GetSession(req *SessionReq) (*Session, error) {
	resp := &Session{}
	err := rpc.Service("wechat").Call("Session.Login", req, resp)
	return resp, err
}

type DecryptUser struct {
	OpenID    string
	NickName  string
	Gender    uint8
	Language  string
	City      string
	Province  string
	Country   string
	AvatarURL string
	UnionID   string
	Watermark struct {
		Timestamp int64
		AppID     string
	}
}

func DecryptData(req *DecryptReq) (*DecryptUser, error) {
	resp := &DecryptUser{}
	err := rpc.Service("wechat").Call("Session.Decrypt", req, resp)
	return resp, err
}
