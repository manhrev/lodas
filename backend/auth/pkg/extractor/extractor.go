package extractor

import (
	"context"
	"errors"
	"strconv"

	"github.com/manhrev/lodas/backend/auth/pkg/header"
	"google.golang.org/grpc/metadata"
)

type Extractor interface {
	Get(ctx context.Context, name string) []string
	GetTokenID(ctx context.Context) string
	GetUserID(ctx context.Context) (int64, error)
}

type extractor struct {
}

func New() Extractor {
	return &extractor{}
}

func (t *extractor) Get(ctx context.Context, name string) []string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil
	}

	return md.Get(name)
}

func (t *extractor) GetTokenID(ctx context.Context) string {
	values := t.Get(ctx, header.TokenID)
	if len(values) == 0 {
		// Return empty in case UniversalID is undefined
		// Empty is default UniversalID
		return ""
	}

	return values[0]
}

func (t *extractor) GetUserID(ctx context.Context) (int64, error) {
	values := t.Get(ctx, header.UserID)
	if len(values) == 0 {
		return 0, errors.New("metadata does not have x-user-id")
	}

	return strconv.ParseInt(values[0], 10, 64)
}
