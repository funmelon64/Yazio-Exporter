package daysstorage

import (
	"testing"
	"time"
)

const content = "{\"2024-05-02\":{\"daily\":{\"date\":\"2024-05-02\"},\"consumed\":{\"products\":{}},\"goals\":{},\"exercises\":{},\"water\":null},\"2024-05-05\":{\"daily\":null,\"consumed\":{},\"goals\":{\"data\":\"...\"},\"exercises\":{},\"water\":{}}}"
func TestStorageFromJson(t *testing.T) {
	storage, err := FromJson([]byte(content))
	if err != nil {
		t.Error(err)
		return
	}
	if len(storage.Days) != 2 {
		t.Error("Days count should be 2")
	}

	day, ok := storage.Days[time.Date(2024,5,2,0,0,0,0,time.UTC)]

	if !ok {
		t.Error(".Days should contain 2024-05-02")
	}
	if len(day.DailyJson) <= 4 {
		t.Error("2024-05-02 should contain non empty daily")
	}

	t.Logf("Storage Days: %v\n", storage.Days)
}

func TestAddToStorage(t *testing.T) {
	storage, err := FromJson([]byte(content))
	if err != nil {
		t.Error(err)
		return
	}

	storage.AddJson(map[time.Time]string{
		time.Date(2024, 05, 01, 0, 0, 0, 0, time.UTC): "GoalsJson",
		time.Date(2024, 05, 05, 0, 0, 0, 0, time.UTC): "RewritedGoals",
	}, GOALS)

	day, ok := storage.Days[time.Date(2024, 05, 01, 0, 0, 0, 0, time.UTC)]
	if !ok {
		t.Error(".Days should contain added 2024-05-01")
		return
	}
	if len(day.GoalsJson) <= 4 {
		t.Error("2024-05-01 should contain non empty goals")
		return
	}

	storage.AddJson(map[time.Time]string{
		time.Date(2024, 05, 05, 0, 0, 0, 0, time.UTC): "",
	}, GOALS)

	day, ok = storage.Days[time.Date(2024, 05, 05, 0, 0, 0, 0, time.UTC)]
	if !ok {
		t.Error(".Days should contain added 2024-05-01")
		return
	}
	if string(day.GoalsJson) != "RewritedGoals" {
		t.Error("added empty string shouldn't rewrite existed goals")
		return
	}

	t.Logf("Storage Days: %v\n", storage.Days)
}

func TestStorageToJson(t *testing.T) {
	storage, err := FromJson([]byte(content))
	if err != nil {
		t.Error(err)
		return
	}

	str, err := storage.ToJson()
	if err != nil {
		t.Error(err)
	}

	t.Log(str)
}

func TestIsInStorage(t *testing.T) {
	const cont = "{\"2024-05-02\":{\"daily\":{\"date\":\"2024-03-30\",\"energy\":2413.892,\"carb\":130.5838,\"protein\":172.1392,\"fat\":131.1716,\"energy_goal\":2000},\"consumed\":\"123\",\"goals\":\"555\",\"exercises\":\"333\",\"water\":{}},\"2024-05-03\":{\"daily\":{\"date\":\"2024-03-228\",\"energy\":45456.892,\"carb\":130.5838,\"protein\":172.1392,\"fat\":131.1716,\"energy_goal\":2000},\"consumed\":{},\"goals\":null,\"exercises\":{},\"water\":[\"123\"]}}"
	storage, err := FromJson([]byte(cont))
	if err != nil {
		t.Error(err)
		return
	}

	date := time.Date(2024, 5, 2, 0, 0, 0, 0, time.UTC)
	if ! storage.IsInStorage(date, DAILY) {
		t.Error("DAILY", date)
	}
	if storage.IsInStorage(date, WATER) {
		t.Error("WATER", date)
	}

	date = time.Date(2024, 5, 8, 0, 0, 0, 0, time.UTC)
	if storage.IsInStorage(date, DAILY) {
		t.Error("DAILY", date)
	}
	if storage.IsInStorage(date, WATER) {
		t.Error("WATER", date)
	}

	date = time.Date(2024, 5, 3, 0, 0, 0, 0, time.UTC)
	if storage.IsInStorage(date, GOALS) {
		t.Error("GOALS", date)
	}
}