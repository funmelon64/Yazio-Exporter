package yzparse

import (
	"encoding/json"
	"log"
	"regexp"
	"time"
)

type DailyDayJson struct {
	Date      time.Time
	DailyJson string
}

var dateRe = regexp.MustCompile(`(?m)"date" *: *"(.*?)"`)

func NutrsDailySplitToDays(nutrsDailyJson string) ([]DailyDayJson, error) {
	diaryDays := []DailyDayJson{}
	jsons := []json.RawMessage{}
	if err := json.Unmarshal([]byte(nutrsDailyJson), &jsons); err != nil {
		return nil, err
	}

	for i := range jsons {
		matchs := dateRe.FindStringSubmatch(string(jsons[i]))
		if len(matchs) >= 2 {
			date, err := time.Parse(time.DateOnly, matchs[1])
			if err != nil {
				log.Printf("[Parser] fail to parse date %s from nutrients-daily JSON, skip: %v", matchs[1], err)
				continue
			}
			diaryDays = append(diaryDays, DailyDayJson{Date: date, DailyJson: string(jsons[i])})
		}
	}

	return diaryDays, nil
}
