package adminctrl

import (
	"admin/app/ctrl"
	e "admin/app/err"
	"admin/app/pkg"
	"admin/app/proto/adminproto"
	"admin/app/service"
	"admin/app/service/adminservice"
	"admin/core/ip"
	"admin/core/opt"
	"admin/core/reg"
	"admin/core/uuid"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/disintegration/imaging"

	"github.com/gin-gonic/gin"
)

type Account struct {
	idFactory    *uuid.IDFactory
	avatarSizes  []*ctrl.ImageSize
	adminService *adminservice.Admin
	groupService *adminservice.Group
	roleService  *adminservice.Role
}

func NewAccount() *Account {
	idFactory, err := uuid.New(int64(opt.Config().Node))
	if err != nil {
		panic(fmt.Sprintf("admin err: %s", err.Error()))
	}
	return &Account{
		idFactory: idFactory,
		avatarSizes: []*ctrl.ImageSize{
			&ctrl.ImageSize{Width: 128, Height: 128},
		},
		adminService: adminservice.NewAdmin(),
		groupService: adminservice.NewGroup(),
		roleService:  adminservice.NewRole(),
	}
}

func (a *Account) findAccount(c *gin.Context) (*adminproto.Admin, error) {
	account := c.Query("account")
	if account == "" {
		return nil, e.ERequest
	}

	admin, err := a.adminService.FindByAccount(account)
	if err != nil {
		return nil, ctrl.Error(err)
	}

	state, err := strconv.Atoi(c.Query("state"))
	if err != nil {
		state = -1
	}
	if state >= 0 && admin.State != uint8(state) {
		return &adminproto.Admin{}, nil
	}

	groupID, err := strconv.Atoi(c.Query("group_id"))
	if err != nil {
		groupID = 0
	}
	if groupID > 0 && admin.GroupID != uint(groupID) {
		return &adminproto.Admin{}, nil
	}

	return admin, nil
}

func (a *Account) List(c *gin.Context) {
	list := &adminproto.QueryResp{
		Admins: []*adminproto.Admin{},
	}

	state, err := strconv.Atoi(c.Query("state"))
	if err != nil {
		state = -1
	}

	groupID, err := strconv.Atoi(c.Query("group_id"))
	if err != nil {
		groupID = 0
	}

	account := c.Query("account")
	if account != "" {
		admin, err := a.findAccount(c)
		if err != nil {
			ctrl.Render(c, "error/500.html", gin.H{"errMsg": err.Error()})
			return
		}
		list.Page = 1
		list.Limit = 20
		if admin.ID > 0 {
			list.TotalPage = 1
			list.Count = 1
			list.Admins = append(list.Admins, admin)
		}

	} else {
		page, err := strconv.Atoi(c.Query("page"))

		if err != nil {
			page = 1
		}

		req := &adminproto.QueryReq{
			GroupID: uint(groupID),
			State:   int8(state),
			Page:    uint(page),
			Limit:   20,
		}

		list, err = a.adminService.Query(req)
		if err != nil {
			ctrl.Render(c, "error/500.html", gin.H{"errMsg": err.Error()})
			return
		}
	}

	groups, _ := a.groupService.Actives()

	admins := map[uint]interface{}{}
	for _, v := range list.Admins {

		groupName := ""
		group, ok := groups[v.GroupID]
		if ok {
			groupName = group.Name
		}

		loginAt := "-"
		if v.LoginAt != nil {
			loginAt = v.LoginAt.Format("2006-01-02 15:04:05")
		}

		admins[v.ID] = struct {
			ID        uint
			Name      string
			Email     string
			Mobile    string
			State     uint8
			GroupID   uint
			LoginAt   string
			LoginIP   string
			CreatedAt string
			GroupName string
		}{
			ID:        v.ID,
			Name:      v.Name,
			Email:     v.Email,
			Mobile:    v.Mobile,
			State:     v.State,
			GroupID:   v.GroupID,
			LoginAt:   loginAt,
			LoginIP:   ip.ToIP(v.LoginIP).String(),
			CreatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
			GroupName: groupName,
		}
	}

	roles, _ := a.roleService.Actives()

	ctrl.Render(c, "admin/list.html", gin.H{
		"groups":    groups,
		"admins":    admins,
		"totalPage": list.TotalPage,
		"qPage":     list.Page,
		"qState":    state,
		"qGroupID":  uint(groupID),
		"qAccount":  account,
		"roles":     roles,
	})
}

func (a *Account) Update(c *gin.Context) {
	if c.Request.Method == "POST" {
		id, err := a.getPostID(c)
		if err != nil {
			ctrl.Json(c, err, nil)
			return
		}

		admin, err := a.adminService.FindByID(id)
		if err != nil {
			ctrl.Json(c, ctrl.Error(err), nil)
			return
		}

		oldAvatar := admin.Avatar

		err = a.parseParam(c, admin)
		if err != nil {
			ctrl.Json(c, err, nil)
			return
		}

		err = a.adminService.Update(admin)
		if err != nil {
			ctrl.Json(c, ctrl.Error(err), nil)
			return
		}

		if admin.Avatar != "" && oldAvatar != admin.Avatar {
			a.removeAvatar(strings.Replace(oldAvatar, "_{size}", "", 1))
		}
		ctrl.Json(c, nil, admin)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ctrl.Render(c, "error/507.html", gin.H{"errMsg": ctrl.LanguageMessage("err_invalid_id")})
		return
	}

	admin, err := a.adminService.FindByID(uint(id))
	if err != nil {
		ctrl.Render(c, "error/507.html", gin.H{"errMsg": ctrl.LanguageMessage("err_admin_name_not_exists")})
		return
	}

	groups, _ := a.groupService.Actives()

	avatar := ""
	if admin.Avatar != "" {
		avatar = fmt.Sprintf("/%s", admin.Avatar)
	}

	data := gin.H{
		"id":       admin.ID,
		"groups":   groups,
		"name":     admin.Name,
		"email":    admin.Email,
		"mobile":   admin.Mobile,
		"password": a.adminService.FakePassword(),
		"groupID":  admin.GroupID,
		"avatar":   avatar,
		"isEnable": "",
	}

	if admin.State == service.StateActive {
		data["isEnable"] = "on"
	}

	ctrl.Render(c, "admin/save.html", data)
}

func (a *Account) Get(c *gin.Context) {

}

func (a *Account) Add(c *gin.Context) {
	if c.Request.Method == "POST" {
		admin := &adminproto.Admin{}
		err := a.parseParam(c, admin)
		if err != nil {
			ctrl.Json(c, err, nil)
			return
		}

		err = a.adminService.Add(admin)
		if err != nil {
			ctrl.Json(c, ctrl.Error(err), nil)
			return
		}

		ctrl.Json(c, nil, admin)
		return
	}

	groups, _ := a.groupService.Actives()

	ctrl.Render(c, "admin/save.html", gin.H{
		"groups":  groups,
		"groupID": 0,
	})
}

func (a *Account) getPostID(c *gin.Context) (uint, error) {
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		return 0, e.EInvalidID
	}
	return uint(id), nil
}

func (a *Account) getPostMobile(c *gin.Context) (string, error) {
	mobile := c.PostForm("mobile")
	if len(mobile) > 0 && !reg.Mobile(mobile) {
		return "", e.EMobile
	}
	return mobile, nil
}

func (a *Account) getPostEmail(c *gin.Context) (string, error) {
	email := c.PostForm("email")
	if len(email) > 0 && !reg.Email(email) {
		return "", e.EEmail
	}
	return email, nil
}

func (a *Account) getPostPassword(c *gin.Context) (string, error) {
	password := c.PostForm("password")
	if len(password) < 8 || reg.Alphabet(password) || reg.Number(password) {
		return "", e.EPassword
	}
	passwordConfirm := c.PostForm("passwordConfirm")
	if password != passwordConfirm {
		return "", e.EConfirmPassword
	}

	return password, nil
}

func (a *Account) getPostGroupID(c *gin.Context) (uint, error) {
	groupID, err := strconv.Atoi(c.PostForm("groupID"))
	if err != nil || groupID < 1 {
		return 0, e.EGroup
	}
	return uint(groupID), nil
}

func (a *Account) getPostName(c *gin.Context) (string, error) {
	name := c.PostForm("name")
	if len(name) < 5 || reg.Number(name) {
		return "", e.EAdminName
	}
	if reg.Email(name) {
		return "", e.EAdminNameCanNotEmail
	}

	return name, nil
}

func (a *Account) getPostState(c *gin.Context) (uint8, error) {
	isEnable := c.PostForm("isEnable")
	var state uint8 = 1
	if isEnable == "on" {
		state = 0
	}
	return state, nil
}

func (a *Account) getPostAvatar(c *gin.Context) (string, error) {
	avatar := c.PostForm("avatar")
	if avatar == "" {
		return "", nil
	}

	imageService := &service.Image{}
	uploadDir := fmt.Sprintf("%s/avatar", imageService.Dir())
	err := os.MkdirAll(uploadDir, 0660)
	if err != nil {
		return "", e.EUpload
	}

	filenameID := a.idFactory.String()
	ext := filepath.Ext(filepath.Base(avatar))
	dstFilename := fmt.Sprintf("%s/%s%s", uploadDir, filenameID, ext)
	err = imageService.Thumbnail(avatar, dstFilename, 400, 400, imaging.Lanczos)
	if err != nil {
		return "", e.EUpload
	}

	for _, file := range a.avatarSizes {
		filename := fmt.Sprintf("%s/%s_%dx%d%s", uploadDir, filenameID, file.Width, file.Height, ext)
		err = imageService.Thumbnail(avatar, filename, file.Width, file.Height, imaging.Lanczos)
		if err != nil {
			a.removeAvatar(dstFilename)
			return "", e.EUpload
		}
	}

	return fmt.Sprintf("%s/%s_{size}%s", uploadDir, filenameID, ext), nil
}

func (a *Account) parseBaseParam(c *gin.Context, admin *adminproto.Admin) (err error) {
	admin.Mobile, err = a.getPostMobile(c)
	if err != nil {
		return err
	}

	admin.Email, err = a.getPostEmail(c)
	if err != nil {
		return err
	}

	admin.Password, err = a.getPostPassword(c)
	if err != nil {
		return err
	}

	admin.Avatar, err = a.getPostAvatar(c)
	return err
}

func (a *Account) parseParam(c *gin.Context, admin *adminproto.Admin) (err error) {
	admin.GroupID, err = a.getPostGroupID(c)
	if err != nil {
		return err
	}

	admin.Name, err = a.getPostName(c)
	if err != nil {
		return err
	}

	admin.State, err = a.getPostState(c)
	if err != nil {
		return err
	}

	return a.parseBaseParam(c, admin)
}

//删除头像
func (a *Account) removeAvatar(srcFile string) {
	if srcFile == "" {
		return
	}

	ext := filepath.Ext(srcFile)
	filename := strings.TrimRight(srcFile, ext)
	os.RemoveAll(srcFile)
	for _, file := range a.avatarSizes {
		os.RemoveAll(fmt.Sprintf("%s_%dx%d%s", filename, file.Width, file.Height, ext))
	}
}

func (a *Account) Delete(c *gin.Context) {
	id, err := a.getPostID(c)
	if err != nil {
		ctrl.Json(c, err, nil)
		return
	}

	err = a.adminService.Delete(id)
	ctrl.Json(c, ctrl.Error(err), nil)
}

func (a *Account) Profile(c *gin.Context) {
	sess := pkg.NewSession(c)

	//adminID, ok := session.Get("adminID").(uint)
	//if !ok {
	//	ctrl.Render(c, "error/507.html", gin.H{"errMsg": ctrl.LanguageMessage("err_admin_name_not_exists")})
	//	return
	//}

	admin, err := a.adminService.FindByID(sess.ID)
	if err != nil {
		ctrl.Render(c, "error/507.html", gin.H{"errMsg": ctrl.LanguageMessage("err_admin_name_not_exists")})
		return
	}
	groups, _ := a.groupService.Actives()

	avatar := ""
	if admin.Avatar != "" {
		avatar = fmt.Sprintf("/%s", admin.Avatar)
	}

	groupName := ""
	group, ok := groups[admin.GroupID]
	if ok {
		groupName = group.Name
	}

	data := gin.H{
		"id":          admin.ID,
		"name":        admin.Name,
		"email":       admin.Email,
		"mobile":      admin.Mobile,
		"password":    a.adminService.FakePassword(),
		"groupName":   groupName,
		"avatar":      avatar,
		"lastLoginAt": sess.LoginAt,
		"lastLoginIP": sess.LoginIP,
		"loginIP":     ip.ToIP(admin.LoginIP).String(),
	}

	ctrl.Render(c, "admin/profile.html", data)
}

func (a *Account) UpdateProfile(c *gin.Context) {
	sess := pkg.NewSession(c)

	admin, err := a.adminService.FindByID(sess.ID)
	if err != nil {
		ctrl.Json(c, e.ERequest, nil)
		return
	}
	oldAvatar := admin.Avatar

	err = a.parseBaseParam(c, admin)
	if err != nil {
		ctrl.Json(c, err, nil)
		return
	}

	err = a.adminService.Update(admin)
	if err != nil {
		ctrl.Json(c, ctrl.Error(err), nil)
		return
	}

	if admin.Avatar != "" && oldAvatar != admin.Avatar {
		a.removeAvatar(strings.Replace(oldAvatar, "_{size}", "", 1))
	}
	ctrl.Json(c, nil, admin)
	return
}

func (a *Account) getID(c *gin.Context) (uint, error) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		return 0, e.EInvalidID
	}
	return uint(id), nil
}

func (a *Account) RoleIDs(c *gin.Context) {
	id, err := a.getID(c)
	if err != nil {
		ctrl.Json(c, err, nil)
		return
	}

	roleIDs, err := a.adminService.RoleIDs(id)
	ctrl.Json(c, ctrl.Error(err), roleIDs)
}

func (a *Account) getPostRoleID(c *gin.Context) (uint, error) {
	id, err := strconv.Atoi(c.PostForm("roleID"))
	if err != nil {
		return 0, e.EInvalidRoleID
	}
	return uint(id), nil
}

func (a *Account) AssociateRole(c *gin.Context) {
	id, err := a.getPostID(c)
	if err != nil {
		ctrl.Json(c, err, nil)
		return
	}

	roleID, err := a.getPostRoleID(c)
	if err != nil {
		ctrl.Json(c, err, nil)
		return
	}

	err = a.adminService.AssociateRole(id, roleID)
	ctrl.Json(c, ctrl.Error(err), nil)
}
