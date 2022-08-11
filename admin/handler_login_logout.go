package admin

import (
	"fmt"
	"net/http"

	"github.com/JenswBE/go-commerce/admin/auth"
	"github.com/JenswBE/go-commerce/admin/entities"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (h *AdminHandler) handleLogin(c *gin.Context) {
	// Check if authentication is enabled
	if h.authVerifier == nil {
		// Authentication disabled => Redirect to home
		log.Debug().Msg("Authentication disabled, redirecting to home")
		c.Redirect(http.StatusSeeOther, PrefixAdmin)
		return
	}

	// Check if already logged in
	if h.sessionAuthenticator != nil {
		_, err := h.sessionAuthenticator.MustHaveSessionLogin(c)
		if err == nil {
			// Already logged in => Redirect to home
			log.Debug().Msg("Already logged in, redirecting to home")
			c.Redirect(http.StatusSeeOther, PrefixAdmin)
			return
		}
	}

	// Default action (non-POST) is to show the template
	if c.Request.Method != http.MethodPost {
		c.HTML(200, "login", entities.BaseData{Title: "Inloggen"})
		return
	}

	// Handle login on POST
	username, err := h.authVerifier.ValidateCredentialsWithRoles(c.Request.Context(), c.PostForm("username"), c.PostForm("password"), []string{auth.RoleAdmin})
	if err != nil {
		handleLoginFailed(c, http.StatusUnauthorized, "Inloggen mislukt", err)
		return
	}

	// Login successful => Start session
	err = h.sessionAuthenticator.StartSession(sessions.Default(c), username)
	if err != nil {
		handleLoginFailed(c, http.StatusInternalServerError, "Aanmaken sessie mislukt", err)
		return
	}

	// Redirect to home
	c.Redirect(http.StatusSeeOther, PrefixAdmin)
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
	c.Redirect(http.StatusSeeOther, PrefixAdmin)
}
