package adminctrl

import (
	"github.com/huangxinchun/hxcgo/admin/app/ctrl"
	"github.com/huangxinchun/hxcgo/admin/app/pkg"
	"github.com/huangxinchun/hxcgo/admin/app/proto/adminroleproto"
	"github.com/huangxinchun/hxcgo/admin/app/service"
	"github.com/huangxinchun/hxcgo/admin/app/service/adminservice"
	"encoding/json"
	"fmt"
	"strconv"

	e "github.com/huangxinchun/hxcgo/admin/app/err"

	"github.com/gin-gonic/gin"
)

type Role struct {
	roleService      *adminservice.Role
	privilegeService *adminservice.Privilege
}

func NewRole() *Role {
	return &Role{
		roleService:      adminservice.NewRole(),
		privilegeService: adminservice.NewPrivilege(),
	}
}

func (r *Role) List(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}

	state, err := strconv.Atoi(c.Query("state"))
	if err != nil {
		state = -1
	}

	req := &adminroleproto.QueryReq{
		Page:  uint(page),
		Limit: 20,
		State: int8(state),
	}

	list, err := r.roleService.Query(req)
	if err != nil {
		ctrl.Render(c, "error/500.html", gin.H{"errMsg": ctrl.LanguageMessage(ctrl.Error(err).Error())})
		return
	}

	ctrl.Render(c, "admin_role/list.html", gin.H{
		"roles":     list.Roles,
		"totalPage": list.TotalPage,
		"qPage":     list.Page,
		"qState":    state,
	})
}

func (r *Role) save(c *gin.Context, id uint) {
	role := &adminroleproto.Role{
		ID: id,
	}
	err := r.parseParam(c, role)
	if err != nil {
		ctrl.Render(c, "error/507.html", gin.H{"errMsg": ctrl.LanguageMessage(err.Error())})
		return
	}

	if id > 0 {
		err = r.roleService.Update(role)
	} else {
		err = r.roleService.Add(role)
	}
	if err != nil {
		ctrl.Render(c, "error/500.html", gin.H{"errMsg": ctrl.LanguageMessage(err.Error())})
		return
	}

	ctrl.Render(c, "success/200.html", gin.H{"redirectURL": fmt.Sprintf("/admin/role/update/%d", role.ID)})
}

func (r *Role) Update(c *gin.Context) {
	if c.Request.Method == "POST" {
		id, err := r.getPostID(c)
		if err != nil {
			ctrl.Render(c, "error/507.html", gin.H{"errMsg": ctrl.LanguageMessage(err.Error())})
			return
		}

		r.save(c, id)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ctrl.Render(c, "error/507.html", gin.H{"errMsg": ctrl.LanguageMessage("err_invalid_id")})
		return
	}

	role, err := r.roleService.FindByID(uint(id))
	if err != nil {
		ctrl.Render(c, "error/507.html", gin.H{"errMsg": ctrl.LanguageMessage("err_role_not_exists")})
		return
	}

	data := gin.H{
		"id":       role.ID,
		"name":     role.Name,
		"state":    role.State,
		"isEnable": "",
	}
	if role.State == service.StateActive {
		data["isEnable"] = "on"
	}
	ctrl.Render(c, "admin_role/save.html", data)
}

func (r *Role) Add(c *gin.Context) {
	if c.Request.Method == "POST" {
		r.save(c, 0)
		return
	}

	ctrl.Render(c, "admin_role/save.html", nil)
}

func (r *Role) getPostID(c *gin.Context) (uint, error) {
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		return 0, e.EInvalidID
	}
	return uint(id), nil
}

func (r *Role) getPostName(c *gin.Context) (string, error) {
	name := c.PostForm("name")
	if name == "" {
		return "", e.EInvalidRoleName
	}
	return name, nil
}

func (r *Role) getPostState(c *gin.Context) (uint8, error) {
	isEnable := c.PostForm("isEnable")
	var state uint8 = 1
	if isEnable == "on" {
		state = 0
	}
	return state, nil
}

func (r *Role) parseParam(c *gin.Context, role *adminroleproto.Role) (err error) {
	role.Name, err = r.getPostName(c)
	if err != nil {
		return err
	}

	role.State, err = r.getPostState(c)
	return err
}

func (r *Role) Delete(c *gin.Context) {
	id, err := r.getPostID(c)
	if err != nil {
		ctrl.Json(c, e.EParam, nil)
		return
	}

	err = r.roleService.Delete(id)
	ctrl.Json(c, ctrl.Error(err), nil)
}

func (r *Role) getPostPrivilegeID(c *gin.Context) (uint, error) {
	privilegeID, err := strconv.Atoi(c.PostForm("privilegeID"))
	if err != nil {
		return 0, e.EInvalidPrivilegeID
	}
	return uint(privilegeID), nil
}

func (r *Role) AssociatePrivilege(c *gin.Context) {
	id, err := r.getPostID(c)
	if err != nil {
		ctrl.Json(c, e.EParam, nil)
		return
	}

	privilegeID, err := r.getPostPrivilegeID(c)
	if err != nil {
		ctrl.Json(c, e.EParam, nil)
		return
	}

	err = r.roleService.AssociatePrivilege(id, privilegeID)
	ctrl.Json(c, ctrl.Error(err), nil)
}

func (r *Role) Privileges(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ctrl.Render(c, "error/507.html", gin.H{"errMsg": ctrl.LanguageMessage("err_invalid_id")})
		return
	}

	privilegeTree, err := pkg.PrivilegeTree()
	if err != nil {
		ctrl.Render(c, "error/500.html", gin.H{"errMsg": ctrl.LanguageMessage(ctrl.Error(err).Error())})
		return
	}

	privilegeIDMap, err := r.roleService.PrivilegeIDMap(uint(id))
	if err != nil {
		ctrl.Render(c, "error/500.html", gin.H{"errMsg": ctrl.LanguageMessage(ctrl.Error(err).Error())})
		return
	}

	js, _ := json.Marshal(privilegeTree)
	privilegeIDsJS, _ := json.Marshal(privilegeIDMap)
	ctrl.Render(c, "admin_role/privileges.html", gin.H{
		"privilegesJSON":   js,
		"privilegeIDsJSON": privilegeIDsJS,
		"roleID":           id,
	})
}
