package admin

import (
	"net/http"

	"github.com/JenswBE/go-commerce/admin/entities"
	"github.com/JenswBE/go-commerce/admin/i18n"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (h *AdminHandler) handleLogin(c *gin.Context) {
	// Default action (non-POST) is to show the template
	if c.Request.Method != http.MethodPost {
		c.HTML(200, "login", entities.BaseData{Title: "Inloggen"})
		return
	}

	// Handle login on POST
	username := c.PostForm("username")
	password := c.PostForm("password")

}

func (h *AdminHandler) handleLogout(c *gin.Context) {
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

func verifyCredentials(username, password string) bool {
	return false
}
