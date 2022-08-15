package admin

import (
	"fmt"
	"net/http"
	"time"

	"github.com/JenswBE/go-commerce/admin/auth"
	"github.com/JenswBE/go-commerce/admin/entities"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (h *Handler) handleLogin(c *gin.Context) {
	// Check if authentication is enabled
	if h.oidcClient == nil {
		// Authentication disabled => Redirect to home
		log.Debug().Msg("Authentication disabled, redirecting to home")
		redirect(c, "")
		return
	}

	// Check if already logged in
	if h.sessionAuthenticator != nil {
		_, err := h.sessionAuthenticator.MustHaveSessionLogin(c)
		if err == nil {
			// Already logged in => Redirect to home
			log.Debug().Msg("Already logged in, redirecting to home")
			redirect(c, "")
			return
		}
	}

	// Redirect to Auth Code URL
	redirectURLScheme := c.Request.Header.Get("X-Forwarded-Proto")
	if redirectURLScheme == "" {
		if c.Request.TLS == nil {
			redirectURLScheme = "http"
		} else {
			redirectURLScheme = "https"
		}
	}
	redirectURL := fmt.Sprintf("%s://%s%slogin/oidc_redirect/", redirectURLScheme, c.Request.Host, PrefixAdmin)
	url, state := h.oidcClient.GetAuthCodeURL(redirectURL)
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("oidc_state", state, int((time.Minute * 5).Seconds()), "", "", true, true)
	c.Redirect(http.StatusSeeOther, url)
}

func (h *Handler) handleLoginOIDCRedirect(c *gin.Context) {
	// Init
	// session := sessions.Default(c)
	queryCode := c.Query("code")
	queryState := c.Query("state")
	cookieState, err := c.Cookie("oidc_state")
	if err != nil {
		handleLoginFailed(c, http.StatusUnauthorized, "Cookie met state niet gevonden", err)
		return
	}

	// Handle login redirect
	username, err := h.oidcClient.ValidateCallbackWithRoles(c.Request.Context(), cookieState, queryState, queryCode, []string{auth.RoleAdmin})
	if err != nil {
		handleLoginFailed(c, http.StatusUnauthorized, "Inloggen mislukt", err)
		return
	}

	// Login successful => Start session
	// Currently, authentication is also based on session.
	// This is a first iteration and saves us from dealing with refreshing OAuth2 tokens.
	// Might be changed in the future to use the provided OAuth2 token as well including refreshing.
	err = h.sessionAuthenticator.StartSession(sessions.Default(c), username)
	if err != nil {
		handleLoginFailed(c, http.StatusInternalServerError, "Aanmaken sessie mislukt", err)
		return
	}

	// Redirect to home
	c.Redirect(http.StatusFound, PrefixAdmin)
}

func handleLoginFailed(c *gin.Context, status int, message string, err error) {
	template := &entities.LoginFailedTemplate{
		BaseData: entities.BaseData{Title: "Inloggen"},
		Reason:   message,
	}
	if err != nil {
		template.Reason = fmt.Sprintf("%s (%v)", message, err)
	}
	html(c, status, template)
}

func (h *Handler) handleLogout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	if err := s.Save(); err != nil {
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to clear session for logout: %v", err)
			return
		}
	}
	html(c, http.StatusOK, &entities.LogoutSuccessfulTemplate{
		BaseData: entities.BaseData{Title: "Successvol afgemeld"},
	})
}
