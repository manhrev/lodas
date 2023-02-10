package record

import "github.com/manhrev/lodas/backend/lodas/pkg/ent"

type Record interface {
}

type recordImpl struct {
	entClient *ent.Client
}

func New(entCient *ent.Client) Record {
	return &recordImpl{
		entClient: entCient,
	}
}
