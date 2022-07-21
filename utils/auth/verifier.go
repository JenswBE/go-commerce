package auth

import "context"

const RoleAdmin = "admin"

type Verifier interface {
	EnforceRoles(ctx context.Context, roles []string, secret string) (string, error)
}
