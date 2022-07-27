package auth

import (
	"context"

	"golang.org/x/exp/slices"
)

const RoleAdmin = "admin"

type Verifier interface {
	ValidateCredentialsWithRoles(ctx context.Context, username, password string, roles []string) (string, error)
}

func verifyRoles(actualRoles, expectedRoles []string) bool {
	for _, expected := range expectedRoles {
		if !slices.Contains(actualRoles, expected) {
			return false
		}
	}
	return true
}
