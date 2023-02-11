package lodas

import (
	"context"

	extractor "github.com/manhrev/lodas/backend/auth/pkg/extractor"

	"github.com/manhrev/lodas/backend/lodas/internal/status"
	lodas_pb "github.com/manhrev/lodas/backend/lodas/pkg/api"
)

func (s *lodasServer) CreateBetSetting(
	ctx context.Context,
	request *lodas_pb.CreateBetSettingRequest,
) (*lodas_pb.CreateBetSettingReply, error) {
	userId, err := extractor.New().GetUserID(ctx)
	if err != nil {
		return nil, status.Internal(err.Error())
	}
	_, err = s.repository.BetSetting.Create(
		ctx,
		userId,
		request.GetBetSettingInfo(),
	)
	if err != nil {
		return nil, err
	}
	return &lodas_pb.CreateBetSettingReply{}, nil
}
