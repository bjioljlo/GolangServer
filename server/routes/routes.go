package routes

import (
	"GolangServer/server/controllers"
	"GolangServer/server/drivers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {

	sindex := router.Group("/", drivers.EnableCookieSession())
	{
		sindex.GET("/", controllers.IndexHome)
		sindex.GET("/index", controllers.IndexHome)
		sindex.GET("/login", controllers.LoginPage)
		sindex.GET("/stock", controllers.StockBacktest)
		sr := sindex.Group("/login")
		{
			sr.POST("/create", controllers.LoginNew)
			sr.POST("/login", controllers.LoginAuth)
			sr.POST("/delete", controllers.LoginDel)
			sr.POST("/logout", controllers.LoginLogout)
			checkUser := sr.Group("/user", drivers.AuthSessionMidc())
			{
				checkUser.POST("/me", controllers.CheckMe)
			}
		}
	}

}
