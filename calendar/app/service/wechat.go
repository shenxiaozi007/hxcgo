package service

import (
	e "calendar/app/err"
	"calendar/app/proto/wechatproto"
	"calendar/core/opt"
)

type Wechat struct {
	appName string
}

func NewWechat() *Wechat {
	return &Wechat{
		appName: opt.Config().AppName,
	}
}

func (w *Wechat) Session(code string) (*wechatproto.Session, error) {
	if code == "" {
		return nil, e.EParam
	}

	req := &wechatproto.SessionReq{
		AppName: w.appName,
		Code:    code,
	}

	return wechatproto.GetSession(req)
}

func (w *Wechat) DecryptData(sessionKey string, iv string, encryptedData string) (*wechatproto.DecryptUser, error) {
	if sessionKey == "" || iv == "" || encryptedData == "" {
		return nil, e.EParam
	}

	req := &wechatproto.DecryptReq{
		AppName:       w.appName,
		SessionKey:    sessionKey,
		IV:            iv,
		EncryptedData: encryptedData,
	}

	return wechatproto.DecryptData(req)
}
