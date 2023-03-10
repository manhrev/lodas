package cache

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	v8 "github.com/go-redis/redis/v8"
	"github.com/manhrev/lodas/backend/auth/internal/cache/token"
	u "github.com/manhrev/lodas/backend/auth/internal/cache/user"
	"github.com/manhrev/lodas/backend/auth/pkg/ent"
)

var _defaultTTL = 24 * time.Hour

var (
	cache_host     string = os.Getenv("CACHE_HOST")
	cache_port     string = os.Getenv("CACHE_PORT")
	cache_password string = os.Getenv("CACHE_PASSWORD")
	cache_db       string = os.Getenv("CACHE_DB")
)

type Cache struct {
	accessToken token.Token
	user        u.User
}

func New(ent *ent.Client) (*Cache, error) {
	var (
		r   *v8.Client
		ttl = _defaultTTL
	)
	db, err := strconv.Atoi(cache_db)
	if err != nil {
		panic(err)
	}
	r = v8.NewClient(&v8.Options{
		Addr:     fmt.Sprintf("%s:%s", cache_host, cache_port),
		Password: cache_password, // no password set
		DB:       db,             // use default DB
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
