package controllers

import (
	"GolangServer/server/models"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var UserData *models.UserInfo

func LoginPage(c *gin.Context) {
	models.CheckTable()
	c.HTML(http.StatusOK, "login.html", nil)
}

func LoginAuth(c *gin.Context) {
	var (
		username string
		password string
	)

	if in, isExist := c.GetPostForm("username"); isExist && in != "" {
		username = in
	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"error": errors.New("必須輸入使用者名稱"),
		})
		return
	}
	if in, isExist := c.GetPostForm("password"); isExist && in != "" {
		password = in
	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"error": errors.New("必須輸入密碼名稱"),
		})
		return
	}
	if err := Auth(username, password); err == nil {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"success": "登入成功",
		})
		return
	} else {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"error": err,
		})
		return
	}
}

func Auth(username string, password string) error {
	if isExist := CheckUserIsExist(username); isExist {
		return CheckPassword(UserData.Password, password)
	} else {
		return errors.New("user is not exist")
	}
}

func CheckPassword(p1 string, p2 string) error {
	if p1 == p2 {
		return nil
	} else {
		return errors.New("password is not correct")
	}
}
func CheckUserIsExist(username string) bool {
	if user, err := models.FindUser(username); err == nil {
		fmt.Println("查詢到 User 為 ", user)
		UserData = user
		return true
	} else {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("沒查詢到 User 為 ", user)
			return false
		}
		panic("查詢 user 失敗，原因為 " + err.Error())
	}
}
