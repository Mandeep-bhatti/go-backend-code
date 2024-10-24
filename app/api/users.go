package api

import (
	"fmt"

	"centocode.com/app/handlers/users"
	"github.com/gin-gonic/gin"
)

type userGroup struct {
	RouterGroup *gin.RouterGroup
}

func (r *userGroup) Init() {
	defer func() {
		fmt.Println("users api has been intillized")
	}()
	r.RouterGroup.GET("/users/get", users.UserGetHandler)
}
