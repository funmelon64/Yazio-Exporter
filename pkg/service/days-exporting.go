package service

import (
	"YazioExporter/cmd/utils"
	"YazioExporter/internal/daysstorage"
	"YazioExporter/internal/yazioscraper"
	"YazioExporter/pkg/yzapi"
	"YazioExporter/pkg/yzparse"
	"fmt"
	"log"
	"sort"
	"time"
)

type JsonsForRequest struct {
	Consumed, Goals, Exercises, Water bool
}

type daysExporter struct{}

func NewDaysExporter() *daysExporter {
	return &daysExporter{}
}

func (de *daysExporter) ExportDaysFromYazioToStorage(dateFrom time.Time, dateTo time.Time, storage *daysstorage.Storage,
	whatRequesting JsonsForRequest, yzFactory yzapi.ClientFactory) {
	dateFrom = utils.TruncToDay(dateFrom)
	dateTo = utils.TruncToDay(dateTo)

	daysForScrape := []time.Time{}
	dailyJsonsMap := map[time.Time]string{}

	dailyJsons := getDaysFromDailyRequestsByRange(dateFrom, dateTo, yzFactory)
	for _, dailyJson := range dailyJsons {
		dailyJsonsMap[dailyJson.Date] = dailyJson.DailyJson
		daysForScrape = append(daysForScrape, dailyJson.Date)
	}

	storage.AddJson(dailyJsonsMap, daysstorage.DAILY)

	sort.Slice(daysForScrape, func(i, j int) bool { return daysForScrape[i].After(daysForScrape[j]) })

	daysForScrape = cropDaysListToRange(daysForScrape, dateFrom, dateTo)

	const workersCount = 5
	scraper := yazioscraper.NewYazioJsonsScraper[time.Time](daysForScrape, workersCount, yzFactory)

	if whatRequesting.Consumed {
		scraper.Scrape(func(cl yzapi.Client, t time.Time) (string, error) {
			if storage.IsInStorage(t, daysstorage.CONSUMED) {
				return "", nil
			}
			return cl.GetConsumed(t)
		}, func(m map[time.Time]string) {
			storage.AddJson(m, daysstorage.CONSUMED)
		})
	}

	if whatRequesting.Goals {
		scraper.Scrape(func(cl yzapi.Client, t time.Time) (string, error) {
			if storage.IsInStorage(t, daysstorage.GOALS) {
				return "", nil
			}
			return cl.GetGoals(t)
		}, func(m map[time.Time]string) {
			storage.AddJson(m, daysstorage.GOALS)
		})
	}

	if whatRequesting.Exercises {
		scraper.Scrape(func(cl yzapi.Client, t time.Time) (string, error) {
			if storage.IsInStorage(t, daysstorage.EXERCISES) {
				return "", nil
			}
			return cl.GetExercises(t)
		}, func(m map[time.Time]string) {
			storage.AddJson(m, daysstorage.EXERCISES)
		})
	}

	if whatRequesting.Water {
		scraper.Scrape(func(cl yzapi.Client, t time.Time) (string, error) {
			if storage.IsInStorage(t, daysstorage.WATER) {
				return "", nil
			}
			return cl.GetWater(t)
		}, func(m map[time.Time]string) {
			storage.AddJson(m, daysstorage.WATER)
		})
	}
}

func cropDaysListToRange(days []time.Time, from time.Time, to time.Time) []time.Time {
	var idxFrom, idxTo int
	for i, day := range days {
		if day == from {
			idxTo = i
		}
		if day == to {
			idxFrom = i
		}
	}
	if idxTo == 0 {
		idxTo = len(days) - 1
	}
	return days[idxFrom : idxTo+1]
}

func requestMonthDaily(month time.Time, yazio yzapi.Client) ([]yzparse.DailyDayJson, error) {
	monthDiaryJson, err := yazio.GetMonthDiary(month)
	if err != nil {
		return nil, fmt.Errorf("error request for month %s: %v", utils.FmtAsMonth(month), err)
	}

	daysJsons, err := yzparse.NutrsDailySplitToDays(monthDiaryJson)
	if err != nil {
		return nil, fmt.Errorf("error parsing month %s: %v", utils.FmtAsMonth(month), err)
	}

	if len(daysJsons) < 1 {
		return nil, nil
	}

	return daysJsons, nil
}

func requestMonthsUntilEnd(from time.Time, step int, yazio yzapi.Client) (result []yzparse.DailyDayJson) {
	const maxMonthsWithoutDays = 3

	monthsWithoutDays := 0
	for curMonth := from; monthsWithoutDays < maxMonthsWithoutDays; curMonth = curMonth.AddDate(0, step, 0) {
		if days, err := requestMonthDaily(curMonth, yazio); days == nil {
			if err != nil {
				log.Println(err)
			} else {
				log.Printf("[Month Dailies] Month %s is empty", utils.FmtAsMonth(curMonth))
			}
			monthsWithoutDays++
		} else {
			result = append(result, days...)
			monthsWithoutDays = 0
		}
	}
	return
}

func requestMonthsInRange(from time.Time, to time.Time, yazio yzapi.Client) (result []yzparse.DailyDayJson) {
	from = utils.TruncToMonth(from)
	to = utils.TruncToMonth(to)

	for curMonth := from; curMonth != to.AddDate(0, 1, 0); curMonth = curMonth.AddDate(0, 1, 0) {
		if days, err := requestMonthDaily(curMonth, yazio); days == nil {
			if err != nil {
				log.Println(err)
			} else {
				log.Printf("[Month Dailies] Month %s is empty", utils.FmtAsMonth(curMonth))
			}
		} else {
			result = append(result, days...)
		}
	}
	return
}

func getDaysFromDailyRequestsByRange(dateFrom time.Time, dateTo time.Time, yzFactory yzapi.ClientFactory) []yzparse.DailyDayJson {
	var dailyJsons []yzparse.DailyDayJson
	yazio := yzFactory.NewClient()

	if dateFrom.IsZero() && dateTo.IsZero() {
		log.Println("[Month Dailies] Requesting months dailies for all months..")
		dailyJsons = requestMonthsUntilEnd(time.Now(), 1, yazio)
		dailyJsons = append(dailyJsons, requestMonthsUntilEnd(time.Now().AddDate(0, -1, 0), -1, yazio)...)
	} else if dateFrom.IsZero() {
		log.Printf("[Month Dailies] Requesting months dailies until %s..", utils.FmtAsMonth(dateTo))
		dailyJsons = requestMonthsUntilEnd(dateTo, -1, yazio)
	} else if dateTo.IsZero() {
		log.Printf("[Month Dailies] Requesting months dailies from %s until last record in diary..", utils.FmtAsMonth(dateFrom))
		dailyJsons = requestMonthsUntilEnd(dateFrom, 1, yazio)
	} else {
		log.Printf("[Month Dailies] Parsing days from %s until %s\n", utils.FmtAsMonth(dateFrom), utils.FmtAsMonth(dateTo))
		dailyJsons = requestMonthsInRange(dateFrom, dateTo, yazio)
	}

	return dailyJsons
}
