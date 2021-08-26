package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHome(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
