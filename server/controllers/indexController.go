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

	} else {

	}
	random_stocks := models.GetBacktestInfo("SearchHistory")

	var tempList map[string]string
	tempList = make(map[string]string)
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

func IndexTData(c *gin.Context) {
	random_stocks := models.GetBacktestInfo("SearchHistory")

	var tempList map[string]string
	tempList = make(map[string]string)
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
func IndexBData(c *gin.Context) {
	random_stocks := models.GetBacktestInfo("SearchHistory")

	var tempList map[string]string
	tempList = make(map[string]string)
	for count := 0; count < len(random_stocks); count++ {
		tempList["save_number"+fmt.Sprint(count+1)] = random_stocks[count]
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
