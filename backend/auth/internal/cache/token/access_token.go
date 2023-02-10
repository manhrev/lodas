package token

import (
	"context"
	"fmt"
	"time"

	v8 "github.com/go-redis/redis/v8"

	"github.com/manhrev/runtracking/backend/auth/internal/service/token/signer"
)

type accessToken struct {
	client *v8.Client
}

func AccessToken(client *v8.Client) Token {
	return &accessToken{
		client: client,
	}
}

func (n *accessToken) IsActive(ctx context.Context, token *signer.Token) (bool, error) {
	key := n.key(token.Id)
	exists, err := n.client.Exists(ctx, key).Uint64()
	if err != nil {
		return false, err
	}

	return exists > 0, nil
}

func (n *accessToken) Set(ctx context.Context, token *signer.Token) error {

	key := n.key(token.Id)
	ttl := time.Unix(token.ExpiresAt, 0).Sub(time.Now())
	return n.client.Set(ctx, key, 1, ttl).Err()
}

func (n *accessToken) Revoke(ctx context.Context, jti string) error {
	key := n.key(jti)
	return n.client.Del(ctx, key).Err()
}

// func (n *accessToken) RevokeBySafeID(ctx context.Context, safeID string) error {
// 	ctx = xcontext.Detach(ctx)
// 	keys, err := n.client.Keys(ctx, n.key(safeID, "*")).Result()
// 	if err != nil {
// 		return err
// 	}
// 	if len(keys) == 0 {
// 		return nil
// 	}

// 	if err := n.client.Del(ctx, keys...).Err(); err != nil {
// 		logging.Logger(ctx).Error("could not revoke", zap.Any("keys", keys), zap.Error(err))
// 	}

// 	return nil
// }

// func (n *accessToken) RevokeAll(ctx context.Context) error {
// 	ctx = xcontext.Detach(ctx)
// 	keys, err := n.client.Keys(ctx, n.key("*", "*")).Result()
// 	if err != nil {
// 		return err
// 	}
// 	if len(keys) == 0 {
// 		return nil
// 	}

// 	if err := n.client.Del(ctx, keys...).Err(); err != nil {
// 		logging.Logger(ctx).Error("could not revoke", zap.Any("keys", keys), zap.Error(err))
// 	}

// 	return nil
// }

func (n *accessToken) key(jti string) string {
	return fmt.Sprintf("access.token:%s", jti)
}
