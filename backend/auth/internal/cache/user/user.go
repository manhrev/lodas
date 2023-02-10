package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	v8 "github.com/go-redis/redis/v8"
	"github.com/manhrev/runtracking/backend/auth/pkg/data"
	"github.com/manhrev/runtracking/backend/auth/pkg/ent"
	"github.com/manhrev/runtracking/backend/auth/pkg/ent/user"
	"google.golang.org/protobuf/proto"
)

type User interface {
	Set(ctx context.Context, user *data.User) error
	Get(ctx context.Context, userID int64) (*data.User, error)
	Delete(ctx context.Context, userID int64) error
}

type redis struct {
	ent    *ent.Client
	client *v8.Client
	ttl    time.Duration
}

func Redis(client *v8.Client, ent *ent.Client, ttl time.Duration) User {
	return &redis{
		ent:    ent,
		client: client,
		ttl:    ttl,
	}
}

func (n redis) Get(ctx context.Context, userID int64) (*data.User, error) {
	bytes, err := n.client.Get(ctx, key(userID)).Bytes()
	if err == nil {
		return unmarshal(bytes)
	}
	if !errors.Is(v8.Nil, err) {
		return nil, errors.New("could not get user on redis")
	}

	u, err := n.ent.User.Query().
		Where(user.ID(userID)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	u0 := &data.User{
		UserId:   u.ID,
		Email:    u.Email,
		UserName: u.Username,
	}

	if err = n.Set(ctx, u0); err != nil {
		return nil, errors.New("could not get user on redis")
	}

	return u0, nil
}

func (n redis) Set(ctx context.Context, u *data.User) error {
	bytes, err := marshal(u)
	if err != nil {
		return err
	}

	return n.client.Set(ctx, key(u.UserId), bytes, n.ttl).Err()
}

func (n redis) Delete(ctx context.Context, userID int64) error {
	return n.client.Del(ctx, key(userID)).Err()
}

func unmarshal(bytes []byte) (*data.User, error) {
	var u data.User
	return &u, proto.Unmarshal(bytes, &u)
}

func marshal(u *data.User) ([]byte, error) {
	return proto.Marshal(u)
}

func key(userID int64) string {
	return fmt.Sprintf("user:%d", userID)
}
