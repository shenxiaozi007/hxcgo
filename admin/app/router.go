package app

import (
	"github.com/huangxinchun/hxcgo/admin/app/ctrl"
	"github.com/huangxinchun/hxcgo/admin/app/ctrl/adminctrl"
	"github.com/huangxinchun/hxcgo/admin/app/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	router.FuncMap["isGranted"] = IsGranted
	session := ctrl.NewSession()

	router.GET("/login", session.Login)
	router.GET("/logout", session.Logout)
	router.POST("/session", session.Authentication)

	captcha := ctrl.NewCaptcha()
	router.GET("/captcha/:width/:height/:id", captcha.Get)

	authorized := router.Group("/")
	authorized.Use(middleware.AuthRequired())

	image := ctrl.NewImage()
	authorized.HEAD("/image/upload", image.Empty)
	authorized.OPTIONS("/image/upload", image.Empty)
	authorized.POST("/image/upload", image.Upload)

	account := adminctrl.NewAccount()
	authorized.GET("/admin/profile", account.Profile)
	authorized.POST("/admin/profile/update", account.UpdateProfile)
	authorized.GET("/admin/accounts", account.List)
	authorized.GET("/admin/account/update/:id", account.Update)
	authorized.POST("/admin/account/update", account.Update)
	authorized.GET("/admin/account//add", account.Add)
	authorized.POST("/admin/account/add", account.Add)
	authorized.POST("/admin/account/delete", account.Delete)
	authorized.POST("/admin/account/associate_role", account.AssociateRole)
	authorized.GET("/admin/account/role_ids", account.RoleIDs)

	adminGroup := adminctrl.NewGroup()
	authorized.GET("/admin/groups", adminGroup.List)
	authorized.GET("/admin/group/update/:id", adminGroup.Update)
	authorized.POST("/admin/group/update", adminGroup.Update)
	authorized.GET("/admin/group/add", adminGroup.Add)
	authorized.POST("/admin/group/add", adminGroup.Add)
	authorized.POST("/admin/group/delete", adminGroup.Delete)

	adminRole := adminctrl.NewRole()
	authorized.GET("/admin/roles", adminRole.List)
	authorized.GET("/admin/role/update/:id", adminRole.Update)
	authorized.POST("/admin/role/update", adminRole.Update)
	authorized.GET("/admin/role/add", adminRole.Add)
	authorized.POST("/admin/role/add", adminRole.Add)
	authorized.POST("/admin/role/delete", adminRole.Delete)
	authorized.POST("/admin/role/associate_privilege", adminRole.AssociatePrivilege)
	authorized.GET("/admin/role/privileges/:id", adminRole.Privileges)

	adminPrivilege := adminctrl.NewPrivilege()
	authorized.GET("admin/privileges", adminPrivilege.List)
	authorized.GET("/admin/privilege/update/:id", adminPrivilege.Update)
	authorized.POST("/admin/privilege/update", adminPrivilege.Update)
	authorized.GET("/admin/privilege/add", adminPrivilege.Add)
	authorized.POST("/admin/privilege/add", adminPrivilege.Add)
	authorized.POST("/admin/privilege/delete", adminPrivilege.Delete)
	authorized.GET("/admin/privilege/role_ids", adminPrivilege.RoleIDs)
}
