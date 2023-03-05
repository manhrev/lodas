package crawler

import (
	"log"
	"regexp"
	"time"

	lodas "github.com/manhrev/lodas/backend/lodas/pkg/api"

	"github.com/manhrev/lodas/backend/lodas/pkg/ent/schema"
	"github.com/mmcdole/gofeed"
)

var (
	resultMTURLs = map[lodas.Province][]string{
		lodas.Province_PHU_YEN:    {"https://xskt.com.vn/rss-feed/phu-yen-xspy.rss"},
		lodas.Province_HUE:        {"https://xskt.com.vn/rss-feed/thua-thien-hue-xstth.rss"},
		lodas.Province_DAK_LAK:    {"https://xskt.com.vn/rss-feed/dac-lac-xsdlk.rss"},
		lodas.Province_QUANG_NAM:  {"https://xskt.com.vn/rss-feed/quang-nam-xsqnm.rss"},
		lodas.Province_DA_NANG:    {"https://xskt.com.vn/rss-feed/da-nang-xsdng.rss"},
		lodas.Province_KHANH_HOA:  {"https://xskt.com.vn/rss-feed/khanh-hoa-xskh.rss"},
		lodas.Province_BINH_DINH:  {"https://xskt.com.vn/rss-feed/binh-dinh-xsbdi.rss"},
		lodas.Province_QUANG_BINH: {"https://xskt.com.vn/rss-feed/quang-binh-xsqb.rss"},
		lodas.Province_QUANG_TRI:  {"https://xskt.com.vn/rss-feed/quang-tri-xsqt.rss"},
		lodas.Province_GIA_LAI:    {"https://xskt.com.vn/rss-feed/gia-lai-xsgl.rss"},
		lodas.Province_NINH_THUAN: {"https://xskt.com.vn/rss-feed/ninh-thuan-xsnt.rss"},
		lodas.Province_DAK_NONG:   {"https://xskt.com.vn/rss-feed/dac-nong-xsdno.rss"},
		lodas.Province_QUANG_NGAI: {"https://xskt.com.vn/rss-feed/quang-ngai-xsqng.rss"},
		lodas.Province_KON_TUM:    {"https://xskt.com.vn/rss-feed/kon-tum-xskt.rss"},
	}
	resultMBURLs = map[lodas.Province][]string{
		lodas.Province_MIEN_BAC: {"https://xskt.com.vn/rss-feed/mien-bac-xsmb.rss"},
	}
)

type KQXS struct {
	Ptime  *time.Time
	Result *schema.PrizeMap
}

func CrawlResult(resultURL string, area lodas.Area, resultTimeRequest *time.Time) (*time.Time, *schema.PrizeMap) {
	goFeed := gofeed.NewParser()
	feed, err := goFeed.ParseURL(resultURL)
	if err != nil {
		log.Println(err)
		return nil, nil
	}
	if len(feed.Items) == 0 {
		return nil, nil
	}
	if resultTimeRequest == nil {
		des := feed.Items[0].Description
		timeStr := feed.Items[0].Link
		resultTime := GetTime(timeStr)
		result := CreateResult(area, des)
		if resultTime != nil && result != nil && time.Now().In(time.FixedZone("UTC+7", 7*60*60)).Format("2-1-2006") == resultTime.Format("2-1-2006") {
			return resultTime, result
		}
		return nil, nil
	}
	for _, item := range feed.Items {
		des := item.Description
		timeStr := item.Link
		resultTime := GetTime(timeStr)
		result := CreateResult(area, des)

		if resultTime != nil && resultTime.Format("2-1-2006") == resultTimeRequest.Format("2-1-2006") && result != nil {
			return resultTime, result
		}

	}
	return nil, nil

}
func GetTime(timeStr string) *time.Time {

	r := regexp.MustCompile(`\d+-\d+\-\d\d\d\d`)
	rawValueList := r.FindAllString(timeStr, -1)
	if len(rawValueList) == 0 {
		return nil
	}
	rawValue := rawValueList[0]
	rs, err := time.ParseInLocation("2-1-2006", rawValue, time.FixedZone("UTC+7", 7*60*60))
	if err != nil {
		println("hh")
		return nil
	}
	return &rs

}

func CreateResult(area lodas.Area, des string) *schema.PrizeMap {
	pM := make(schema.PrizeMap)

	r := regexp.MustCompile(`\d\d+`)
	resultList := r.FindAllString(des, -1)

	n := 0
	if area == lodas.Area_AREA_BAC {
		n = 27
	} else if area == lodas.Area_AREA_TRUNG {
		n = 18
	}
	if len(resultList) == n {
		if n == 27 {
			pM[lodas.Prize_PRIZE_1DB] = resultList[0]
			pM[lodas.Prize_PRIZE_1G1] = resultList[1]
			pM[lodas.Prize_PRIZE_1G2] = resultList[2]
			pM[lodas.Prize_PRIZE_2G2] = resultList[3]
			pM[lodas.Prize_PRIZE_1G3] = resultList[4]
			pM[lodas.Prize_PRIZE_2G3] = resultList[5]
			pM[lodas.Prize_PRIZE_3G3] = resultList[6]
			pM[lodas.Prize_PRIZE_4G3] = resultList[7]
			pM[lodas.Prize_PRIZE_5G3] = resultList[8]
			pM[lodas.Prize_PRIZE_6G3] = resultList[9]
			pM[lodas.Prize_PRIZE_1G4] = resultList[10]
			pM[lodas.Prize_PRIZE_2G4] = resultList[11]
			pM[lodas.Prize_PRIZE_3G4] = resultList[12]
			pM[lodas.Prize_PRIZE_4G4] = resultList[13]
			pM[lodas.Prize_PRIZE_1G5] = resultList[14]
			pM[lodas.Prize_PRIZE_2G5] = resultList[15]
			pM[lodas.Prize_PRIZE_3G5] = resultList[16]
			pM[lodas.Prize_PRIZE_4G5] = resultList[17]
			pM[lodas.Prize_PRIZE_5G5] = resultList[18]
			pM[lodas.Prize_PRIZE_6G5] = resultList[19]
			pM[lodas.Prize_PRIZE_1G6] = resultList[20]
			pM[lodas.Prize_PRIZE_2G6] = resultList[21]
			pM[lodas.Prize_PRIZE_3G6] = resultList[22]
			pM[lodas.Prize_PRIZE_1G7] = resultList[23]
			pM[lodas.Prize_PRIZE_2G7] = resultList[24]
			pM[lodas.Prize_PRIZE_3G7] = resultList[25]
			pM[lodas.Prize_PRIZE_4G7] = resultList[26]
		}
		if n == 18 {
			resultList[16] = resultList[16][:len(resultList[16])-1]
			pM[lodas.Prize_PRIZE_1DB] = resultList[0]
			pM[lodas.Prize_PRIZE_1G1] = resultList[1]
			pM[lodas.Prize_PRIZE_1G2] = resultList[2]
			pM[lodas.Prize_PRIZE_1G3] = resultList[3]
			pM[lodas.Prize_PRIZE_2G3] = resultList[4]
			pM[lodas.Prize_PRIZE_1G4] = resultList[5]
			pM[lodas.Prize_PRIZE_2G4] = resultList[6]
			pM[lodas.Prize_PRIZE_3G4] = resultList[7]
			pM[lodas.Prize_PRIZE_4G4] = resultList[8]
			pM[lodas.Prize_PRIZE_5G4] = resultList[9]
			pM[lodas.Prize_PRIZE_6G4] = resultList[10]
			pM[lodas.Prize_PRIZE_7G4] = resultList[11]
			pM[lodas.Prize_PRIZE_1G5] = resultList[12]
			pM[lodas.Prize_PRIZE_1G6] = resultList[13]
			pM[lodas.Prize_PRIZE_2G6] = resultList[14]
			pM[lodas.Prize_PRIZE_3G6] = resultList[15]
			pM[lodas.Prize_PRIZE_1G7] = resultList[16]
			pM[lodas.Prize_PRIZE_1G8] = resultList[17]
		}
	}
	if len(pM) != 0 {
		return &pM
	}
	return nil
}

func Crawl() map[lodas.Province]*KQXS {
	results := map[lodas.Province]*KQXS{}
	ti := time.Now().AddDate(0, 0, 0).In(time.FixedZone("UTC+7", 7*60*60))
	for province, urls := range resultMTURLs {
		for _, url := range urls {
			resultTime, result := CrawlResult(url, lodas.Area_AREA_TRUNG, &ti)
			results[province] = &KQXS{resultTime, result}
			if resultTime != nil && result != nil {
				break
			}
		}

	}
	for province, urls := range resultMBURLs {
		for _, url := range urls {
			resultTime, result := CrawlResult(url, lodas.Area_AREA_BAC, &ti)
			results[province] = &KQXS{resultTime, result}
			if resultTime != nil && result != nil {
				break
			}
		}

	}
	// fmt.Println(results)
	return results
}
