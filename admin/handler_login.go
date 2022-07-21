package admin

import (
	"github.com/gin-gonic/gin"
)

func handleLogin(c *gin.Context) {
	c.HTML(200, "login", gin.H{
		"title": "Inloggen",
	})
}
