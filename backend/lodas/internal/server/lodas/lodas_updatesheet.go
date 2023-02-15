package lodas

import (
	"context"

	extractor "github.com/manhrev/lodas/backend/auth/pkg/extractor"
	"github.com/manhrev/lodas/backend/lodas/internal/status"
	lodas_pb "github.com/manhrev/lodas/backend/lodas/pkg/api"
)

func (s *lodasServer) UpdateSheet(
	ctx context.Context,
	request *lodas_pb.UpdateSheetRequest,
) (*lodas_pb.UpdateSheetReply, error) {
	userId, err := extractor.New().GetUserID(ctx)
	if err != nil {
		return nil, status.Internal(err.Error())
	}
	err = s.repository.Sheet.Update(
		ctx,
		userId,
		request.GetSheetId(),
		request.GetName(),
		request.GetArea(),
		request.GetProvince(),
		request.GetRatio(),
		request.GetResultTime())
	if err != nil {
		return nil, status.Internal(err.Error())
	}
	return &lodas_pb.UpdateSheetReply{}, nil
}
