package admin

import (
	"net/http"

	"github.com/JenswBE/go-commerce/admin/entities"
	"github.com/JenswBE/go-commerce/admin/i18n"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func handleLogin(c *gin.Context) {
	c.HTML(200, "login", gin.H{
		"title": "Inloggen",
	})
}

func handleLogout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	if err := s.Save(); err != nil {
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to clear session for logout: %v", err)
			return
		}
	}
	redirectWithMessage(c, s, entities.MessageTypeSuccess, i18n.LogoutSuccessful(), "")
}
