package lodas

import (
	"context"

	extractor "github.com/manhrev/lodas/backend/auth/pkg/extractor"
	"github.com/manhrev/lodas/backend/lodas/internal/status"
	lodas_pb "github.com/manhrev/lodas/backend/lodas/pkg/api"
)

func (s *lodasServer) SubmitSheet(
	ctx context.Context,
	request *lodas_pb.SubmitSheetRequest,
) (*lodas_pb.SubmitSheetReply, error) {
	userId, err := extractor.New().GetUserID(ctx)
	if err != nil {
		return nil, status.Internal(err.Error())
	}
	err = s.repository.Sheet.Submit(
		ctx,
		userId,
		request.GetId())
	if err != nil {
		return nil, status.Internal(err.Error())
	}
	return &lodas_pb.SubmitSheetReply{}, nil
}
