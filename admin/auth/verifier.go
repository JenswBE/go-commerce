package auth

import "context"

const RoleAdmin = "admin"

type Verifier interface {
	ValidateCredentialsWithRoles(ctx context.Context, username, password string, roles []string) error
	EnforceRoles(ctx context.Context, roles []string, secret string) (string, error)
}
