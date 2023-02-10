package lodas

import (
	"context"

	lodas_pb "github.com/manhrev/lodas/backend/lodas/pkg/api"
)

func (s *lodasServer) ListRecords(
	ctx context.Context,
	request *lodas_pb.ListRecordsRequest,
) (*lodas_pb.ListRecordsReply, error) {

	return &lodas_pb.ListRecordsReply{}, nil
}
