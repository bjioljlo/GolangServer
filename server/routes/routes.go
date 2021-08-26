package routes

import (
	"GolangServer/server/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/", controllers.IndexHome)
	router.GET("/index", controllers.IndexHome)
	router.GET("/login", controllers.LoginPage)
	router.POST("/login", controllers.LoginAuth)
	router.GET("/stock", controllers.Stock_backtest)
}
