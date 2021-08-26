package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"GolangServer/server/drivers"
)

func Stock_backtest(c *gin.Context) {
	var ctx = context.Background()
	input := c.Query("stock")

	val, err := drivers.RedisDB.LRange(ctx, input, 0, -1).Result() // => GET key
	if err != nil {
		panic(err)
	}

	fmt.Println("key", val[0])
	type Stock_Info struct {
		Equity_Final string `json:"Equity Final [$]"`
		Max_Drawdown string `json:"Max. Drawdown [%]"`
		Win_Rate     string `json:"Win Rate [%]"`
		SQN          string `json:"SQN"`
	}
	type Tomorrow_Action struct {
		UnitNumber string `json:"UnitNumber"`
		BuyPrice   string `json:"BuyPrice"`
		SellPrice  string `json:"SellPrice"`
		Buy_sell   string `json:"Buy_sell"`
	}
	var jsonBlob = []byte(val[1])
	var stock_info Stock_Info
	err2 := json.Unmarshal(jsonBlob, &stock_info)
	if err2 != nil {
		fmt.Println("error:", err2)
	}

	jsonBlob = []byte(val[0])
	var tomorrow_action Tomorrow_Action
	err3 := json.Unmarshal(jsonBlob, &tomorrow_action)
	if err3 != nil {
		fmt.Println("error:", err3)
	}
	c.HTML(http.StatusOK, input+".html", tomorrow_action)
}
