package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(signingKey string, validity time.Duration) (string, error) {
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(validity)),
	})

	// Sign token
	signedToken, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", fmt.Errorf("failed to sign JWT token: %w", err)
	}
	return signedToken, nil
}

func ValidateJWT(tokenString string, signingKey string) error {
	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) { return []byte(signingKey), nil },
		jwt.WithValidMethods([]string{jwt.SigningMethodHS512.Name}),
	)
	if err != nil {
		return fmt.Errorf("failed to parse provided token: %w", err)
	}
	if !token.Valid {
		return errors.New("provided token is not valid")
	}
	return nil
}
