package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
	"wechat/app/cachekey"
	e "wechat/app/err"
	"wechat/core/cache"
)

var accessTokenLock = sync.Mutex{}

func AccessToken(appID string,secret string) (string,error) {
	var token string

	cKey := cachekey.AccessToken(appID)
	if cache.Get(cKey,&token) == nil {
		return token,nil
	}

	accessTokenLock.Lock()
	defer accessTokenLock.Unlock()

	if cache.Get(cKey,&token) == nil {
		return token,nil
	}

	urlStr := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", appID, secret)
	resp, err := http.Get(urlStr)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	type Result struct {
		AccessToken string `json:"access_token"`
		Expires int `json:"expires_in"`
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
	}

	result := &Result{}
	err = json.Unmarshal(buf, result)
	if err != nil {
		return "", err
	}

	switch result.ErrCode {
	case 0:
		//成功

		cache.Set(cKey,result.AccessToken,time.Duration(result.Expires-120)*time.Second)
		return result.AccessToken, nil
	case -1:
		return "", e.ESystemBusy
	case 40001:
		return "", e.EInvalidAppSecret
	case 40002:
		return "", e.EInvalidGrantType
	case 40013:
		return "", e.EInvalidAppID
	}

	return "", e.EUnknown
}
