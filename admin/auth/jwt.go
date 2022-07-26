package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(signingKey [64]byte, validity time.Duration) (string, error) {
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(validity)),
	})

	// Sign token
	signedToken, err := token.SignedString(signingKey[:])
	if err != nil {
		return "", fmt.Errorf("failed to sign JWT token: %w", err)
	}
	return signedToken, nil
}

func ValidateJWT(tokenString string, signingKey [64]byte) (jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&jwt.RegisteredClaims{},
		func(token *jwt.Token) (interface{}, error) { return signingKey[:], nil },
		jwt.WithValidMethods([]string{jwt.SigningMethodHS512.Name}),
	)
	if err != nil {
		return jwt.RegisteredClaims{}, fmt.Errorf("failed to parse provided token: %w", err)
	}
	if !token.Valid {
		return jwt.RegisteredClaims{}, errors.New("provided token is not valid")
	}
	return *token.Claims.(*jwt.RegisteredClaims), nil
}
