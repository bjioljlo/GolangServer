package models

import (
	"GolangServer/server/drivers"
	"context"
	"encoding/json"
	"fmt"
	"os"
)

type Stock_Info struct {
	Equity_Final string `json:"Equity Final [$]"`
	Max_Drawdown string `json:"Max. Drawdown [%]"`
	Win_Rate     string `json:"Win Rate [%]"`
	SQN          string `json:"SQN"`
	Return       string `json:"Return [%]"`
}
type Tomorrow_Action struct {
	Date       string `json:"Date"`
	UnitNumber string `json:"UnitNumber"`
	BuyPrice   string `json:"BuyPrice"`
	SellPrice  string `json:"SellPrice"`
	Buy_sell   string `json:"Buy_sell"`
}

func GetBacktestInfo(name string) []string {
	ctx := context.Background()
	val, err := drivers.RedisDB.LRange(ctx, name, 0, -1).Result() // => GET key
	if err != nil {
		panic(err)
	}
	return val
}

func JsonToStruck(val []byte, v interface{}) {
	err := json.Unmarshal(val, v)
	if err != nil {
		fmt.Println("error:", err)
	}
}

func StruckToJson(v interface{}) []byte {
	val, err := json.Marshal(v)
	if err != nil {
		fmt.Println("error:", err)
	}
	return val
}

func SaveHtml(val []byte, name string) {
	f, err := os.Create("./server/view/html/" + name + ".html")
	defer f.Close()
	if err != nil {
		fmt.Println("error4:", err.Error())
	} else {
		_, err5 := f.Write(val)
		if err5 != nil {
			fmt.Println("error5", err5.Error())
		}
	}
}
