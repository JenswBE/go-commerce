package middlewares

import (
	"context"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
)

type OIDCMiddleware struct {
	verifier *oidc.IDTokenVerifier
}

func NewOIDCMiddleware(issuerURL string, clientID string) (*OIDCMiddleware, error) {
	// Get provider
	provider, err := oidc.NewProvider(context.Background(), issuerURL)
	if err != nil {
		return nil, err
	}

	// Build middleware
	config := &oidc.Config{ClientID: clientID}
	return &OIDCMiddleware{verifier: provider.Verifier(config)}, nil
}

func (mw *OIDCMiddleware) Handle(c *gin.Context) {
	token := strings.TrimPrefix(c.Request.Header.Get("Authorization"), "Bearer ")
	_, err := mw.verifier.Verify(c.Request.Context(), token)
	if err != nil {
		c.AbortWithStatus(401)
		return
	}
	c.Next()
}