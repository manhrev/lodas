package record

import (
	"context"
	"time"

	"github.com/manhrev/lodas/backend/lodas/internal/status"
	lodas_pb "github.com/manhrev/lodas/backend/lodas/pkg/api"
	"github.com/manhrev/lodas/backend/lodas/pkg/code"
	"github.com/manhrev/lodas/backend/lodas/pkg/ent"
	"github.com/manhrev/lodas/backend/lodas/pkg/ent/record"
	"github.com/manhrev/lodas/backend/lodas/pkg/ent/sheet"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Record interface {
	Create(
		// Note: check validation before calling this function
		ctx context.Context,
		sheetId int64,
		numbers []string,
		cashAmount int64,
		prizes []lodas_pb.Prize,
		betType lodas_pb.BetType,
	) (*ent.Record, error)

	List(
		ctx context.Context,
		sheetId int64,
		limit uint32,
		offset uint64,
		ascending bool,
		sortBy lodas_pb.RecordSortBy,
		from *timestamppb.Timestamp,
		to *timestamppb.Timestamp,
	) ([]*ent.Record, int64, error)

	Delete(
		ctx context.Context,
		userId int64,
		ids []int64,
	) error
}

type recordImpl struct {
	entClient *ent.Client
}

func New(entCient *ent.Client) Record {

	return &recordImpl{
		entClient: entCient,
	}

}

func (r *recordImpl) Create(
	ctx context.Context,
	sheetId int64,
	numbers []string,
	cashAmount int64,
	prizes []lodas_pb.Prize,
	betType lodas_pb.BetType,
) (*ent.Record, error) {
	intList := []int{}
	for _, prize := range prizes {
		intList = append(intList, int(prize))
	}
	create_record, err := r.entClient.Record.Create().
		SetSheetID(sheetId).
		SetNumbers(numbers).
		SetCashAmount(cashAmount).
		SetPrize(intList).
		SetBetType(int64(betType)).
		SetCreatedTime(time.Now().In(time.FixedZone("UTC+7", 7*60*60))).
		Save(ctx)
	if err != nil {
		return nil, status.Internal(err.Error())
	}

	return create_record, nil
}
func (s *recordImpl) List(
	ctx context.Context,
	sheetId int64,
	limit uint32,
	offset uint64,
	ascending bool,
	sortBy lodas_pb.RecordSortBy,
	from *timestamppb.Timestamp,
	to *timestamppb.Timestamp,
) ([]*ent.Record, int64, error) {
	query := s.entClient.Record.Query().Where(record.HasSheetWith(sheet.IDEQ(sheetId)))

	// ascending?
	if ascending {
		query.Order(ent.Asc(record.FieldCreatedTime))
	} else {
		query.Order(ent.Desc(record.FieldCreatedTime))
	}
	// time range
	if from != nil && to != nil {
		query.Where(
			record.CreatedTimeGTE(from.AsTime().Local()),
			record.CreatedTimeLTE(to.AsTime().Local()),
		)
	}
	// Count number of  records
	total, err := query.Count(ctx)

	if err != nil {
		return nil, 0, status.Error(code.Code_INTERNAL)
	}
	if offset > 0 {
		query.Offset(int(offset))
	} else {
		query.Offset(0)
	}

	records, err := query.WithSheet().All(ctx)
	if err != nil {
		return nil, 0, status.Internal(err.Error())
	}

	return records, int64(total), nil

}

func (s *recordImpl) Delete(
	ctx context.Context,
	userId int64,
	ids []int64,
) error {
	numDeleted, err := s.entClient.Record.
		Delete().
		Where(record.HasSheetWith(sheet.UserIDEQ(userId)), record.IDIn(ids...)).Exec(ctx)
	if err != nil {
		return status.Internal(err.Error())
	}
	if numDeleted == 0 {
		return status.Internal("Can't delete record")
	}
	return nil
}
