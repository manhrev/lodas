package cache

import (
	"context"
	"fmt"
	"log"
	"time"

	v8 "github.com/go-redis/redis/v8"
	"github.com/manhrev/runtracking/backend/auth/config"
	"github.com/manhrev/runtracking/backend/auth/internal/cache/token"
	u "github.com/manhrev/runtracking/backend/auth/internal/cache/user"
	"github.com/manhrev/runtracking/backend/auth/pkg/ent"
)

var _defaultTTL = 24 * time.Hour

type Cache struct {
	accessToken token.Token
	user        u.User
}

func New(ent *ent.Client, config *config.Configuration) (*Cache, error) {
	var (
		r   *v8.Client
		ttl = _defaultTTL
	)

	conf := config.Cache
	r = v8.NewClient(&v8.Options{
		Addr:     fmt.Sprintf("%s:%s", conf.Host, conf.Port),
		Password: conf.Password, // no password set
		DB:       conf.Db,       // use default DB
	})

	log.Printf("Redis was created : %v", r)

	log.Printf("Key are got from redis : %s ", r.Get(context.Background(), "started"))

	return &Cache{
		accessToken: token.AccessToken(r),
		user:        u.Redis(r, ent, ttl),
	}, nil
}

func (c *Cache) AccessToken() token.Token {
	return c.accessToken
}

func (c *Cache) User() u.User {
	return c.user
}
