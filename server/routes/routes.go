package routes

import (
	"GolangServer/server/controllers"
	"GolangServer/server/models"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {

	sindex := router.Group("/", models.EnableCookieSession())
	{
		sindex.GET("/", controllers.IndexHome)
		sindex.GET("/index", controllers.IndexHome)
		sindex.GET("/Tdata", controllers.IndexTData) //ajax
		sindex.GET("/Bdata", controllers.IndexBData) //ajax
		sindex.GET("/Sdata", controllers.IndexSData) //ajax
		sindex.GET("/login", controllers.LoginPage)
		sindex.GET("/stock", controllers.StockBacktest)
		sindex.GET("/save", controllers.UpdateAddStocks)       //ajax
		sindex.GET("/delet", controllers.UpdateDeletStocks)    //ajax
		sindex.GET("/search", controllers.StockSearch)         //ajax
		sindex.GET("/deleteInfo", controllers.StockDeleteInfo) //ajax
		sr := sindex.Group("/login")
		{
			sr.POST("/create", controllers.LoginNew)
			sr.POST("/login", controllers.LoginAuth)
			sr.POST("/delete", controllers.LoginDel)
			sr.POST("/logout", controllers.LoginLogout)
			checkUser := sr.Group("/user", models.AuthSessionMidc())
			{
				checkUser.POST("/me", controllers.CheckMe)
			}
		}
	}

}
