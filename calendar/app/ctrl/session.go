package ctrl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	e "calendar/app/err"
	"calendar/app/proto/userproto"
	"calendar/app/service"
	"calendar/core/opt"
	"calendar/core/session"
	"calendar/core/uuid"

	"github.com/gin-gonic/gin"
)

type Session struct {
	wechatService *service.Wechat
	userService   *service.User
	idFactory     *uuid.IDFactory
}

func NewSession() *Session {
	idFactory, err := uuid.New(int64(opt.Config().Node))
	if err != nil {
		panic(fmt.Sprintf("calendar err: %s", err.Error()))
	}
	return &Session{
		wechatService: service.NewWechat(),
		userService:   service.NewUser(),
		idFactory:     idFactory,
	}
}

func (s *Session) Login(c *gin.Context) {
	type Args struct {
		Code          string `json:"code"`
		IV            string `json:"iv"`
		EncryptedData string `json:"encryptedData"`
	}

	buf, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		Json(c, e.EParam, nil)
		return
	}

	args := &Args{}
	err = json.Unmarshal(buf, args)
	if err != nil || args.Code == "" {
		Json(c, e.EParam, nil)
		return
	}

	wxSession, err := s.wechatService.Session(args.Code)
	if err != nil {
		Json(c, e.ECode, nil)
		return
	}

	data, err := s.wechatService.DecryptData(wxSession.Key, args.IV, args.EncryptedData)
	if err != nil {
		Json(c, Error(err), nil)
		return
	}

	user, err := s.userService.FindByOpenID(wxSession.OpenID)
	if err != nil {
		if CompareError(err, e.ENotFound) == true {
			user = &userproto.User{}
		} else {
			Json(c, e.ERequest, nil)
			return
		}

	}

	user.OpenID = wxSession.OpenID
	user.UnionID = wxSession.UnionID
	user.NickName = data.NickName
	user.Avatar = data.AvatarURL
	user.Gender = data.Gender
	user.Country = data.Country
	user.Province = data.Province
	user.City = data.City

	if user.ID > 0 {
		err = s.userService.Update(user)
	} else {
		err = s.userService.Add(user)
	}

	if err != nil {
		Json(c, Error(err), nil)
		return
	}

	sess := session.Default(c)
	token := sess.Create()
	sess.Set("openID", user.OpenID)
	sess.Set("userID", user.ID)
	sess.Save()

	Json(c, nil, gin.H{
		"token": token,
	})
}

func (s *Session) Logout(c *gin.Context) {

}
