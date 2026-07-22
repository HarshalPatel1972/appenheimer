package auth

import "context"

type IdentityProvider interface {
	Authenticate(ctx context.Context, email, password string) (string, error)
	ResetPassword(ctx context.Context, email string) error
	ConfirmReset(ctx context.Context, token, newPassword string) error
}
