package admin

import (
	"fmt"
	"net/http"
	"time"

	"github.com/JenswBE/go-commerce/admin/entities"
	"github.com/JenswBE/go-commerce/admin/i18n"
	"github.com/JenswBE/go-commerce/utils/auth"
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
	err := h.authVerifier.ValidateCredentialsWithRoles(c.Request.Context(), c.PostForm("username"), c.PostForm("password"), []string{auth.RoleAdmin})
	if err != nil {
		handleLoginFailed(c, http.StatusUnauthorized, "Inloggen mislukt", err)
		return
	}

	// Login successful => Set token in session
	err = setNewTokenInSession(c, h.jwtSigningKey)
	if err != nil {
		handleLoginFailed(c, http.StatusInternalServerError, "Token toevoegen aan sessie mislukt", err)
		return
	}

	// Redirect to home
	c.Redirect(http.StatusSeeOther, PrefixAdmin)
}

func setNewTokenInSession(c *gin.Context, jwtSigningKey [64]byte) error {
	// Login successful => Set token in session
	token, err := auth.GenerateJWT(jwtSigningKey, time.Hour*24*7)
	if err != nil {
		return fmt.Errorf("aanmaken van token mislukt: %w", err)
	}

	// Set token in session
	session := sessions.Default(c)
	session.Set("token", token)
	err = session.Save()
	if err != nil {
		return fmt.Errorf("opslaan van token in sessie mislukt: %w", err)
	}
	return nil
}

func handleLoginFailed(c *gin.Context, status int, message string, err error) {
	c.HTML(status, "login", entities.BaseData{Title: "Inloggen", Messages: []entities.Message{{
		Type:    entities.MessageTypeError,
		Content: fmt.Sprintf("%s: %v", message, err.Error()),
	}}})
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
