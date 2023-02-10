package lodas

import (
	"context"

	lodas_pb "github.com/manhrev/lodas/backend/lodas/pkg/api"
)

func (s *lodasServer) DeleteSheet(
	ctx context.Context,
	request *lodas_pb.DeleteSheetRequest,
) (*lodas_pb.DeleteSheetReply, error) {

	return &lodas_pb.DeleteSheetReply{}, nil
}
