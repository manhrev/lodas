package lodas

import (
	"context"

	lodas_pb "github.com/manhrev/lodas/backend/lodas/pkg/api"
)

func (s *lodasServer) CreateRecord(
	ctx context.Context,
	request *lodas_pb.CreateRecordRequest,
) (*lodas_pb.CreateRecordReply, error) {
	//userId, err := extractor.New().GetUserID(ctx)

	// if err != nil {
	// 	return nil, status.Internal(err.Error())
	// }
	_, err := s.repository.Record.Create(
		ctx,
		request.GetSheetId(),
		request.GetNumbers(),
		request.GetCash(),
		request.GetPrizes(),
		request.GetBetType(),
	)
	if err != nil {
		return nil, err
	}
	return &lodas_pb.CreateRecordReply{}, nil
}
