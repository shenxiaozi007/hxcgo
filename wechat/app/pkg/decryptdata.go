package pkg

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	e "wechat/app/err"
	"wechat/app/proto"
)

func Decrypt(appID, sessionKey, iv, encryptedData string) (*proto.DecryptUser, error) {
	dkey, err := base64.StdEncoding.DecodeString(sessionKey)
	if err != nil {
		return nil, err
	}

	div, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return nil, err
	}

	ddata, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return nil, err
	}

	//解密
	block, err := aes.NewCipher(dkey)
	if err != nil {
		return nil, err
	}
	mode := cipher.NewCBCDecrypter(block, div)
	mode.CryptBlocks(ddata, ddata)

	l := len(ddata)
	var i = 1
	for i <= l {
		if ddata[l-i] == 125 {
			ddata = ddata[:l-i+1]
			break
		}
		i++
	}

	user := &proto.DecryptUser{}
	err = json.Unmarshal(ddata, user)
	if err != nil {
		return nil, err
	}

	if user.Watermark.AppID != appID {
		return nil, e.EInvalidAppID
	}

	return user, nil
}
