package lodas

import (
	"context"

	lodas_pb "github.com/manhrev/lodas/backend/lodas/pkg/api"
	"github.com/manhrev/lodas/backend/lodas/pkg/ent"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *lodasServer) ListRecords(
	ctx context.Context,
	request *lodas_pb.ListRecordsRequest,
) (*lodas_pb.ListRecordsReply, error) {
	// userId, err := extractor.New().GetUserID(ctx)
	// if err != nil {
	// 	return nil, status.Internal(err.Error())
	// }
	recordList, total, err := s.repository.Record.List(
		ctx,
		request.GetSheetId(),
		request.GetLimit(),
		request.GetOffset(),
		request.GetAscending(),
		request.GetSortBy(),
		request.GetFrom(),
		request.GetTo(),
	)
	if err != nil {
		return nil, err
	}
	return &lodas_pb.ListRecordsReply{
		Records: transformEntRecordListToRecordList(recordList),
		Total:   total,
	}, nil
}

func transformEntRecordListToRecordList(entRecordList []*ent.Record) []*lodas_pb.Record {
	recordList := []*lodas_pb.Record{}
	for _, recordEnt := range entRecordList {
		prizes := []lodas_pb.Prize{}
		for _, prize := range recordEnt.Prize {
			prizes = append(prizes, lodas_pb.Prize(prize))
		}
		winInfo := make(map[int64]string)
		if recordEnt.WinInfo != nil {
			for key, value := range *recordEnt.WinInfo {
				// wininfo contain only 1 number -> value[0]
				winInfo[int64(key)] = value[0]
			}
		}

		recordInfo := &lodas_pb.Record{
			Id:         recordEnt.ID,
			SheetId:    recordEnt.Edges.Sheet.ID,
			Numbers:    recordEnt.Numbers,
			Cash:       recordEnt.CashAmount,
			Prizes:     prizes,
			BetType:    lodas_pb.BetType(recordEnt.BetType),
			CashIn:     recordEnt.CashIn,
			WinInfo:    winInfo,
			CashOut:    recordEnt.CashOut,
			CreateTime: timestamppb.New(recordEnt.CreatedTime),
		}
		recordList = append(recordList, recordInfo)
	}
	return recordList
}
