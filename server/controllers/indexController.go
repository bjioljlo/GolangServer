package controllers

import (
	"fmt"
	"net/http"

	"GolangServer/server/models"

	"github.com/gin-gonic/gin"
)

func IndexHome(c *gin.Context) {
	if hasSession := models.HasSession(c); hasSession {
		temp := models.GetSessionValue(models.GetSession(c))
		checkUserIdsExist(temp)

	}
	random_stocks := models.GetBacktestInfo("SearchHistory")

	var tempList = make(map[string]string)
	for count := 0; count < len(random_stocks); count++ {
		tempList["save_number"+fmt.Sprint(count+1)] = random_stocks[count]
		val := models.GetBacktestInfo(tempList["save_number"+fmt.Sprint(count+1)])
		if val == nil {
			continue
		}

		var tomorrow_action models.Tomorrow_Action
		models.JsonToStruck([]byte(val[0]), &tomorrow_action)

		tempList["save_number"+fmt.Sprint(count+1)+"_Date"] = tomorrow_action.Date
		tempList["save_number"+fmt.Sprint(count+1)+"_UnitNumber"] = tomorrow_action.UnitNumber
		tempList["save_number"+fmt.Sprint(count+1)+"_BuyPrice"] = tomorrow_action.BuyPrice
		tempList["save_number"+fmt.Sprint(count+1)+"_SellPrice"] = tomorrow_action.SellPrice
		tempList["save_number"+fmt.Sprint(count+1)+"_Buy_sell"] = tomorrow_action.Buy_sell
		if UserData != nil {
			tempList["session"] = UserData.Username
		}

	}
	c.HTML(http.StatusOK, "index.html", tempList)
}

//tomorrow action的表格回傳
func IndexTData(c *gin.Context) {
	random_stocks := models.GetBacktestInfo("SearchHistory")

	var tempList = make(map[string]string)
	for count := 0; count < len(random_stocks); count++ {
		tempList["save_number"+fmt.Sprint(count+1)] = random_stocks[count]
		val := models.GetBacktestInfo(tempList["save_number"+fmt.Sprint(count+1)])
		if val == nil {
			continue
		}

		var tomorrow_action models.Tomorrow_Action
		models.JsonToStruck([]byte(val[0]), &tomorrow_action)

		tempList["save_number"+fmt.Sprint(count+1)+"_Date"] = tomorrow_action.Date
		tempList["save_number"+fmt.Sprint(count+1)+"_UnitNumber"] = tomorrow_action.UnitNumber
		tempList["save_number"+fmt.Sprint(count+1)+"_BuyPrice"] = tomorrow_action.BuyPrice
		tempList["save_number"+fmt.Sprint(count+1)+"_SellPrice"] = tomorrow_action.SellPrice
		tempList["save_number"+fmt.Sprint(count+1)+"_Buy_sell"] = tomorrow_action.Buy_sell
	}
	c.JSON(http.StatusOK, tempList)
}

//總賺錢的表格回傳
func IndexBData(c *gin.Context) {
	if !models.HasSession(c) {
		c.JSON(http.StatusOK, nil)
		return
	}
	userId := models.GetSessionValue(models.GetSession(c))
	userData, err := models.FindId(userId)
	if err != nil {
		fmt.Println("func IndexSData 查詢不到 id 為 ", userId)
	}
	var stocks []string
	models.JsonToStruck([]byte(userData.Stocks), &stocks)

	var tempList = make(map[string]string)
	for count := 0; count < len(stocks); count++ {
		tempList["save_number"+fmt.Sprint(count+1)] = stocks[count]
		val := models.GetBacktestInfo(tempList["save_number"+fmt.Sprint(count+1)])
		if val == nil {
			continue
		}

		var stock_info models.Stock_Info
		models.JsonToStruck([]byte(val[1]), &stock_info)

		tempList["save_number"+fmt.Sprint(count+1)+"Return"] = stock_info.Return
	}
	c.JSON(http.StatusOK, tempList)
}

//使用者存檔的回傳
func IndexSData(c *gin.Context) {
	if !models.HasSession(c) {
		c.JSON(http.StatusOK, nil)
		return
	}
	userId := models.GetSessionValue(models.GetSession(c))
	userData, err := models.FindId(userId)
	if err != nil {
		fmt.Println("func IndexSData 查詢不到 id 為 ", userId)
	}
	var stocks []string
	models.JsonToStruck([]byte(userData.Stocks), &stocks)
	var tempList = make(map[string]string)
	for count := 0; count < len(stocks); count++ {
		tempList["save_number"+fmt.Sprint(count+1)] = stocks[count]
	}
	c.JSON(http.StatusOK, tempList)
}
