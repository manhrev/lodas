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
			winInfo = map[int64]string{
				1:  recordEnt.WinInfo.PRIZE_1DB,
				2:  recordEnt.WinInfo.PRIZE_1G1,
				3:  recordEnt.WinInfo.PRIZE_1G2,
				4:  recordEnt.WinInfo.PRIZE_2G2,
				5:  recordEnt.WinInfo.PRIZE_1G3,
				6:  recordEnt.WinInfo.PRIZE_2G3,
				7:  recordEnt.WinInfo.PRIZE_3G3,
				8:  recordEnt.WinInfo.PRIZE_4G3,
				9:  recordEnt.WinInfo.PRIZE_5G3,
				10: recordEnt.WinInfo.PRIZE_6G3,
				11: recordEnt.WinInfo.PRIZE_1G4,
				12: recordEnt.WinInfo.PRIZE_2G4,
				13: recordEnt.WinInfo.PRIZE_3G4,
				14: recordEnt.WinInfo.PRIZE_4G4,
				15: recordEnt.WinInfo.PRIZE_5G4,
				16: recordEnt.WinInfo.PRIZE_6G4,
				17: recordEnt.WinInfo.PRIZE_7G4,
				18: recordEnt.WinInfo.PRIZE_1G5,
				19: recordEnt.WinInfo.PRIZE_2G5,
				20: recordEnt.WinInfo.PRIZE_3G5,
				21: recordEnt.WinInfo.PRIZE_4G5,
				22: recordEnt.WinInfo.PRIZE_5G5,
				23: recordEnt.WinInfo.PRIZE_6G5,
				24: recordEnt.WinInfo.PRIZE_1G6,
				25: recordEnt.WinInfo.PRIZE_2G6,
				26: recordEnt.WinInfo.PRIZE_3G6,
				27: recordEnt.WinInfo.PRIZE_1G7,
				28: recordEnt.WinInfo.PRIZE_2G7,
				29: recordEnt.WinInfo.PRIZE_3G7,
				30: recordEnt.WinInfo.PRIZE_4G7,
				31: recordEnt.WinInfo.PRIZE_1G8,
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
