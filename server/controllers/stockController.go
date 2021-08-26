package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"GolangServer/server/models"
)

func Stock_backtest(c *gin.Context) {
	input := c.Query("stock")
	//取得DB回測存檔
	val := models.GetBacktestInfo(input)
	//DB的回測資訊
	var stock_info models.Stock_Info
	models.JsonToStruck([]byte(val[1]), &stock_info)
	//DB的行動資訊
	var tomorrow_action models.Tomorrow_Action
	models.JsonToStruck([]byte(val[0]), &tomorrow_action)
	//DB的html資料
	models.SaveHtml([]byte(val[2]), input)
	c.HTML(http.StatusOK, input+".html", tomorrow_action)
}
