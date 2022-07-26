package auth

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/rs/zerolog/log"
	"golang.org/x/oauth2"
)

var _ Verifier = &OIDCVerifier{}

type OIDCVerifier struct {
	clientID string
	provider *oidc.Provider
	verifier *oidc.IDTokenVerifier
}

func NewOIDCVerifier(issuerURL, clientID string) (*OIDCVerifier, error) {
	// Get provider
	provider, err := oidc.NewProvider(context.Background(), issuerURL)
	if err != nil {
		return nil, err
	}

	// Build middleware
	return &OIDCVerifier{
		clientID: clientID,
		provider: provider,
		verifier: provider.Verifier(&oidc.Config{SkipClientIDCheck: true}),
	}, nil
}

func (v *OIDCVerifier) ValidateCredentialsWithRoles(ctx context.Context, username, password string, roles []string) error {
	// To keep things simple, we will just use the Direct/Password Grant flow.
	// We might improve later by using e.g. a Token Grant flow.
	oauth2Config := oauth2.Config{
		ClientID: v.clientID,
		Endpoint: v.provider.Endpoint(),
		Scopes:   []string{oidc.ScopeOpenID},
	}
	token, err := oauth2Config.PasswordCredentialsToken(ctx, username, password)
	if err != nil {
		return fmt.Errorf("failed to validate credentials: %w", err)
	}

	// Validate roles
	_, err = v.EnforceRoles(ctx, roles, token.AccessToken)
	if err != nil {
		return fmt.Errorf("failed to enforce roles %v: %w", roles, err)
	}
	return nil
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
