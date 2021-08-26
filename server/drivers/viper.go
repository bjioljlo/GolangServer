package drivers

import (
	"fmt"

	"github.com/spf13/viper"
)

var Viper *viper.Viper

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./server/config")
	viper.SetDefault("GolangServer.IP", "localhost")
	err := viper.ReadInConfig()
	if err != nil {
		panic("讀取設定檔出現錯誤，原因為：" + err.Error())
	}
	fmt.Println("讀取設定檔成功")
	Viper = viper.GetViper()
}
