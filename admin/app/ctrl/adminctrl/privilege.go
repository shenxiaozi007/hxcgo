package adminctrl

import (
	"admin/app/ctrl"
	e "admin/app/err"
	"admin/app/pkg"
	"admin/app/proto/adminprivilegeproto"
	"admin/app/service/adminservice"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Privilege struct {
	privilegeService *adminservice.Privilege
	roleService      *adminservice.Role
}

func NewPrivilege() *Privilege {
	return &Privilege{
		privilegeService: adminservice.NewPrivilege(),
		roleService:      adminservice.NewRole(),
	}
}

func (p *Privilege) List(c *gin.Context) {
	privilegeTree, err := pkg.PrivilegeTree()
	if err != nil {
		ctrl.Render(c, "error/500.html", gin.H{"errMsg": ctrl.LanguageMessage(ctrl.Error(err).Error())})
		return
	}

	js, _ := json.Marshal(privilegeTree)
	roles, _ := p.roleService.Actives()
	ctrl.Render(c, "admin_privilege/list.html", gin.H{
		"privilegesJSON": js,
		"roles":          roles,
	})
}

func (p *Privilege) getPostID(c *gin.Context) (uint, error) {
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		return 0, e.EInvalidID
	}
	return uint(id), nil
}

func (p *Privilege) getPostPID(c *gin.Context) (uint, error) {
	id, err := strconv.Atoi(c.PostForm("pid"))
	if err != nil {
		return 0, e.EInvalidPID
	}
	return uint(id), nil
}

func (p *Privilege) getPostName(c *gin.Context) (string, error) {
	name := c.PostForm("name")
	if name == "" {
		return "", e.EInvalidName
	}
	return name, nil
}

func (p *Privilege) getPostSortOrder(c *gin.Context) (int, error) {
	id, err := strconv.Atoi(c.PostForm("sortOrder"))
	if err != nil {
		id = 0
	}
	return id, nil
}

func (p *Privilege) getPostIsMenu(c *gin.Context) (uint16, error) {
	isMenu, _ := strconv.Atoi(c.PostForm("isMenu"))
	if isMenu != 1 {
		isMenu = 0
	}
	return uint16(isMenu), nil
}

func (p *Privilege) parseParam(c *gin.Context, privilege *adminprivilegeproto.Privilege) (err error) {
	privilege.PID, _ = p.getPostPID(c)
	privilege.IsMenu, _ = p.getPostIsMenu(c)

	privilege.Name, err = p.getPostName(c)
	if err != nil {
		return err
	}

	privilege.SortOrder, err = p.getPostSortOrder(c)
	if err != nil {
		return err
	}

	privilege.Icon = c.PostForm("icon")
	privilege.URIRule = c.PostForm("uriRule")
	return nil
}

func (p *Privilege) save(c *gin.Context, id uint) {
	privilege := &adminprivilegeproto.Privilege{
		ID: id,
	}
	err := p.parseParam(c, privilege)
	if err != nil {
		ctrl.Render(c, "error/507.html", gin.H{"errMsg": ctrl.LanguageMessage(err.Error())})
		return
	}

	if id > 0 {
		if privilege.ID == privilege.PID {
			ctrl.Render(c, "error/507.html", gin.H{"errMsg": ctrl.LanguageMessage(e.EInvalidPID.Error())})
			return
		}
		err = p.privilegeService.Update(privilege)
	} else {
		err = p.privilegeService.Add(privilege)
	}
	if err != nil {
		log.Println(err)
		ctrl.Render(c, "error/500.html", gin.H{"errMsg": ctrl.LanguageMessage(err.Error())})
		return
	}

	ctrl.Render(c, "success/200.html", gin.H{"redirectURL": fmt.Sprintf("/admin/privilege/update/%d", privilege.ID)})
}

func (p *Privilege) Add(c *gin.Context) {
	if c.Request.Method == "POST" {
		p.save(c, 0)
		return
	}

	privileges, _ := p.privilegeService.FindAll()
	js, _ := json.Marshal(privileges)
	ctrl.Render(c, "admin_privilege/save.html", gin.H{
		"isMenu":         0,
		"privilegesJSON": js,
	})
}

func (p *Privilege) Update(c *gin.Context) {
	if c.Request.Method == "POST" {
		id, err := p.getPostID(c)
		log.Println(id)
		if err != nil {
			ctrl.Render(c, "error/507.html", gin.H{"errMsg": ctrl.LanguageMessage(err.Error())})
			return
		}

		p.save(c, id)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	log.Println("param id: ", id)
	if err != nil {
		ctrl.Render(c, "error/507.html", gin.H{"errMsg": ctrl.LanguageMessage("err_invalid_id")})
		return
	}

	privilege, err := p.privilegeService.FindByID(uint(id))
	if err != nil {
		ctrl.Render(c, "error/507.html", gin.H{"errMsg": ctrl.LanguageMessage("err_privilege_not_exists")})
		return
	}

	privileges, _ := p.privilegeService.FindAll()
	js, _ := json.Marshal(privileges)

	var parentName string
	if privilege.PID > 0 {
		parent, err := p.privilegeService.FindByID(privilege.PID)
		if err == nil {
			parentName = parent.Name
		}
	}

	data := gin.H{
		"id":             privilege.ID,
		"isMenu":         privilege.IsMenu,
		"name":           privilege.Name,
		"pid":            privilege.PID,
		"icon":           privilege.Icon,
		"uriRule":        privilege.URIRule,
		"sortOrder":      privilege.SortOrder,
		"pName":          parentName,
		"privilegesJSON": js,
	}

	ctrl.Render(c, "admin_privilege/save.html", data)
}

func (p *Privilege) Delete(c *gin.Context) {
	id, err := p.getPostID(c)
	if err != nil {
		ctrl.Json(c, e.EParam, nil)
		return
	}

	err = p.privilegeService.Delete(id)
	ctrl.Json(c, ctrl.Error(err), nil)
}
func (p *Privilege) getID(c *gin.Context) (uint, error) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		return 0, e.EInvalidID
	}
	return uint(id), nil
}

func (p *Privilege) RoleIDs(c *gin.Context) {
	id, err := p.getID(c)
	if err != nil {
		ctrl.Json(c, err, nil)
		return
	}

	roleIDs, err := p.privilegeService.RoleIDs(id)
	ctrl.Json(c, ctrl.Error(err), roleIDs)
}
