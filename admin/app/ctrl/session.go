package ctrl

import (
	"github.com/huangxinchun/hxcgo/admin/app/pkg"
	"github.com/huangxinchun/hxcgo/admin/app/service"
	"github.com/huangxinchun/hxcgo/admin/app/service/adminservice"
	"github.com/huangxinchun/hxcgo/admin/core/encrypt"
	"github.com/huangxinchun/hxcgo/admin/core/ip"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	e "github.com/huangxinchun/hxcgo/admin/app/err"

	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Session struct {
}

func NewSession() *Session {
	return &Session{}
}

func (s *Session) getCaptchaID() string {
	return captcha.NewLen(4)
}

func (s *Session) Login(c *gin.Context) {
	c.HTML(http.StatusOK, "session/login.html", gin.H{
		"title":     "Posts",
		"captchaID": s.getCaptchaID(),
	})
}

func (s *Session) Authentication(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	captchaID := c.PostForm("captcha_id")
	captchaCode := c.PostForm("captcha_code")

	resp := map[string]interface{}{}
	resp["captcha_id"] = s.getCaptchaID()
	if !captcha.VerifyString(captchaID, captchaCode) {
		Json(c, e.ECaptcha, resp)
		return
	}

	//ipStr := c.ClientIP()

	err := s.login(c, username, password)

	if err != nil {
		Json(c, e.EUsernameOrPassword, resp)
		return
	}

	Json(c, nil, nil)
}

func (s *Session) login(c *gin.Context, username string, password string) error {
	//判断IP是否限制登录
	//todo

	adminService := &adminservice.Admin{}
	admin, err := adminService.FindByName(username)
	if err != nil {
		return err
	}

	if admin.State != service.StateActive {
		return e.EInvalidAdmin
	}

	if encrypt.CompareHashAndPassword(admin.Password, password) != nil {
		return e.EInvalidAdmin
	}

	roleIDs, err := adminService.RoleIDs(admin.ID)
	if err != nil || len(roleIDs) == 0 {
		return e.ERestrictedAdmin
	}

	lastLoginAt := "-"
	if admin.LoginAt != nil {
		lastLoginAt = admin.LoginAt.Format("2006-01-02 15:03:04")
	}

	ipStr := c.ClientIP()

	session := pkg.NewSession(c)
	session.ID = admin.ID
	session.Name = admin.Name
	if admin.Avatar != "" {
		session.Avatar = fmt.Sprintf("/%s", admin.Avatar)
	}
	session.LoginAt = lastLoginAt
	session.LoginIP = ip.ToIP(admin.LoginIP).String()
	session.RoleIDs = roleIDs
	session.GroupID = admin.GroupID

	log.Println(session.Save())

	//更新最后登录时间和IP
	now := time.Now()
	admin.LoginAt = &now
	admin.LoginIP = ip.ToInt64(net.ParseIP(ipStr))
	adminService.UpdateLoginTimeAndIP(admin)
	return nil
}

func (s *Session) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.Redirect(http.StatusFound, "/login")
}
