package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	e "wechat/app/err"
	"wechat/app/proto"
)

func Session(appID string, secret string, code string) (*proto.Session, error) {
	urlStr := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", appID, secret, code)
	resp, err := http.Get(urlStr)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	session := &proto.Session{}
	err = json.Unmarshal(buf, session)
	if err != nil {
		return nil, err
	}

	switch session.ErrCode {
	case 0:
		//成功
		return session, nil
	case -1:
		return nil, e.ESystemBusy
	case 40029:
		return nil, e.EInvalidCode
	case 45011:
		return nil, e.EAPILimit
	}

	return nil, e.EUnknown
}
