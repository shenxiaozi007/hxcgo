package middleware

import (
	"github.com/huangxinchun/hxcgo/admin/app/pkg"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {

	return func(c *gin.Context) {
		admin := pkg.NewSession(c)
		if admin.ID == 0 || admin.GroupID == 0 {
			c.HTML(http.StatusUnauthorized, "error/logout.html", nil)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if !admin.IsGranted(c.Request.URL.Path) {
			//c.HTML(http.StatusUnauthorized, "error/401.html", nil)
			//c.AbortWithStatus(http.StatusUnauthorized)
			//return
		}

		admin.ActiveTime = time.Now().Format("2006-01-02 15:03:04")
		admin.Save()
		c.Next()
	}

}
