package sheet

import "github.com/manhrev/lodas/backend/lodas/pkg/ent"

type Sheet interface {
}

type sheetImpl struct {
	entClient *ent.Client
}

func New(entCient *ent.Client) Sheet {
	return &sheetImpl{
		entClient: entCient,
	}
}
