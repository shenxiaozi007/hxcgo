package ctrl

import (
	"net/http"
	"strconv"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
)

type Captcha struct {
	allowCaptchaSizes []*ImageSize
}

func NewCaptcha() *Captcha {
	return &Captcha{
		allowCaptchaSizes: []*ImageSize{
			&ImageSize{Height: 34, Width: 90},
		},
	}

}

func (pt *Captcha) Get(c *gin.Context) {
	captchaID := c.Param("id")
	if captchaID == "" {
		c.String(http.StatusNotFound, "404 page not found")
	}

	height, err := strconv.Atoi(c.Param("height"))
	if err != nil || height <= 0 {
		c.String(http.StatusNotFound, "404 page not found")
	}

	width, err := strconv.Atoi(c.Param("width"))
	if err != nil || height <= 0 {
		c.String(http.StatusNotFound, "404 page not found")
	}

	var isAllow bool = false
	for _, img := range pt.allowCaptchaSizes {
		if img.Height == height && img.Width == width {
			isAllow = true
			break
		}
	}

	if !isAllow {
		c.String(http.StatusNotImplemented, "Not Implemented ")
	}

	captcha.Reload(captchaID)
	captcha.WriteImage(c.Writer, captchaID, 90, 34)
}
