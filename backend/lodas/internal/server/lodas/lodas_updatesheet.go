package lodas

import (
	"context"

	lodas_pb "github.com/manhrev/lodas/backend/lodas/pkg/api"
)

func (s *lodasServer) UpdateSheet(
	ctx context.Context,
	request *lodas_pb.UpdateSheetRequest,
) (*lodas_pb.UpdateSheetReply, error) {

	return &lodas_pb.UpdateSheetReply{}, nil
}
