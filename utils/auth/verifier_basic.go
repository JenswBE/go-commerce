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
		expectedSecret: base64.StdEncoding.EncodeToString([]byte(username + ":" + password)),
	}
}

func (v *BasicVerifier) EnforceRoles(ctx context.Context, roles []string, basicSecret string) (string, error) {
	basicSecret = strings.TrimPrefix(basicSecret, "Basic ")
	if basicSecret != v.expectedSecret {
		log.Debug().Msg("BasicVerifier: Provided basic secret does not match expected secret")
		return "", errors.New("provided basic secret invalid")
	}
	return v.username, nil
}
