package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Routers struct {
	Router *gin.Engine
}

func (r *Routers) Init() {

	r.Router.GET("/swagger/*filepath", func(c *gin.Context) {
		httpSwagger.WrapHandler.ServeHTTP(c.Writer, c.Request)
	})

	v1 := r.Router.Group("/api/v1")

	userGroup := userGroup{
		RouterGroup: v1,
	}

	userGroup.Init()

	defer func() {
		fmt.Println("Router has been intialized..")
	}()
}
