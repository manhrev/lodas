package betsetting

import (
	"github.com/manhrev/lodas/backend/lodas/pkg/ent"
	"github.com/manhrev/lodas/backend/lodas/pkg/ent/schema"
)

type BetSetting interface {
}

type betSettingImpl struct {
	entClient *ent.Client
}

func New(entCient *ent.Client) BetSetting {
	entCient.Record.Create().SetWinInfo(&schema.PrizeMap{})

	return &betSettingImpl{
		entClient: entCient,
	}

}
