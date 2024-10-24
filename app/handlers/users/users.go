package users

import (
	"centocode.com/app/helpers/response"
	"github.com/gin-gonic/gin"
)

func UserGetHandler(c *gin.Context) {

	response.SendSuccessResponse(c, map[string]string{
		"name": "mandeep",
	})
}
