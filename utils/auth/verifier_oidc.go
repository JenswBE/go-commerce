package auth

import (
	"context"
	"errors"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/rs/zerolog/log"
)

var _ Verifier = &OIDCVerifier{}

type OIDCVerifier struct {
	verifier *oidc.IDTokenVerifier
}

func NewOIDCVerifier(issuerURL string) (*OIDCVerifier, error) {
	// Get provider
	provider, err := oidc.NewProvider(context.Background(), issuerURL)
	if err != nil {
		return nil, err
	}

	// Build middleware
	config := &oidc.Config{SkipClientIDCheck: true}
	return &OIDCVerifier{verifier: provider.Verifier(config)}, nil
}

func (v *OIDCVerifier) EnforceRoles(ctx context.Context, roles []string, rawToken string) (string, error) {
	// Validate token
	rawToken = strings.TrimPrefix(rawToken, "Bearer ")
	token, err := v.verifier.Verify(ctx, rawToken)
	if err != nil {
		log.Debug().Err(err).Msg("VerifyToken: Token is invalid")
		return "", err
	}

	// Extract claims
	var claims struct {
		Roles []string `json:"roles"`
	}
	if err := token.Claims(&claims); err != nil {
		log.Debug().Err(err).Msg("VerifyToken: Failed to extract custom claims")
		return "", err
	}

	// Verify roles
	if !verifyRoles(claims.Roles, roles) {
		log.Debug().Strs("expected", roles).Strs("actual", claims.Roles).Msg("VerifyToken: Expected role missing")
		return "", errors.New("expected roles missing")
	}
	return token.Subject, nil
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
