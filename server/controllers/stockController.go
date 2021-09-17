package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"GolangServer/server/models"
)

func StockBacktest(c *gin.Context) {
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
func StockSearch(c *gin.Context) {
	var input = string(c.Query("stock"))
	if input == "" {
		fmt.Println("沒輸入東西")
		return
	}
	models.SendMsg(1, 1, input)
	c.JSON(http.StatusOK, nil)
}
func StockDeleteInfo(c *gin.Context) {
	var input = string(c.Query("stock"))
	if input == "" {
		fmt.Println("沒輸入東西")
		return
	}
	models.SendMsg(2, 1, input)
	c.JSON(http.StatusOK, nil)
}
