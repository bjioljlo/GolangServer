package controllers

import (
	"GolangServer/server/models"
	"errors"
	"fmt"
	"math/rand"
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
	if err := auth(username, password); err == nil {
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
func LoginNew(c *gin.Context) {
	var (
		username string
		password string
	)
	username = c.Query("username")
	if username == "" {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"error": errors.New("必須輸入使用者名稱"),
		})
		return
	}
	password = c.Query("password")
	if password == "" {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"error": errors.New("必須輸入密碼名稱"),
		})
		return
	}
	if err := create(username, password); err == nil {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"success": "建立成功",
		})
		return
	}
}

func create(username string, password string) error {
	if isExist := checkUserIsExist(username); isExist {
		return errors.New("user is already exist!")
	} else {
		var temp models.UserInfo
		temp.Username = username
		temp.Password = password
		temp.ID = int64(rand.Intn(100))
		if isExist := checkUserIdsExist(temp.ID); isExist {
			return errors.New("id is already exist!")
		}
		return models.CreateUser(&temp)
	}
}
func auth(username string, password string) error {
	if isExist := checkUserIsExist(username); isExist {
		return checkPassword(UserData.Password, password)
	} else {
		return errors.New("user is not exist")
	}
}
func checkPassword(p1 string, p2 string) error {
	if p1 == p2 {
		return nil
	} else {
		return errors.New("password is not correct")
	}
}
func checkUserIsExist(username string) bool {
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
func checkUserIdsExist(id int64) bool {
	if user, err := models.FindId(id); err == nil {
		fmt.Println("查詢到 id 為 ", user)
		UserData = user
		return true
	} else {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("沒查詢到 id 為 ", user)
			return false
		}
		panic("查詢 id 失敗，原因為 " + err.Error())
	}
}
