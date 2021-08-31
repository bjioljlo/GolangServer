package drivers

import (
	"net/http"

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
	session.Set("ID", id)
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

//用session取得userid
func GetUserId(c *gin.Context) uint {
	session := sessions.Default(c)
	sessionValue := session.Get("ID")
	if sessionValue == nil {
		return 0
	}
	return sessionValue.(uint)
}
