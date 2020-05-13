package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	e "wechat/app/err"
	"wechat/app/proto"
)

func QRCodeUnlimited(accessToken string,args *proto.QRCodeUnlimitedReq) ([]byte,string,error) {
	if accessToken == "" {
		return nil,"",e.EInvalidAccessToken
	}
	urlStr := fmt.Sprintf("https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=%s",accessToken)

	js,err := json.Marshal(args)
	if err != nil {
		return nil,"",err
	}

	reader := bytes.NewBuffer(nil)
	reader.Write(js)

	resp,err := http.Post(urlStr,"application/json",reader)
	if err != nil {
		return nil,"",err
	}
	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,"", err
	}

	mimeType := resp.Header.Get("Content-Type")
	if IsImage(mimeType) {
		return buf,mimeType,nil
	}

	type Result struct {
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
	}

	result := &Result{}
	err = json.Unmarshal(buf, result)
	if err != nil {
		return nil,"", err
	}

	switch result.ErrCode {
	case 45009:
		return nil,"", e.EAPILimit
	case 41030:
		return nil,"", e.EPageNotFoundOrMiniProgramNotPublish
	}

	return nil,"", e.EUnknown
}

func IsImage(mtype string) bool {
	mimeTypes := []string{
		"image/gif",
		"image/jpeg",
		"image/jpg",
		"image/png",
	}

	isImage := false
	for _, v := range mimeTypes {
		if v == mtype {
			isImage = true
			break
		}
	}
	return isImage
}