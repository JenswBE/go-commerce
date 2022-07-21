package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	verifier Verifier
}

func NewAuthMiddleware(verifier Verifier) *AuthMiddleware {
	return &AuthMiddleware{verifier: verifier}
}

func (mw *AuthMiddleware) EnforceRoles(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Validate token
		_, err := mw.verifier.EnforceRoles(c.Request.Context(), roles, c.Request.Header.Get("Authorization"))
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Verification successful
		c.Next()
	}
}
