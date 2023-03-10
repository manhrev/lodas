package sheet

import (
	"context"
	"fmt"

	lodas "github.com/manhrev/lodas/backend/lodas/pkg/api"
	"github.com/manhrev/lodas/backend/lodas/pkg/ent"
	"github.com/manhrev/lodas/backend/lodas/pkg/ent/record"
	"github.com/manhrev/lodas/backend/lodas/pkg/ent/sheet"
)

func processCashIn(entClient *ent.Client, sheetObj *ent.Sheet) error {
	// Create transaction
	tx, err := entClient.Tx(context.Background())
	if err != nil {
		return err
	}

	// Get all records of sheet
	records, err := entClient.Record.Query().Where(record.HasSheetWith(sheet.ID(sheetObj.ID))).All(context.Background())
	if err != nil {
		return err
	}
	// compute total cash in for sheet
	totalCashIn := int64(0)
	// compute cash in for each record
	for _, recordObj := range records {
		cashIn := int64(len(recordObj.Numbers)) * int64(len(recordObj.Prize)) * recordObj.CashAmount
		fmt.Println("recordId: ", recordObj.ID, "cashIn: ", cashIn, "cashAmount: ", recordObj.CashAmount, "len(recordObj.Numbers): ", len(recordObj.Numbers), "len(recordObj.Prize): ", len(recordObj.Prize))
		err := entClient.Record.UpdateOneID(recordObj.ID).SetCashIn(cashIn).Exec(context.Background())
		if err != nil {
			return tx.Rollback()
		}
		totalCashIn += cashIn
	}

	// Update sheet status to SUBMITTED
	err = entClient.Sheet.UpdateOneID(sheetObj.ID).SetTotalCashin(totalCashIn).SetStatus(int64(lodas.SheetStatus_SHEET_STATUS_SUBMITTED)).Exec(context.Background())
	if err != nil {
		return tx.Rollback()
	}
	return nil
}
