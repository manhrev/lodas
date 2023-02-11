package lodas

import (
	"context"

	extractor "github.com/manhrev/lodas/backend/auth/pkg/extractor"
	"github.com/manhrev/lodas/backend/lodas/internal/status"
	lodas_pb "github.com/manhrev/lodas/backend/lodas/pkg/api"
)

func (s *lodasServer) CreateSheet(
	ctx context.Context,
	request *lodas_pb.CreateSheetRequest,
) (*lodas_pb.CreateSheetReply, error) {
	userId, err := extractor.New().GetUserID(ctx)
	if err != nil {
		return nil, status.Internal(err.Error())
	}

	_, err = s.repository.Sheet.Create(
		ctx,
		userId,
		request.GetName(),
		request.GetArea(),
		request.GetProvince(),
		request.GetRatio(),
		request.GetResultTime())
	if err != nil {
		return nil, err
	}

	return &lodas_pb.CreateSheetReply{}, nil
}
