package result

import (
	"context"
	"fmt"
	"time"

	lodas "github.com/manhrev/lodas/backend/lodas/pkg/api"
	"github.com/manhrev/lodas/backend/lodas/pkg/ent"
	"github.com/manhrev/lodas/backend/lodas/pkg/ent/record"
	"github.com/manhrev/lodas/backend/lodas/pkg/ent/schema"
	"github.com/manhrev/lodas/backend/lodas/pkg/ent/sheet"
)

func getTimeRange(start *time.Time, end *time.Time) (time.Time, time.Time) {
	startTime := time.Time{}
	endTime := time.Time{}
	if start != nil {
		year, month, date := start.In(time.FixedZone("UTC+7", 7*60*60)).Date()
		startTime = time.Date(year, month, date, 0, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60))
	}
	if end != nil {
		year, month, date := end.In(time.FixedZone("UTC+7", 7*60*60)).Date()
		endTime = time.Date(year, month, date, 23, 59, 59, 0, time.FixedZone("UTC+7", 7*60*60))
	}
	return startTime, endTime
}

func checkResult(entClient *ent.Client, sheetObj *ent.Sheet, resultObj *ent.Result, betSettingObj *ent.BetSetting) error {
	if sheetObj != nil && resultObj != nil && betSettingObj != nil {
		doubleDigitResult := schema.PrizeMap{}
		tripleDigitResult := schema.PrizeMap{}
		for prize, value := range *resultObj.PrizeMap {
			if len(value) >= 2 {
				doubleDigitResult[prize] = value[len(value)-2:]
			}
			if len(value) >= 3 {
				tripleDigitResult[prize] = value[len(value)-3:]
			}
		}
		// check result if sheet is submitted
		if sheetObj.Status == int64(lodas.SheetStatus_SHEET_STATUS_SUBMITTED) {
			// get all records of this sheet
			records, err := entClient.Record.Query().Where(record.HasSheetWith(sheet.IDEQ(sheetObj.ID))).All(context.Background())
			if err != nil {
				return err
			}
			// check result
			for _, recordObj := range records {
				// get numbers, prizers
				winInfo := schema.PrizeMap{}
				numbers := recordObj.Numbers
				prizers := recordObj.Prize

				// check result Lo 2
				if recordObj.BetType == int64(lodas.BetType_LO2) {
					fmt.Println("L02")

					digitResult := doubleDigitResult
					betSetting := betSettingObj.Values.LO2
					count := int64(0)
					for _, prize := range prizers {
						for _, number := range numbers {
							if number == digitResult[lodas.Prize(prize)] {
								winInfo[lodas.Prize(prize)] = number
								count++
							}

						}
					}
					cashOut := count * betSetting * recordObj.CashAmount
					err := entClient.Record.UpdateOneID(recordObj.ID).SetWinInfo(&winInfo).SetCashOut(cashOut).Exec(context.Background())
					if err != nil {
						return err
					}
					fmt.Println("winInfo: ", winInfo)
					fmt.Println("cashOut: ", cashOut)
				}
				// check result Lo 3
				if recordObj.BetType == int64(lodas.BetType_LO3) {
					fmt.Println("L03")

					digitResult := tripleDigitResult
					betSetting := betSettingObj.Values.LO3
					count := int64(0)
					for _, prize := range prizers {
						for _, number := range numbers {
							if number == digitResult[lodas.Prize(prize)] {
								winInfo[lodas.Prize(prize)] = number
								count++
							}

						}
					}
					cashOut := count * betSetting * recordObj.CashAmount
					entClient.Record.UpdateOneID(recordObj.ID).SetWinInfo(&winInfo).SetCashOut(cashOut).Exec(context.Background())
					fmt.Println("winInfo: ", winInfo)
					fmt.Println("cashOut: ", cashOut)
				}
				// check result De 2
				// Chỉ xét giải đặc biệt một lần duy nhất
				if recordObj.BetType == int64(lodas.BetType_DE2) {
					fmt.Println("DE2")

					digitResult := doubleDigitResult
					betSetting := betSettingObj.Values.DE2
					count := int64(0)
					for _, number := range numbers {
						if number == digitResult[lodas.Prize_PRIZE_1DB] {
							winInfo[lodas.Prize_PRIZE_1DB] = number
							count = 1
						}

					}
					cashOut := count * betSetting
					err := entClient.Record.UpdateOneID(recordObj.ID).SetWinInfo(&winInfo).SetCashOut(cashOut).Exec(context.Background())
					if err != nil {
						return err
					}
					fmt.Println("winInfo: ", winInfo)
					fmt.Println("cashOut: ", cashOut)
				}
				// check result De 3
				// Chỉ xét giải đặc biệt một lần duy nhất
				if recordObj.BetType == int64(lodas.BetType_DE3) {
					fmt.Println("DE3")

					digitResult := tripleDigitResult
					betSetting := betSettingObj.Values.DE3
					count := int64(0)
					for _, number := range numbers {
						if number == digitResult[lodas.Prize_PRIZE_1DB] {
							winInfo[lodas.Prize_PRIZE_1DB] = number
							count = 1
						}

					}
					cashOut := count * betSetting * recordObj.CashAmount
					err := entClient.Record.UpdateOneID(recordObj.ID).SetWinInfo(&winInfo).SetCashOut(cashOut).Exec(context.Background())
					if err != nil {
						return err
					}
					fmt.Println("winInfo: ", winInfo)
					fmt.Println("cashOut: ", cashOut)
				}
				// check result Da
				if recordObj.BetType == int64(lodas.BetType_DA) {
					fmt.Println("DA2")

					digitResult := doubleDigitResult
					betSetting := betSettingObj.Values.DA
					winMap := map[string]int64{}
					count := int64(0)
					for _, prize := range prizers {
						for _, number := range numbers {
							if number == digitResult[lodas.Prize(prize)] {
								winInfo[lodas.Prize(prize)] = number

								// If number is in winMap increase count 1 else set count = 1
								if _, ok := winMap[number]; ok {
									winMap[number] = winMap[number] + 1
								} else {
									winMap[number] = 1
								}
							}
						}
					}
					// get array of win count of each number
					valueArray := []int64{}

					for _, value := range winMap {
						valueArray = append(valueArray, value)
					}
					if len(valueArray) > 1 {
						// if A number have duplicate less than B number, A number will be count
						// {56:2, 57:3, 58:3} => count = 2+2+3 = 7
						for i := 0; i < len(valueArray)-1; i++ {
							for j := i + 1; j < len(valueArray); j++ {
								if valueArray[i] <= valueArray[j] {
									count = count + valueArray[i]
								} else {
									count = count + valueArray[j]
								}
							}
						}
					}
					cashOut := count * betSetting * recordObj.CashAmount
					err := entClient.Record.UpdateOneID(recordObj.ID).SetWinInfo(&winInfo).SetCashOut(cashOut).Exec(context.Background())
					if err != nil {
						return err
					}
					fmt.Println("winInfo: ", winInfo)
					fmt.Println("cashOut: ", cashOut)
				}
			}
			// Update sheet status
			err = entClient.Sheet.UpdateOneID(sheetObj.ID).SetStatus(int64(lodas.SheetStatus_SHEET_STATUS_GOT_RESULT)).Exec(context.Background())
			if err != nil {
				return err
			}

		}

	}
	return nil
}
