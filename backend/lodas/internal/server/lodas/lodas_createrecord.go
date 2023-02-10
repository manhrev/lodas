package lodas

import (
	"context"

	//extractor "github.com/manhrev/lodas/backend/auth/pkg/extractor"

	lodas_pb "github.com/manhrev/lodas/backend/lodas/pkg/api"
)

func (s *lodasServer) CreateRecord(
	ctx context.Context,
	request *lodas_pb.CreateRecordRequest,
) (*lodas_pb.CreateRecordReply, error) {
	// userId, err := extractor.New().GetUserID(ctx)
	// if err != nil {
	// 	return nil, status.Internal(err.Error())
	// }

	return &lodas_pb.CreateRecordReply{}, nil
}
