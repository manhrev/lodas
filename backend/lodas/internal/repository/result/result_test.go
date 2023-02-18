package result

import (
	"context"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	lodas "github.com/manhrev/lodas/backend/lodas/pkg/api"
	"github.com/manhrev/lodas/backend/lodas/pkg/ent/enttest"
	"github.com/manhrev/lodas/backend/lodas/pkg/ent/schema"
)

const (
	_driver = "mysql"
	_url    = "root:password@1@tcp(localhost:33306)/lodas?charset=utf8&parseTime=true"
)

func TestCreate(t *testing.T) {
	var (
		ctx    = context.Background()
		client = enttest.Open(t, _driver, _url)
	)
	result := New(client)
	rM := schema.PrizeMap{
		lodas.Prize_PRIZE_1DB: {"dfsdf"},
	}
	err := result.Create(ctx, rM)
	if err != nil {
		log.Fatal(err)
	}
}
