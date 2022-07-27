package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/rs/zerolog/log"
	"golang.org/x/oauth2"
)

type OIDCVerifier struct {
	clientID string
	endpoint oauth2.Endpoint
	verifier *oidc.IDTokenVerifier
}

var _ Verifier = &OIDCVerifier{} // Enforce interface

func NewOIDCVerifier(issuerURL, clientID string) (*OIDCVerifier, error) {
	// Get provider
	provider, err := oidc.NewProvider(context.Background(), issuerURL)
	if err != nil {
		return nil, err
	}

	// Build middleware
	return &OIDCVerifier{
		clientID: clientID,
		endpoint: provider.Endpoint(),
		verifier: provider.Verifier(&oidc.Config{SkipClientIDCheck: true}),
	}, nil
}

func (v *OIDCVerifier) ValidateCredentialsWithRoles(ctx context.Context, username, password string, roles []string) (subject string, err error) {
	// To keep things simple, we will just use the Direct/Password Grant flow.
	// We might improve later by using e.g. a Token Grant flow.
	oauth2Config := oauth2.Config{
		ClientID: v.clientID,
		Endpoint: v.endpoint,
		Scopes:   []string{oidc.ScopeOpenID},
	}
	oauth2Token, err := oauth2Config.PasswordCredentialsToken(ctx, username, password)
	if err != nil {
		return "", fmt.Errorf("failed to validate credentials using password grant: %w", err)
	}

	// Parse token
	oidcToken, err := v.verifier.Verify(ctx, oauth2Token.AccessToken)
	if err != nil {
		log.Debug().Err(err).Str("token", oauth2Token.AccessToken).Msg("VerifyToken: Token is invalid")
		return "", fmt.Errorf("failed to verify OIDC token: %w", err)
	}

	// Extract claims
	var claims struct {
		Roles []string `json:"roles"`
	}
	if err := oidcToken.Claims(&claims); err != nil {
		log.Debug().Err(err).Msg("VerifyToken: Failed to extract custom claims")
		return "", fmt.Errorf("failed to extract custom claims from OIDC token: %w", err)
	}

	// Verify roles
	if !verifyRoles(claims.Roles, roles) {
		log.Debug().Strs("expected", roles).Strs("actual", claims.Roles).Msg("VerifyToken: Expected role missing")
		return "", errors.New("expected roles missing")
	}
	return oidcToken.Subject, nil
}
