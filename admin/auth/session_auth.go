package auth

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type SessionAuthenticator struct {
	sessionValidity time.Duration
}

func NewSessionAuthenticator(sessionValidity time.Duration) *SessionAuthenticator {
	return &SessionAuthenticator{
		sessionValidity: sessionValidity,
	}
}

func (s *SessionAuthenticator) MustHaveSessionLogin(c *gin.Context) (username string, err error) {
	// Extract expiration from session
	session := sessions.Default(c)
	validUntilRaw := session.Get("valid_until")
	if validUntilRaw == nil {
		return "", errors.New("valid_until not present in session")
	}

	// Validate type of expiration
	validUntilString, ok := validUntilRaw.(string)
	if !ok {
		session.Clear()
		if err := session.Save(); err != nil {
			log.Warn().Err(err).Interface("valid_until", validUntilRaw).Msg("Failed to clear invalid session")
			return "", fmt.Errorf("session invalid, but failed to clear session: %w", err)
		}
		return "", errors.New("session invalid and therefore cleared")
	}

	// Parse expiration
	validUntil, timeErr := time.Parse(time.RFC3339, validUntilString)
	if timeErr != nil {
		session.Clear()
		if err := session.Save(); err != nil {
			log.Warn().Err(err).AnErr("time_parse_error", timeErr).Interface("valid_until", validUntilRaw).Msg("Failed to clear invalid session")
			return "", fmt.Errorf("session invalid, but failed to clear session: %w", err)
		}
		return "", errors.New("session invalid and therefore cleared")
	}

	// Clear session if expired
	if time.Now().After(validUntil) {
		session.Clear()
		if err := session.Save(); err != nil {
			log.Warn().Err(err).Time("valid_until", validUntil).Msg("Failed to clear expired session")
			return "", fmt.Errorf("session expired, but failed to clear session: %w", err)
		}
		return "", errors.New("session expired and therefore cleared")
	}

	// Refresh expiration if beyond 10% of expiration time
	if time.Until(validUntil) < s.sessionValidity/10 {
		newValidUntil := time.Now().Add(s.sessionValidity)
		session.Set("valid_until", newValidUntil)
		if err := session.Save(); err != nil {
			log.Warn().Err(err).Time("valid_until_old", validUntil).Time("valid_until_new", newValidUntil).Msg("Failed to extend session")
			return "", fmt.Errorf("failed to extend session: %w", err)
		}
	}

	// Session is valid
	usernameRaw := session.Get("username")
	if usernameRaw == nil {
		return "none", nil
	}
	username, ok = usernameRaw.(string)
	if !ok {
		log.Warn().Err(err).Interface("username", usernameRaw).Msg("Username has unexpected type in session")
		session.Clear()
		if err := session.Save(); err != nil {
			log.Warn().Err(err).Interface("username", usernameRaw).Msg("Username has unexpected type in session, but failed to clear invalid session")
			return "", fmt.Errorf("username in session invalid, but failed to clear session: %w", err)
		}
		return "", errors.New("username in session has unexpected type and therefore session is cleared")
	}
	return username, nil
}

func (s *SessionAuthenticator) StartSession(session sessions.Session, username string) error {
	session.Set("username", username)
	session.Set("valid_until", time.Now().Add(s.sessionValidity).Format(time.RFC3339))
	if err := session.Save(); err != nil {
		log.Warn().Err(err).Str("username", username).Msg("Failed to start session")
		return fmt.Errorf("failed to start session: %w", err)
	}
	return nil
}

func (s *SessionAuthenticator) MW(redirectPathNotLoggedIn string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := s.MustHaveSessionLogin(c); err != nil {
			c.Redirect(http.StatusSeeOther, redirectPathNotLoggedIn)
			c.Abort()
			return
		}
	}
}
