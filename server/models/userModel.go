package models

import (
	"GolangServer/server/drivers"
)

var Users *[]UserInfo

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

func DeleteUser(user *UserInfo) error {
	return drivers.MysqlDB.Delete(user).Error
}

func FindUser(username string) (*UserInfo, error) {
	user := new(UserInfo)
	user.Username = username
	err := drivers.MysqlDB.Where("username = ?", username).First(&user).Error
	return user, err
}

func FindId(id int64) (*UserInfo, error) {
	user := new(UserInfo)
	user.ID = id
	err := drivers.MysqlDB.Where("id = ?", id).First(&user).Error
	return user, err
}
func FindSession(session uint) (*UserInfo, error) {
	user := new(UserInfo)
	err := drivers.MysqlDB.Where("id = ?", session).First(&user).Error
	return user, err
}

func IsNotFoundError(err error) bool {
	return drivers.IsNotFoundError(err)
}
