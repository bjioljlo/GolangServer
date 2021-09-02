package models

import (
	"GolangServer/server/drivers"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const KEY = "AEN233"

// 使用 Cookie 保存 session
func EnableCookieSession() gin.HandlerFunc {
	store := cookie.NewStore([]byte(KEY))
	return sessions.Sessions("SAMPLE", store)
}

// 註冊和登入要保存seesion
func SaveAuthSession(c *gin.Context, id uint) {
	session := sessions.Default(c)
	session.Set("ID", id+7749)
	err := session.Save()
	if err != nil {
		panic("session error:" + err.Error())
	}
}

//登出要刪除session
func ClearAuthSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}

//Session是否存在
func HasSession(c *gin.Context) bool {
	session := sessions.Default(c)
	if sessionValue := session.Get("ID"); sessionValue == nil {
		return false
	}
	return true
}

//Session中間件
func AuthSessionMidc() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionValue := session.Get("ID")
		if sessionValue == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			c.Abort()
			return
		}
		c.Set("ID", sessionValue.(uint))
		c.Next()
		return
	}
}

//取得session的值
func GetSession(c *gin.Context) uint {
	session := sessions.Default(c)
	sessionValue := session.Get("ID")
	if sessionValue == nil {
		return 0
	}
	return sessionValue.(uint)
}

//取得Session對應資料
func GetSessionValue(session uint) int64 {
	return int64(session - 7749)
	//下面是有存資料在db中用的
	ctx := context.Background()
	val, err := drivers.RedisDB.LRange(ctx, strconv.Itoa(int(session)), 0, -1).Result()
	if err != nil {
		panic(err)
	}
	result, err := strconv.Atoi(val[0])
	if err != nil {
		panic(err)
	}
	return int64(result)
}
