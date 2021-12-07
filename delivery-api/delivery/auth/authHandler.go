package auth

import (
	"context"

	"encore.dev/beta/auth"
	"encore.dev/beta/errs"
)

var secrets struct {
	Token string
}

// AuthHandler check for an API key to verify that user has correct permissions
//encore:authhandler
func AuthHandler(_ context.Context, header string) (auth.UID, error) {
	if header != secrets.Token {
		return "", &errs.Error{
			Code:    errs.Unauthenticated,
			Message: "invalid token",
		}
	}
	return "", nil
}
