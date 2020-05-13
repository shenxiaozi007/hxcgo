package ctrl

import (
	"wechat/app/api"
	e "wechat/app/err"
	"wechat/app/proto"
)

type QRCode struct {
	
}
func NewQRCode() *QRCode {
	return &QRCode{}
}

func(qr *QRCode) Unlimited(req *proto.QRCodeUnlimitedReq,resp *proto.QRCode) error {
	if req.AppName == "" {
		return e.EParam
	}

	cfg,err := getWechatConfig(req.AppName)
	if err != nil {
		return err
	}

	accessToken,err := api.AccessToken(cfg.AppID,cfg.Secret)
	if err != nil {
		return err
	}

	resp.Buffer,resp.MIMEType,err = api.QRCodeUnlimited(accessToken,req)
	return err
}
