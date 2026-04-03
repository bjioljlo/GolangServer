package server

import (
	"GolangServer/internal/drivers"
	"GolangServer/internal/routes"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

var HttpServer *gin.Engine

func RunServer() {
	var err error
	for {
		err = drivers.RunMysqlDB()
		if err == nil {
			break
		}
		fmt.Printf("mysql connect error: %v, retry in 5 seconds...\n", err)
		time.Sleep(5 * time.Second)
	}

	for {
		err = drivers.RunRedisDB()
		if err == nil {
			break
		}
		fmt.Printf("redis connect error: %v, retry in 5 seconds...\n", err)
		time.Sleep(5 * time.Second)
	}

	HttpServer = gin.Default()
	HttpServer.LoadHTMLGlob(drivers.Viper.GetString("GolangServer.ViewsPath"))
	HttpServer.Static("/assets", "./internal/view/assets")
	routes.RegisterRoutes(HttpServer)
	httpErr := HttpServer.Run(drivers.Viper.GetString("GolangServer.Host") + ":" + drivers.Viper.GetString("GolangServer.Port"))
	if httpErr != nil {
		panic("HttpServer error:" + httpErr.Error())
	}
}
