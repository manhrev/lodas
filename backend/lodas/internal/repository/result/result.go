package result

import (
	"context"
	"log"

	"github.com/manhrev/lodas/backend/lodas/internal/service/crawler"
	"github.com/manhrev/lodas/backend/lodas/internal/status"
	lodas "github.com/manhrev/lodas/backend/lodas/pkg/api"
	"github.com/manhrev/lodas/backend/lodas/pkg/ent"
	"github.com/manhrev/lodas/backend/lodas/pkg/ent/betsetting"
	"github.com/manhrev/lodas/backend/lodas/pkg/ent/result"
	"github.com/manhrev/lodas/backend/lodas/pkg/ent/sheet"
)

type Result interface {
	//Get result from nhà cái
	CreateDaily(
		ctx context.Context,
		results map[lodas.Province]*crawler.KQXS,
	) error
	//Check result daily
	CheckResultDaily(
		ctx context.Context,
	) error
}

type resultImpl struct {
	entClient *ent.Client
}

func New(entClient *ent.Client) Result {
	return &resultImpl{
		entClient: entClient,
	}
}

func (r *resultImpl) CreateDaily(
	ctx context.Context,
	results map[lodas.Province]*crawler.KQXS,
) error {
	bulk := []*ent.ResultCreate{}
	numDeleted := 0
	for province, kqxs := range results {
		resultTime := kqxs.Ptime
		re := kqxs.Result
		if resultTime != nil && re != nil {
			num, _ := r.entClient.Result.Delete().Where(result.ProvinceEQ(int64(province)), result.CreatedTimeEQ(*resultTime)).Exec(ctx)
			numDeleted += num
			bulk = append(bulk, r.entClient.Result.Create().SetCreatedTime(*resultTime).SetPrizeMap(re).SetProvince(int64(province)))
		}
	}
	created, err := r.entClient.Result.CreateBulk(bulk...).Save(ctx)
	log.Print("Num of deleted:", numDeleted)
	log.Print("Num of created:", len(created))

	if err != nil {
		return status.Internal(err.Error())
	}
	// ke := (*created[0].PrizeMap)
	// fmt.Println(ke)
	return nil
}
func (r *resultImpl) CheckResultDaily(
	ctx context.Context,
) error {
	sheets, err := r.entClient.Sheet.Query().Where(sheet.StatusEQ(int64(lodas.SheetStatus_SHEET_STATUS_SUBMITTED))).All(ctx)
	if err != nil {
		log.Printf("Error CheckResultDaily get Sheet: %v", err)
		return status.Internal(err.Error())
	}
	for _, sheetObj := range sheets {
		betSettingObj, err := r.entClient.BetSetting.Query().Where(
			betsetting.HasSheetsWith(sheet.IDEQ(sheetObj.ID)),
			betsetting.UserIDEQ(sheetObj.UserID),
		).Only(ctx)
		if err != nil {
			log.Printf("Error CheckResultDaily get BetSetting: %v", err)
			// return status.Internal(err.Error())
		}
		startTime, endTime := getTimeRange(&sheetObj.ResultTime, &sheetObj.ResultTime)
		// get resultObj which has the same province and result time with sheetObj

		resultObj, err := r.entClient.Result.Query().Where(result.ProvinceEQ(sheetObj.Province), result.CreatedTimeGTE(startTime), result.CreatedTimeLTE(endTime), result.ProvinceEQ(sheetObj.Province)).Only(ctx)
		if err != nil {
			log.Printf("Error CheckResultDaily get result: %v", err)
			// return status.Internal(err.Error())
		}
		err = checkResult(r.entClient, sheetObj, resultObj, betSettingObj)
		if err != nil {
			log.Printf("Error CheckResultDaily check result: %v", err)
			// return status.Internal(err.Error())
		}
	}

	return nil
}
