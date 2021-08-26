package drivers

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlDB *gorm.DB

func RunMysqlDB() {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Viper.GetString("MysqlDB.User"),
		Viper.GetString("MysqlDB.Password"),
		Viper.GetString("MysqlDB.Network"),
		Viper.GetString("MysqlDB.IP"),
		Viper.GetInt("MysqlDB.Port"),
		Viper.GetString("MysqlDB.DB"))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("使用 gorm 連線 DB 發生錯誤，原因為 " + err.Error())
	}
	fmt.Println("MysqlDB OK")

	MysqlDB = db
}
