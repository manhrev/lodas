package result

import (
	"context"
	"log"

	"github.com/manhrev/lodas/backend/lodas/internal/status"
	lodas "github.com/manhrev/lodas/backend/lodas/pkg/api"
	"github.com/manhrev/lodas/backend/lodas/pkg/ent"
	"github.com/manhrev/lodas/backend/lodas/pkg/ent/schema"
)

type Result interface {
	Create(
		ctx context.Context,
		resultMap schema.PrizeMap,
	) error
}

type resultImpl struct {
	entClient *ent.Client
}

func New(entClient *ent.Client) Result {
	return &resultImpl{
		entClient: entClient,
	}
}

func (r *resultImpl) Create(ctx context.Context, resultMap schema.PrizeMap) error {
	res, err := r.entClient.Result.Create().SetPrizeMap(&resultMap).Save(ctx)
	if err != nil {
		return status.Internal(err.Error())
	}
	log.Println((*res.PrizeMap)[lodas.Prize_PRIZE_1DB])
	return nil
}
