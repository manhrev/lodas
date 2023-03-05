package result

import (
	"context"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"github.com/manhrev/lodas/backend/lodas/pkg/ent/enttest"
	// "github.com/manhrev/lodas/backend/lodas/pkg/ent/enttest"
)

const (
	_driver = "mysql"
	_url    = "root:password@1@tcp(localhost:33306)/lodas?charset=utf8&parseTime=true"
)

func TestCreate(t *testing.T) {
	var (
		ctx    = context.Background()
		client = enttest.Open(t, _driver, _url)
	)
	cl := New(client)
	// results := crawler.Crawl()

	// // fmt.Println(results[lodas.Province_BINH_DINH])
	// err := result.CreateDaily(ctx, results)
	// sheetObj := client.Sheet.Query().Where(sheet.IDEQ(11)).OnlyX(ctx)
	// resultObj := client.Result.Query().Where(result.ProvinceEQ(int64(lodas.Province_MIEN_BAC)), result.CreatedTimeGTE(time.Unix(1677517200, 0)), result.CreatedTimeLTE(time.Unix(1677517200, 0))).OnlyX(ctx)
	// betSettingObj := client.BetSetting.Query().Where(betsetting.CreatedTimeGTE(time.Unix(1677517200, 0))).OnlyX(ctx)
	// err := checkResult(client, sheetObj, resultObj, betSettingObj)
	err := cl.CheckResultDaily(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
