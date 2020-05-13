package wechatproto

type SessionReq struct {
	AppName string
	Code    string
}

type DecryptReq struct {
	AppName       string
	SessionKey    string
	IV            string
	EncryptedData string
}
