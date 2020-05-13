package ctrl

import "github.com/gin-gonic/gin"

type User struct {
}

func NewUser() *User {
	return &User{}
}

func (u *User) List(c *gin.Context) {

}
