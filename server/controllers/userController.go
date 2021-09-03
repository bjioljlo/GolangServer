package controllers

import (
	"GolangServer/server/models"
	"errors"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

var UserData *models.UserInfo

func LoginPage(c *gin.Context) {
	models.CheckTable()
	if hasSession := models.HasSession(c); hasSession {
		temp := models.GetSessionValue(models.GetSession(c))
		checkUserIdsExist(temp)
		c.HTML(http.StatusOK, "login.html", gin.H{
			"success":  "已經登入",
			"UserData": UserData,
		})
		return

	}
	c.HTML(http.StatusOK, "login.html", nil)
}
func LoginAuth(c *gin.Context) {
	var (
		username string
		password string
	)
	username, password, err := checkFormat(c)
	if err != nil {
		return
	}
	if hasSession := models.HasSession(c); hasSession {
		temp, err := models.FindUser(username)
		if err == nil {
			if checkSession(c, temp.ID) {
				checkUserIdsExist(temp.ID)
				c.HTML(http.StatusOK, "login.html", gin.H{
					"success":  "已經登入",
					"UserData": UserData,
				})
				return
			} else {
				models.ClearAuthSession(c)
			}
		}
	}

	if err := auth(username, password); err == nil {
		models.SaveAuthSession(c, uint(UserData.ID))
		c.HTML(http.StatusOK, "login.html", gin.H{
			"success":  "登入成功",
			"UserData": UserData,
		})
		return
	} else {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"error": err,
		})
		return
	}
}
func LoginLogout(c *gin.Context) {
	if hasSession := models.HasSession(c); hasSession {
		models.ClearAuthSession(c)
		UserData = nil
		c.HTML(http.StatusOK, "login.html", gin.H{
			"success": "已經登出",
		})
		return
	} else {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"success": "沒帳號啦",
		})
	}
}
func LoginNew(c *gin.Context) {
	var (
		username string
		password string
	)
	username, password, err := checkFormat(c)
	if err != nil {
		return
	}
	if err := create(username, password); err == nil {
		models.SaveAuthSession(c, uint(UserData.ID))
		c.HTML(http.StatusOK, "login.html", gin.H{
			"success":  "建立成功",
			"UserData": UserData,
		})
		return
	} else {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"error": err,
		})
		return
	}
}
func LoginDel(c *gin.Context) {
	var (
		username string
		password string
	)
	username, password, err := checkFormat(c)
	if err != nil {
		return
	}
	if err := delete(username, password); err == nil {
		if hasSession := models.HasSession(c); hasSession {
			temp, err := models.FindUser(username)
			if err == nil {
				if checkSession(c, temp.ID) {
					models.ClearAuthSession(c)
					UserData = nil
				}
			}
		}

		c.HTML(http.StatusOK, "login.html", gin.H{
			"success": "刪除成功",
		})
		return
	} else {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"error": err,
		})
		return
	}
}
func CheckMe(c *gin.Context) {
	currentUser := c.MustGet("ID").(uint)
	temp, err := models.FindId(int64(currentUser))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"sessionData": currentUser,
			"error":       err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"sessionData": currentUser,
		"userData":    temp,
	})
}

func checkFormat(c *gin.Context) (string, string, error) {
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
		return username, password, errors.New("必須輸入使用者名稱")
	}
	if in, isExist := c.GetPostForm("password"); isExist && in != "" {
		password = in
	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"error": errors.New("必須輸入密碼名稱"),
		})
		return username, password, errors.New("必須輸入密碼名稱")
	}
	return username, password, nil
}
func delete(username string, password string) error {
	if isExist := checkUserIsExist(username); isExist {
		err := checkPassword(UserData.Password, password)
		if err != nil {
			return err
		}
		return models.DeleteUser(UserData)
	} else {
		return errors.New("user is not exist!")
	}
}
func create(username string, password string) error {
	if isExist := checkUserIsExist(username); isExist {
		return errors.New("user is already exist!")
	} else {
		var temp models.UserInfo
		for {
			temp.Username = username
			temp.Password = password
			temp.ID = int64(rand.Intn(100))
			if isExist := checkUserIdsExist(temp.ID); isExist {
				continue //errors.New("id is already exist!")
			}
			break
		}
		UserData = &temp
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
		if models.IsNotFoundError(err) {
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
		if models.IsNotFoundError(err) {
			fmt.Println("沒查詢到 id 為 ", user)
			return false
		}
		panic("查詢 id 失敗，原因為 " + err.Error())
	}
}
func checkSession(c *gin.Context, id int64) bool {
	userId := models.GetSessionValue(models.GetSession(c))
	if userId == 0 || uint(userId) != uint(id) {
		return false
	}
	return true
}

func addStocks(addstock string) {
	var stocks []string
	models.JsonToStruck([]byte(UserData.Stocks), &stocks)
	for i := 0; i < len(stocks); i++ {
		if stocks[i] == addstock {
			fmt.Println("已經重複存檔： ", addstock)
			return
		}
	}
	stocks = append(stocks, addstock)
	val := models.StruckToJson(stocks)
	UserData.Stocks = string(val)
}

func UpdateStocks(c *gin.Context) {
	input := c.Query("stock")
	if input == "" {
		fmt.Println("沒輸入東西")
		c.Redirect(http.StatusMovedPermanently, "/index")
		return
	}
	addStocks(input)
	models.SaveStocks(UserData)
	c.Redirect(http.StatusMovedPermanently, "/index")
}
