package lodas

import (
	"context"

	lodas_pb "github.com/manhrev/lodas/backend/lodas/pkg/api"
)

func (s *lodasServer) CreateSheet(
	ctx context.Context,
	request *lodas_pb.CreateSheetRequest,
) (*lodas_pb.CreateSheetReply, error) {

	return &lodas_pb.CreateSheetReply{}, nil
}
