package proto

type AppConfig struct {
	AppID  string
	Secret string
}

type Session struct {
	OpenID  string `json:"openid"`
	Key     string `json:"session_key"`
	UnionID string `json:"unionid"`
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type DecryptUser struct {
	OpenID    string `json:"openId"`
	NickName  string `json:"nickName"`
	Gender    uint8  `json:"gender"`
	Language  string `json:"language"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	AvatarURL string `json:"avatarUrl"`
	UnionID   string `json:"unionId"`
	Watermark struct {
		Timestamp int64  `json:"timestamp"`
		AppID     string `json:"appid"`
	} `json:"watermark"`
}

type QRCode struct {
	Buffer []byte
	MIMEType string
}
