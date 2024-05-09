package service

import (
	"YazioExporter/cmd/utils"
	"YazioExporter/internal/daysstorage"
	"YazioExporter/test/mockyzapi"
	"log"
	"regexp"
	"strings"
	"testing"
	"time"
)

type callbackWriter func([]byte) (int, error)

func (c callbackWriter) Write(p []byte) (int, error) {
	return c(p)
}

func TestDaysExport(t *testing.T) {
	e := NewDaysExporter()

	storage := daysstorage.NewEmpty()

	mockSets := mockyzapi.NewMockSettings()
	mockSets.DailyDayFrom = utils.Date(2024, 4, 8)
	mockSets.DailyDayTo = utils.Date(2024, 5, 4)

	daysDif := int(mockSets.DailyDayTo.AddDate(0, 0, 1).Sub(mockSets.DailyDayFrom).Hours() / 24)

	// export goals for all available days
	e.ExportDaysFromYazioToStorage(time.Time{}, time.Time{}, storage,
		JsonsForRequest{false, true, false, false},
		mockyzapi.NewMockClientFactory(mockSets))

	if len(storage.Days) != daysDif {
		t.Errorf("Expected %v Days, got %v", daysDif, len(storage.Days))
		return
	}

	day, ok := storage.Days[utils.Date(2024, 4, 19)]
	if !ok {
		t.Error("Expected day 2024-4-19 to exist")
		return
	}
	if len(day.GoalsJson) <= 4 {
		t.Error("Expected GoalsJson for 2024-4-19 be not empty")
		return
	}
}

func TestDaysExportFrom(t *testing.T) {
	e := NewDaysExporter()

	storage := daysstorage.NewEmpty()
	mockSets := mockyzapi.NewMockSettings()
	mockSets.DailyDayFrom = utils.Month(2024, 4)
	mockSets.DailyDayTo = utils.Date(2024, 4, 30)

	var dateRe = regexp.MustCompile(`(?m)\d{4}-\d\d-\d\d`)

	writer := callbackWriter(func(data []byte) (int, error) {
		dateStr := dateRe.FindString(string(data))
		if dateStr != "" {
			date, _ := time.Parse(time.DateOnly, dateStr)
			if date.Before(utils.Month(2024, 04)) {
				t.Error("invalid dates in call:")
			}
		}
		log.Print(string(data))
		return len(data), nil
	})

	mockSets.Logger = log.New(writer, "", 0)

	e.ExportDaysFromYazioToStorage(utils.Date(2024, 04, 18), time.Time{}, storage,
		JsonsForRequest{false, true, false, false},
		mockyzapi.NewMockClientFactory(mockSets))
}

func TestExportDaysAppendingJsons(t *testing.T) {
	e := NewDaysExporter()

	mockSets := mockyzapi.NewMockSettings()

	storage := daysstorage.NewEmpty()
	// fill storage with days not containing water jsons (2024-04-01 - 2024-04-30)
	e.ExportDaysFromYazioToStorage(utils.Month(2024, 04), utils.Date(2024, 04, 30), storage,
		JsonsForRequest{false, true, false, false},
		mockyzapi.NewMockClientFactory(mockSets))

	// export waters for days in (2024-04-27 - 2024-05-02)
	e.ExportDaysFromYazioToStorage(utils.Date(2024, 04, 23), utils.Date(2024, 05, 03), storage,
		JsonsForRequest{false, false, false, true},
		mockyzapi.NewMockClientFactory(mockSets))

	day, ok := storage.Days[utils.Date(2024, 4, 12)]
	if !ok {
		t.Error("Expected day 2024-4-12 to exist")
		return
	}
	if len(day.WaterJson) > 4 {
		t.Error("Expected WaterJson for 2024-4-12 to be empty")
		return
	}

	day, ok = storage.Days[utils.Date(2024, 4, 29)]
	if !ok {
		t.Error("Expected day 2024-4-29 to exist")
		return
	}
	if len(day.WaterJson) <= 4 {
		t.Error("Expected WaterJson for 2024-4-29 to be not empty")
		return
	}

	day, ok = storage.Days[utils.Date(2024, 5, 1)]
	if !ok {
		t.Error("Expected day 2024-5-1 to exist")
		return
	}
	if len(day.WaterJson) <= 4 {
		t.Error("Expected WaterJson for 2024-5-1 to be not empty")
		return
	}

}

func TestExportDaysSkipExistedJsons(t *testing.T) {
	e := NewDaysExporter()

	mockSets := mockyzapi.NewMockSettings()

	storage := daysstorage.NewEmpty()

	// fill storage with days not containing water jsons (2024-04-01 - 2024-04-30)
	e.ExportDaysFromYazioToStorage(utils.Month(2024, 04), utils.Date(2024, 04, 30), storage,
		JsonsForRequest{false, true, false, false},
		mockyzapi.NewMockClientFactory(mockSets))

	// add to storage 10 days with water jsons (2024-04-01 - 2024-04-10)
	e.ExportDaysFromYazioToStorage(utils.Month(2024, 04), utils.Date(2024, 04, 10), storage,
		JsonsForRequest{false, false, false, true},
		mockyzapi.NewMockClientFactory(mockSets))

	waterCalls := 0

	writer := callbackWriter(func(data []byte) (int, error) {
		if strings.Contains(string(data), "GetWater") {
			waterCalls++
		}
		log.Print(string(data))
		return len(data), nil
	})

	mockSets.Logger = log.New(writer, "", 0)

	// try to fill all month with water jsons (2024-04-01 - 2024-04-30), because there are 10 days
	// 	with water jsons, there should be only 30-10=20 GetWater calls
	e.ExportDaysFromYazioToStorage(utils.Month(2024, 4), utils.Date(2024, 4, 30), storage,
		JsonsForRequest{false, false, false, true},
		mockyzapi.NewMockClientFactory(mockSets))

	if waterCalls != 20 {
		t.Error("Expected waterCalls to be 20, got ", waterCalls)
		return
	}
}
