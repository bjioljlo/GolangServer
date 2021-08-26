package models

import (
	"GolangServer/server/drivers"
)

type UserInfo struct {
	ID       int64  `json:"id" gorm:"primary_key;auto_increase'"`
	Username string `json:"username"`
	Password string `json:""`
}

func CheckTable() {
	if err := drivers.MysqlDB.AutoMigrate(new(UserInfo)); err != nil {
		panic("資料庫 Migrate 失敗，原因為 " + err.Error())
	}
}

func CreateUser(user *UserInfo) error {
	return drivers.MysqlDB.Create(user).Error
}

func FindUser(username string) (*UserInfo, error) {
	user := new(UserInfo)
	user.Username = username
	err := drivers.MysqlDB.Where("username = ?", username).First(&user).Error
	return user, err
}