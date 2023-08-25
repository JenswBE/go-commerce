package auth

import (
	"context"
	"errors"
	"fmt"
	"slices"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"golang.org/x/oauth2"
)

const RoleAdmin = "admin"

type OIDCClient struct {
	verifier     *oidc.IDTokenVerifier
	oauth2Config oauth2.Config
}

func NewOIDCClient(issuerURL, clientID, clientSecret string) (*OIDCClient, error) {
	// Get provider
	provider, err := oidc.NewProvider(context.Background(), issuerURL)
	if err != nil {
		return nil, err
	}

	// Configure an OpenID Connect aware OAuth2 client.
	oauth2Config := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		// "openid" is a required scope for OpenID Connect flows.
		Scopes: []string{oidc.ScopeOpenID},
	}

	// Return client
	return &OIDCClient{
		verifier:     provider.Verifier(&oidc.Config{ClientID: clientID}),
		oauth2Config: oauth2Config,
	}, nil
}

// GetAuthCodeURL generates a random state and returns an auth code URL and the generated state
func (c *OIDCClient) GetAuthCodeURL(redirectURL string) (url, state string) {
	state = uuid.NewString()
	c.oauth2Config.RedirectURL = redirectURL
	return c.oauth2Config.AuthCodeURL(state), state
}

func (c *OIDCClient) ValidateCallbackWithRoles(ctx context.Context, expectedState, queryState, queryCode string, roles []string) (subject string, err error) {
	// Validate state
	if expectedState != queryState {
		return "", fmt.Errorf("provided state %s did not match expected state %s", queryState, expectedState)
	}

	// Exchange token
	oauth2Token, err := c.oauth2Config.Exchange(ctx, queryCode)
	if err != nil {
		return "", fmt.Errorf("failed to exchange token: %w", err)
	}

	// Extract the ID Token from OAuth2 token.
	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		log.Debug().Err(err).Msg("ValidateCallbackWithRoles: OAuth2 token did not include an id_token")
		return "", errors.New("oauth2 token has a missing id_token field")
	}

	// Parse token
	oidcToken, err := c.verifier.Verify(ctx, rawIDToken)
	if err != nil {
		log.Debug().Err(err).Str("token", rawIDToken).Msg("ValidateCallbackWithRoles: Token is invalid")
		return "", fmt.Errorf("failed to verify OIDC token: %w", err)
	}

	// Extract claims
	var claims struct {
		Roles []string `json:"roles"`
	}
	if err := oidcToken.Claims(&claims); err != nil {
		log.Debug().Err(err).Msg("ValidateCallbackWithRoles: Failed to extract custom claims")
		return "", fmt.Errorf("failed to extract custom claims from OIDC token: %w", err)
	}

	// Verify roles
	if !verifyRoles(claims.Roles, roles) {
		log.Debug().Strs("expected", roles).Strs("actual", claims.Roles).Msg("ValidateCallbackWithRoles: Expected role missing")
		return "", errors.New("expected roles missing")
	}
	return oidcToken.Subject, nil
}

func verifyRoles(actualRoles, expectedRoles []string) bool {
	for _, expected := range expectedRoles {
		if !slices.Contains(actualRoles, expected) {
			return false
		}
	}
	return true
}
