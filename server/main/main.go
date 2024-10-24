package main

import (
	"fmt"
	"log"

	"centocode.com/app/api"
	"centocode.com/app/config"
	"github.com/gin-gonic/gin"
)

// @title Example API
// @version 1.0
// @description This is a sample API
// @host localhost:8080
// @BasePath /api/v1
func main() {
	configData, err := config.LoadConfigs("config/config.yaml")

	if err != nil {
		log.Fatal("failed to load config file", err)
	}

	fmt.Println(configData, "config file is loaded")

	router := gin.Default()

	MyRouters := api.Routers{
		Router: router,
	}

	MyRouters.Init()

	serverAddress := configData.App.Host + ":" + configData.App.Port

	err = router.Run(serverAddress)

	if err != nil {
		log.Fatal("unable to start the server.", err)
	}
}
