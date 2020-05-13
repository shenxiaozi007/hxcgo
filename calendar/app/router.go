package app

import (
	"calendar/app/ctrl"
	"calendar/core/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	router.Use(middleware.AuthRequired)
	session := ctrl.NewSession()
	router.POST("/session", session.Login)
}
