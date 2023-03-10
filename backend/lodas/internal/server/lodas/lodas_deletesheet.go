package lodas

import (
	"context"

	extractor "github.com/manhrev/lodas/backend/auth/pkg/extractor"
	"github.com/manhrev/lodas/backend/lodas/internal/status"
	lodas_pb "github.com/manhrev/lodas/backend/lodas/pkg/api"
)

func (s *lodasServer) DeleteSheet(
	ctx context.Context,
	request *lodas_pb.DeleteSheetRequest,
) (*lodas_pb.DeleteSheetReply, error) {
	userId, err := extractor.New().GetUserID(ctx)
	if err != nil {
		return nil, status.Internal(err.Error())
	}
	err = s.repository.Sheet.Delete(
		ctx,
		userId,
		request.GetIds())
	if err != nil {
		return nil, status.Internal(err.Error())
	}
	return &lodas_pb.DeleteSheetReply{}, nil
}
