package record

import (
	"github.com/manhrev/lodas/backend/lodas/pkg/ent"
	"github.com/manhrev/lodas/backend/lodas/pkg/ent/schema"
)

type Record interface {
}

type recordImpl struct {
	entClient *ent.Client
}

func New(entCient *ent.Client) Record {
	entCient.Record.Create().SetWinInfo(&schema.PrizeMap{})

	return &recordImpl{
		entClient: entCient,
	}

}
