package sheet

import (
	"context"
	"log"
	"time"

	"github.com/manhrev/lodas/backend/lodas/internal/status"
	lodas "github.com/manhrev/lodas/backend/lodas/pkg/api"
	lodas_pb "github.com/manhrev/lodas/backend/lodas/pkg/api"
	"github.com/manhrev/lodas/backend/lodas/pkg/code"
	"github.com/manhrev/lodas/backend/lodas/pkg/ent"
	"github.com/manhrev/lodas/backend/lodas/pkg/ent/betsetting"
	"github.com/manhrev/lodas/backend/lodas/pkg/ent/record"
	"github.com/manhrev/lodas/backend/lodas/pkg/ent/schema"
	"github.com/manhrev/lodas/backend/lodas/pkg/ent/sheet"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Sheet interface {
	Create(
		ctx context.Context,
		userId int64,
		name string,
		area lodas_pb.Area,
		province lodas_pb.Province,
		ratio float64,
		win_ratio float64,
		result_time *timestamppb.Timestamp,
	) (*ent.Sheet, error)

	List(
		ctx context.Context,
		userId int64,
		limit uint32,
		offset uint64,
		ascending bool,
		sortBy lodas_pb.SheetSortBy,
		from *timestamppb.Timestamp,
		to *timestamppb.Timestamp,
		ids []int64,
		sheetStatus lodas.SheetStatus,
	) ([]*ent.Sheet, int64, error)

	Update(
		ctx context.Context,
		userId int64,
		sheetId int64,
		name string,
		area lodas_pb.Area,
		province lodas_pb.Province,
		ratio float64,
		win_ratio float64,
		result_time *timestamppb.Timestamp,
		status int64,
	) error

	Delete(
		ctx context.Context,
		userId int64,
		ids []int64,
	) error

	Submit(
		ctx context.Context,
		userId int64,
		id int64,
	) error
}

type sheetImpl struct {
	entClient *ent.Client
}

func New(entCient *ent.Client) Sheet {
	return &sheetImpl{
		entClient: entCient,
	}
}

func (s *sheetImpl) Create(
	ctx context.Context,
	userId int64,
	name string,
	area lodas_pb.Area,
	province lodas_pb.Province,
	ratio float64,
	win_ratio float64,
	result_time *timestamppb.Timestamp,
) (*ent.Sheet, error) {
	newBetSetting := &ent.BetSetting{}
	newBetSettings, err := s.entClient.BetSetting.Query().Where(betsetting.UserIDEQ(userId)).Order(ent.Desc(betsetting.FieldCreatedTime)).All(ctx)
	if err != nil || len(newBetSettings) == 0 {
		newBetSetting, err = s.entClient.BetSetting.Create().SetUserID(userId).SetValues(&schema.BetSettingMap{
			LO2: 70,
			LO3: 600,
			DE2: 70,
			DE3: 600,
			DA:  600,
		}).SetCreatedTime(time.Now().In(time.FixedZone("UTC+7", 7*60*60))).Save(ctx)
		if err != nil {
			return nil, status.Internal(err.Error())
		}
		return nil, status.Internal(err.Error())
	}
	else{
		newBetSetting = newBetSettings[0]
	}

	createdSheet, err := s.entClient.Sheet.Create().
		SetUserID(userId).
		SetName(name).
		SetArea(int64(area)).
		SetProvince(int64(province)).
		SetRatio(float64(ratio)).
		SetWinRatio(float64(win_ratio)).
		SetResultTime(result_time.AsTime()).
		SetBetSetting(newBetSetting).
		SetCreatedTime(time.Now()).
		SetUpdatedTime(time.Now()).
		Save(ctx)

	if err != nil {
		return nil, status.Internal(err.Error())
	}

	return createdSheet, nil
}

func (s *sheetImpl) List(
	ctx context.Context,
	userId int64,
	limit uint32,
	offset uint64,
	ascending bool,
	sortBy lodas_pb.SheetSortBy,
	from *timestamppb.Timestamp,
	to *timestamppb.Timestamp,
	ids []int64,
	sheetStatus lodas.SheetStatus,
) ([]*ent.Sheet, int64, error) {
	query := s.entClient.Sheet.Query().
		Where(sheet.UserIDEQ(userId))

	if len(ids) > 0 {
		query.Where(sheet.IDIn(ids...))
	}
	switch int64(sheetStatus) {
	case int64(lodas.SheetStatus_SHEET_STATUS_NOT_SUBMITTED):
		query.Where(sheet.StatusEQ(int64(lodas.SheetStatus_SHEET_STATUS_NOT_SUBMITTED)))
	case int64(lodas.SheetStatus_SHEET_STATUS_SUBMITTED):
		query.Where(sheet.StatusIn(int64(lodas.SheetStatus_SHEET_STATUS_SUBMITTED), int64(lodas.SheetStatus_SHEET_STATUS_GOT_RESULT)))
	case int64(lodas.SheetStatus_SHEET_STATUS_GOT_RESULT):
		query.Where(sheet.StatusEQ(int64(lodas.SheetStatus_SHEET_STATUS_GOT_RESULT)))
	}
	// ascending?
	if ascending {
		query.Order(ent.Asc(sheet.FieldUpdatedTime))
	} else {
		query.Order(ent.Desc(sheet.FieldUpdatedTime))
	}

	// time range
	if from != nil && to != nil {
		query.Where(
			sheet.CreatedTimeGTE(from.AsTime().Local()),
			sheet.CreatedTimeLTE(to.AsTime().Local()),
		)
	}
	// Count number of records
	total, err := query.Count(ctx)
	if err != nil {
		return nil, 0, status.Error(code.Code_INTERNAL)
	}

	//limit offset
	if limit <= 0 {
		query.Limit(10)
	} else {
		query.Limit(int(limit))
	}

	if offset > 0 {
		query.Offset(int(offset))
	} else {
		query.Offset(0)
	}

	sheets, err := query.All(ctx)
	if err != nil {
		return nil, 0, status.Internal(err.Error())
	}

	return sheets, int64(total), nil
}
func (s *sheetImpl) Update(
	ctx context.Context,
	userId int64,
	sheetId int64,
	name string,
	area lodas_pb.Area,
	province lodas_pb.Province,
	ratio float64,
	win_ratio float64,
	result_time *timestamppb.Timestamp,
	sheet_status int64,
) error {
	numUpdated, err := s.entClient.Sheet.Update().
		SetName(name).
		SetArea(int64(area)).
		SetProvince(int64(province)).
		SetRatio(float64(ratio)).
		SetWinRatio(float64(win_ratio)).
		SetUpdatedTime(time.Now()).
		SetResultTime(result_time.AsTime()).
		SetStatus(sheet_status).
		Where(sheet.IDEQ(sheetId), sheet.UserIDEQ(userId)).
		Save(ctx)

	if err != nil {
		return status.Internal(err.Error())
	}
	if numUpdated == 0 {
		return status.Internal("Can't update sheet info")
	}
	return nil
}

func (s *sheetImpl) Delete(
	ctx context.Context,
	userId int64,
	ids []int64,
) error {
	// Delete all records belong to sheets
	_, err := s.entClient.Record.Delete().Where(record.HasSheetWith(sheet.IDIn(ids...))).Exec(ctx)
	if err != nil {
		log.Printf("Error delete sheet - delete record: %v", err)
		return status.Internal(err.Error())
	}
	// Delete sheets
	_, err = s.entClient.Sheet.
		Delete().
		Where(sheet.IDIn(ids...), sheet.UserIDEQ(userId)).
		Exec(ctx)
	if err != nil {
		log.Printf("Error delete sheet - delete sheet: %v", err)
		return status.Internal(err.Error())
	}
	return nil
}

func (s *sheetImpl) Submit(
	ctx context.Context,
	userId int64,
	id int64,
) error {
	sheetObj, err := s.entClient.Sheet.Query().Where(sheet.IDEQ(id), sheet.UserIDEQ(userId)).Only(ctx)
	if err != nil {
		log.Printf("Error Submit sheet - not found sheet: %v", err)
		return status.Internal(err.Error())
	}
	err = processCashIn(s.entClient, sheetObj)
	if err != nil {
		log.Printf("Error Submit sheet - processCashIn: %v", err)
		return status.Internal(err.Error())
	}

	return nil
}
