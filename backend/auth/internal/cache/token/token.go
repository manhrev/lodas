package token

import (
	"context"

	"github.com/manhrev/runtracking/backend/auth/internal/service/token/signer"
)

type Token interface {
	IsActive(ctx context.Context, token *signer.Token) (bool, error)
	Set(ctx context.Context, token *signer.Token) error
	Revoke(ctx context.Context, jti string) error
}
