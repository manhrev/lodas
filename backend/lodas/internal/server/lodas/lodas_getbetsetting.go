package lodas

import (
	"context"
	"log"

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

	log.Println(userId)
	return &lodas_pb.GetBetSettingReply{}, nil
}
