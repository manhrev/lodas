package betsetting

import (
	"context"
	"time"

	"github.com/manhrev/lodas/backend/lodas/internal/status"
	lodas_pb "github.com/manhrev/lodas/backend/lodas/pkg/api"
	"github.com/manhrev/lodas/backend/lodas/pkg/ent"
	"github.com/manhrev/lodas/backend/lodas/pkg/ent/betsetting"
	"github.com/manhrev/lodas/backend/lodas/pkg/ent/schema"
)

type BetSetting interface {
	Create(
		ctx context.Context,
		userId int64,
		betSetting *lodas_pb.BetSettingInfo,
	) (*ent.BetSetting, error)
	Get(
		ctx context.Context,
		userId int64,
	) (*ent.BetSetting, error)
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
func (b *betSettingImpl) Create(
	ctx context.Context,
	userId int64,
	betSetting *lodas_pb.BetSettingInfo,
) (*ent.BetSetting, error) {
	betSettingEnt, err := b.entClient.BetSetting.Create().SetUserID(userId).SetValues(&schema.BetSettingMap{
		LO2: betSetting.LO2,
		LO3: betSetting.LO3,
		DE2: betSetting.DE2,
		DE3: betSetting.DE3,
		DA:  betSetting.DA,
	}).SetCreatedTime(time.Now()).Save(ctx)
	if err != nil {
		return nil, status.Internal(err.Error())
	}
	return betSettingEnt, nil
}
func (b *betSettingImpl) Get(
	ctx context.Context,
	userId int64,
) (*ent.BetSetting, error) {
	betSettingEnt, err := b.entClient.BetSetting.Query().Where(betsetting.UserIDEQ(userId)).Order(ent.Desc(betsetting.FieldCreatedTime)).All(ctx)
	if err != nil {
		return nil, status.Internal(err.Error())
	}
	if len(betSettingEnt) == 0 {
		return nil, status.Internal("no bet setting found")
	}

	return betSettingEnt[0], nil

}
