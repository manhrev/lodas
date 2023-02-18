package result

import (
	"context"

	"github.com/manhrev/lodas/backend/lodas/pkg/ent"
	"github.com/manhrev/lodas/backend/lodas/pkg/ent/schema"
)

type Result interface {
	Create(
		ctx context.Context,
		resultMap schema.PrizeMap,
	)
}

type resultImpl struct {
	entClient *ent.Client
}

func New(entClient *ent.Client) Result {
	return &resultImpl{
		entClient: entClient,
	}
}

func (r *resultImpl) Create(ctx context.Context, resultMap schema.PrizeMap) {
	r.entClient.Result.Create().SetPrizeMap(&resultMap)
}
