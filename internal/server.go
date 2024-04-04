package server

import (
	"GolangServer/internal/drivers"
	"GolangServer/internal/routes"

	"github.com/gin-gonic/gin"
)

var HttpServer *gin.Engine

func RunServer() {
	drivers.RunMysqlDB()
	drivers.RunRedisDB()

	HttpServer = gin.Default()
	HttpServer.LoadHTMLGlob(drivers.Viper.GetString("GolangServer.ViewsPath"))
	HttpServer.Static("/assets", "./internal/view/assets")
	routes.RegisterRoutes(HttpServer)
	err := HttpServer.Run(drivers.Viper.GetString("GolangServer.Host") + ":" + drivers.Viper.GetString("GolangServer.Port"))
	if err != nil {
		panic("HttpServer error:" + err.Error())
	}
}
