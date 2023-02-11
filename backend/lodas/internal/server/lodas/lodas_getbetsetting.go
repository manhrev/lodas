package lodas

import (
	"context"

	extractor "github.com/manhrev/lodas/backend/auth/pkg/extractor"

	"github.com/manhrev/lodas/backend/lodas/internal/status"
	lodas_pb "github.com/manhrev/lodas/backend/lodas/pkg/api"
)

func (s *lodasServer) GetBetSetting(
	ctx context.Context,
	request *lodas_pb.GetBetSettingRequest,
) (*lodas_pb.GetBetSettingReply, error) {
	userId, err := extractor.New().GetUserID(ctx)
	if err != nil {
		return nil, status.Internal(err.Error())
	}

	betSettingEnt, err := s.repository.BetSetting.Get(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &lodas_pb.GetBetSettingReply{
		BetSettingInfo: &lodas_pb.BetSettingInfo{
			LO2: betSettingEnt.Values.LO2,
			LO3: betSettingEnt.Values.LO3,
			DE2: betSettingEnt.Values.DE2,
			DE3: betSettingEnt.Values.DE3,
			DA:  betSettingEnt.Values.DA,
		},
	}, nil
}
