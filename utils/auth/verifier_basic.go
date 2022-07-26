package auth

import (
	"context"
	"encoding/base64"
	"errors"
	"strings"

	"github.com/rs/zerolog/log"
)

var _ Verifier = &BasicVerifier{}

type BasicVerifier struct {
	username       string
	expectedSecret string
}

func NewBasicVerifier(username, password string) *BasicVerifier {
	return &BasicVerifier{
		username:       username,
		expectedSecret: encodeBasicLogin(username, password),
	}
}

func (v *BasicVerifier) ValidateCredentialsWithRoles(ctx context.Context, username, password string, roles []string) error {
	if v.expectedSecret != encodeBasicLogin(username, password) {
		return errors.New("invalid credentials")
	}
	// Basic auth doesn't support roles and is always considered and admin.
	// Should only be used for testing!
	return nil
}

func (v *BasicVerifier) EnforceRoles(ctx context.Context, roles []string, basicSecret string) (string, error) {
	basicSecret = strings.TrimPrefix(basicSecret, "Basic ")
	if basicSecret != v.expectedSecret {
		log.Debug().Msg("BasicVerifier: Provided basic secret does not match expected secret")
		return "", errors.New("provided basic secret invalid")
	}
	return v.username, nil
}

func encodeBasicLogin(username, password string) string {
	return base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
}
