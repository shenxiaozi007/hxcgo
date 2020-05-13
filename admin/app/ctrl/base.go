package ctrl

import (
	"github.com/huangxinchun/hxcgo/admin/app/pkg"
	"github.com/huangxinchun/hxcgo/admin/conf"
	"github.com/huangxinchun/hxcgo/admin/conf/language"
	"net/http"

	"github.com/gin-gonic/gin"

	e "github.com/huangxinchun/hxcgo/admin/app/err"
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

func LanguageMessage(key string) string {
	switch Language {
	case "zh_ch":
		return language.ZH_CN[key]
	case "en":
		return language.EN[key]

	}

	return language.ZH_CN[key]
}

func Render(c *gin.Context, templateName string, data gin.H) {
	if data == nil {
		data = gin.H{}
	}
	sess := pkg.NewSession(c)
	//data["currentAdminName"] = sess.Name
	//data["currentAdminID"] = sess.ID
	//data["currentAdminAvatar"] = sess.Avatar
	data["sidebarMenus"] = sess.Menus()

	var activeMenuID uint
	privilege, exists := pkg.PrivilegeURITree().Match(c.Request.URL.String())
	if exists {
		activeMenuID = privilege.PID
	}
	data["activeMenuID"] = activeMenuID
	data["session"] = sess

	c.HTML(http.StatusOK, templateName, data)
}
