package drivers

import (
	"errors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlDB *gorm.DB

func RunMysqlDB() error {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Viper.GetString("MysqlDB.User"),
		Viper.GetString("MysqlDB.Password"),
		Viper.GetString("MysqlDB.Network"),
		Viper.GetString("MysqlDB.IP"),
		Viper.GetInt("MysqlDB.Port"),
		Viper.GetString("MysqlDB.DB"))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	fmt.Println("MysqlDB OK")

	MysqlDB = db
	return nil
}

func IsNotFoundError(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
