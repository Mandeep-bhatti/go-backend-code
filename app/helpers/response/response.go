package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendSuccessResponse[T any](c *gin.Context, data T) {
	jsonResponse := gin.H{
		"status": "ok",
		"data":   data,
	}
	c.JSON(http.StatusOK, jsonResponse)
}

func SendBadRequestResponse[T comparable](c *gin.Context, message T) {
	jsonResponse := gin.H{
		"status":  "bad",
		"message": message,
	}
	c.JSON(http.StatusBadRequest, jsonResponse)
}
