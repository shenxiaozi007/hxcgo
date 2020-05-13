package ctrl

import (
	"wechat/app/api"
	"wechat/app/pkg"
	"wechat/app/proto"
	e "wechat/app/err"
)

type Session struct {

}

func NewSession() *Session {
	return &Session{}
}

func (s *Session) Login(req *proto.SessionReq, resp *proto.Session) error {
	if req.AppName == "" || req.Code == "" {
		return e.EParam
	}

	cfg,err := getWechatConfig(req.AppName)
	if err != nil {
		return err
	}

	sess, err := api.Session(cfg.AppID, cfg.Secret, req.Code)
	if err != nil {
		return err
	}

	resp.OpenID = sess.OpenID
	resp.Key = sess.Key
	resp.UnionID = sess.UnionID
	resp.ErrCode = sess.ErrCode
	resp.ErrMsg = sess.ErrMsg

	return nil
}

func (s *Session) Decrypt(req *proto.DecryptReq, resp *proto.DecryptUser) error {
	if req.AppName == "" || req.SessionKey == "" || req.IV == "" || req.EncryptedData == "" {
		return e.EParam
	}

	cfg,err := getWechatConfig(req.AppName)
	if err != nil {
		return err
	}

	user, err := pkg.Decrypt(cfg.AppID, req.SessionKey, req.IV, req.EncryptedData)
	if err != nil {
		return err
	}

	resp.OpenID = user.OpenID
	resp.UnionID = user.UnionID
	resp.Gender = user.Gender
	resp.Language = user.Language
	resp.NickName = user.NickName
	resp.Country = user.Country
	resp.Province = user.Province
	resp.City = user.City
	resp.AvatarURL = user.AvatarURL
	resp.Watermark = user.Watermark
	return nil
}
