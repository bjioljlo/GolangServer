package server

import (
	"GolangServer/server/drivers"
	_ "GolangServer/server/drivers"
	"GolangServer/server/routes"

	"github.com/gin-gonic/gin"
)

var HttpServer *gin.Engine

func RunServer() {
	drivers.RunMysqlDB()
	drivers.RunRedisDB()

	HttpServer = gin.Default()
	HttpServer.LoadHTMLGlob(drivers.Viper.GetString("GolangServer.ViewsPath"))
	HttpServer.Static("/assets", "./server/view/assets")
	routes.RegisterRoutes(HttpServer)
	err := HttpServer.Run(drivers.Viper.GetString("GolangServer.Host") + ":" + drivers.Viper.GetString("GolangServer.Port"))
	if err != nil {
		panic("HttpServer error:" + err.Error())
	}
}
