package lodas

import (
	"context"

	lodas_pb "github.com/manhrev/lodas/backend/lodas/pkg/api"
)

func (s *lodasServer) ListSheets(
	ctx context.Context,
	request *lodas_pb.ListSheetsRequest,
) (*lodas_pb.ListSheetsReply, error) {

	return &lodas_pb.ListSheetsReply{}, nil
}
