package ctrl

import (
	e "wechat/app/err"
	"wechat/app/proto"
)

type Wechat struct {
}

func NewWechat() *Wechat {
	return &Wechat{}
}

func (w *Wechat) Config(appName string, resp *proto.AppConfig) error {
	if appName == "" {
		return e.EParam
	}

	wxConfig, err := getWechatConfig(appName)
	if err != nil {
		return err
	}

	resp.AppID = wxConfig.AppID
	resp.Secret = wxConfig.Secret
	return nil
}

