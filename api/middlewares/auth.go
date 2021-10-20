package middlewares

import (
	"context"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type OIDCMiddleware struct {
	verifier *oidc.IDTokenVerifier
}

func NewOIDCMiddleware(issuerURL string) (*OIDCMiddleware, error) {
	// Get provider
	provider, err := oidc.NewProvider(context.Background(), issuerURL)
	if err != nil {
		return nil, err
	}

	// Build middleware
	config := &oidc.Config{SkipClientIDCheck: true}
	return &OIDCMiddleware{verifier: provider.Verifier(config)}, nil
}

func (mw *OIDCMiddleware) EnforceRoles(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract token
		rawToken := strings.TrimPrefix(c.Request.Header.Get("Authorization"), "Bearer ")

		// Validate token
		token, err := mw.verifier.Verify(c.Request.Context(), rawToken)
		if err != nil {
			log.Debug().Err(err).Msg("VerifyToken: Token is invalid")
			c.AbortWithStatus(401)
			return
		}

		// Extract claims
		var claims struct {
			Roles []string `json:"roles"`
		}
		if err := token.Claims(&claims); err != nil {
			log.Debug().Err(err).Msg("VerifyToken: Failed to extract custom claims")
			c.AbortWithStatus(401)
			return
		}

		// Verify roles
		if !verifyRoles(claims.Roles, roles) {
			log.Debug().Err(err).Strs("expected", roles).Strs("actual", claims.Roles).Msg("VerifyToken: Expected role missing")
			c.AbortWithStatus(401)
			return
		}

		// Verification successful
		c.Next()
	}
}

func verifyRoles(actualRoles, expectedRoles []string) bool {
outer:
	for _, expected := range expectedRoles {
		for _, actual := range actualRoles {
			if actual == expected {
				continue outer
			}
		}
		// Current expected not found in actual roles
		return false
	}
	return true
}
