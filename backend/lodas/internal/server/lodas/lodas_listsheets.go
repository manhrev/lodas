package lodas

import (
	"context"

	extractor "github.com/manhrev/lodas/backend/auth/pkg/extractor"
	"github.com/manhrev/lodas/backend/lodas/internal/status"
	lodas_pb "github.com/manhrev/lodas/backend/lodas/pkg/api"
	"github.com/manhrev/lodas/backend/lodas/pkg/ent"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *lodasServer) ListSheets(
	ctx context.Context,
	request *lodas_pb.ListSheetsRequest,
) (*lodas_pb.ListSheetsReply, error) {
	userId, err := extractor.New().GetUserID(ctx)
	if err != nil {
		return nil, status.Internal(err.Error())
	}

	sheetList, total, err := s.repository.Sheet.List(
		ctx,
		userId,
		request.GetLimit(),
		request.GetOffset(),
		request.GetAscending(),
		request.GetSortBy(),
		request.GetFrom(),
		request.GetTo(),
		request.GetIds(),
		request.GetStatus(),
	)
	if err != nil {
		return nil, err
	}

	return &lodas_pb.ListSheetsReply{
		Sheets: transformEntSheetListToSheetList(sheetList),
		Total:  total,
	}, nil
}

func transformEntSheetListToSheetList(entSheetList []*ent.Sheet) []*lodas_pb.SheetInfo {
	sheetList := []*lodas_pb.SheetInfo{}
	for _, sheetEnt := range entSheetList {
		sheetInfo := &lodas_pb.SheetInfo{
			Id:          sheetEnt.ID,
			Status:      lodas_pb.SheetStatus(sheetEnt.Status),
			Name:        sheetEnt.Name,
			Area:        lodas_pb.Area(sheetEnt.Area),
			Province:    lodas_pb.Province(sheetEnt.Province),
			Ratio:       sheetEnt.Ratio,
			ResultTime:  timestamppb.New(sheetEnt.ResultTime),
			CreatedTime: timestamppb.New(sheetEnt.CreatedTime),
			UpdatedTime: timestamppb.New(sheetEnt.UpdatedTime),
		}
		sheetList = append(sheetList, sheetInfo)
	}
	return sheetList
}
