package middleware

import (
	"calendar/core/session"
	"time"

	"github.com/gin-gonic/gin"
)

func AuthRequired(c *gin.Context) {
	sess := session.Default(c)
	//adminID := session.Get("adminID")
	//adminName := session.Get("adminName")
	//if adminID == nil || adminName == nil {
	//	c.HTML(http.StatusUnauthorized, "error/logout.html", nil)
	//	c.AbortWithStatus(http.StatusUnauthorized)
	//	return
	//}
	sess.Set("activeTime", time.Now().Format("2006-01-02 15:03:04"))
	sess.Save()
}
