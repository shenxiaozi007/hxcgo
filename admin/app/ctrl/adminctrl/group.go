package adminctrl

import (
	"github.com/huangxinchun/hxcgo/admin/app/ctrl"
	"github.com/huangxinchun/hxcgo/admin/app/proto/admingroupproto"
	"github.com/huangxinchun/hxcgo/admin/app/service"
	"github.com/huangxinchun/hxcgo/admin/app/service/adminservice"
	"fmt"
	"strconv"

	e "github.com/huangxinchun/hxcgo/admin/app/err"

	"github.com/gin-gonic/gin"
)

type Group struct {
	groupService *adminservice.Group
}

//管理组不参与权限分配
func NewGroup() *Group {
	return &Group{
		groupService: adminservice.NewGroup(),
	}
}

func (g *Group) List(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}

	state, err := strconv.Atoi(c.Query("state"))
	if err != nil {
		state = -1
	}

	req := &admingroupproto.QueryReq{
		Page:  uint(page),
		Limit: 20,
		State: int8(state),
	}

	list, err := g.groupService.Query(req)
	if err != nil {
		ctrl.Render(c, "error/500.html", gin.H{"errMsg": ctrl.LanguageMessage(ctrl.Error(err).Error())})
		return
	}

	ctrl.Render(c, "admin_group/list.html", gin.H{
		"groups":    list.Groups,
		"totalPage": list.TotalPage,
		"qPage":     list.Page,
		"qState":    state,
	})
}

func (g *Group) save(c *gin.Context, id uint) {
	group := &admingroupproto.Group{
		ID: id,
	}
	err := g.parseParam(c, group)
	if err != nil {
		ctrl.Render(c, "error/507.html", gin.H{"errMsg": ctrl.LanguageMessage(err.Error())})
		return
	}

	if id > 0 {
		err = g.groupService.Update(group)
	} else {
		err = g.groupService.Add(group)
	}
	if err != nil {
		ctrl.Render(c, "error/500.html", gin.H{"errMsg": ctrl.LanguageMessage(err.Error())})
		return
	}

	ctrl.Render(c, "success/200.html", gin.H{"redirectURL": fmt.Sprintf("/admin/group/update/%d", group.ID)})
}
func (g *Group) Update(c *gin.Context) {
	if c.Request.Method == "POST" {
		id, err := g.getPostID(c)
		if err != nil {
			ctrl.Render(c, "error/507.html", gin.H{"errMsg": ctrl.LanguageMessage(err.Error())})
			return
		}

		g.save(c, id)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ctrl.Render(c, "error/507.html", gin.H{"errMsg": ctrl.LanguageMessage("err_invalid_id")})
		return
	}

	group, err := g.groupService.FindByID(uint(id))
	if err != nil {
		ctrl.Render(c, "error/507.html", gin.H{"errMsg": ctrl.LanguageMessage("err_group")})
		return
	}

	data := gin.H{
		"id":       group.ID,
		"name":     group.Name,
		"state":    group.State,
		"isEnable": "",
	}
	if group.State == service.StateActive {
		data["isEnable"] = "on"
	}
	ctrl.Render(c, "admin_group/save.html", data)
}

func (g *Group) Add(c *gin.Context) {
	if c.Request.Method == "POST" {
		g.save(c, 0)
		return
	}

	ctrl.Render(c, "admin_group/save.html", nil)
}

func (g *Group) getPostID(c *gin.Context) (uint, error) {
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		return 0, e.EInvalidID
	}
	return uint(id), nil
}

func (g *Group) getPostName(c *gin.Context) (string, error) {
	name := c.PostForm("name")
	if name == "" {
		return "", e.EInvalidGroupName
	}
	return name, nil
}

func (g *Group) getPostState(c *gin.Context) (uint8, error) {
	isEnable := c.PostForm("isEnable")
	var state uint8 = 1
	if isEnable == "on" {
		state = 0
	}
	return state, nil
}

func (g *Group) parseParam(c *gin.Context, group *admingroupproto.Group) (err error) {
	group.Name, err = g.getPostName(c)
	if err != nil {
		return err
	}

	group.State, err = g.getPostState(c)
	return err
}

func (g *Group) Delete(c *gin.Context) {
	id, err := g.getPostID(c)
	if err != nil {
		ctrl.Json(c, e.EParam, nil)
		return
	}

	err = g.groupService.Delete(id)
	ctrl.Json(c, ctrl.Error(err), nil)
}
