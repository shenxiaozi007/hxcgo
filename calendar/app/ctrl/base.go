package ctrl

import (
	"net/http"
	"calendar/conf"
	"calendar/conf/language"

	"github.com/gin-gonic/gin"

	e "calendar/app/err"
)

type ImageSize struct {
	Height int
	Width  int
}

var Language = "zh_ch"

func Json(c *gin.Context, err error, data interface{}) {
	var key string
	if err != nil {
		key = err.Error()
	} else {
		key = "success"
	}
	code, ok := conf.ResponseCodes[key]
	if !ok {
		code = 4999
	}

	msg := LanguageMessage(key)

	ret := map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	}
	c.Header("token", "123")
	c.JSON(http.StatusOK, ret)
}

func Error(err error) error {
	if err == nil {
		return nil
	}
	_, ok := conf.ResponseCodes[err.Error()]
	if !ok {
		return e.ERequest
	}
	return err
}

func CompareError(err1 error, err2 error) bool {
	return err1.Error() == err1.Error()
}

func LanguageMessage(key string) string {
	switch Language {
	case "zh_ch":
		return language.ZH_CN[key]
	case "en":
		return language.EN[key]

	}

	return language.ZH_CN[key]
}
