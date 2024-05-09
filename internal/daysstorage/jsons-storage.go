package daysstorage

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"time"
)

type Storage struct {
	Days map[time.Time]dayJsons
}

type JsonsField int
const (
	DAILY     JsonsField = iota
	CONSUMED  JsonsField = iota
	GOALS     JsonsField = iota
	EXERCISES JsonsField = iota
	WATER     JsonsField = iota
)

func (s *Storage) AddJson(jsons map[time.Time]string, field JsonsField) {
	for date, jsonStr := range jsons {
		if len(jsonStr) <= 2 {
			continue
		}
		day := s.Days[date]

		reflect.ValueOf(&day).Elem().
			FieldByIndex([]int{int(field)}).Set(reflect.ValueOf(json.RawMessage(jsonStr)))

		s.Days[date] = day
	}
}

func (s *Storage) IsInStorage(date time.Time, field JsonsField) bool {
	day, ok := s.Days[date]
	if !ok { return false }
	val := string(reflect.ValueOf(&day).Elem().FieldByIndex([]int{int(field)}).Interface().(json.RawMessage))
	return len(val) > 2 && val != "null"
}

func NewEmpty() *Storage {
	return &Storage{ Days: make(map[time.Time]dayJsons, 0) }
}

// FromJson creates and returns empty Storage if error != nil
// if string is empty returns empty Storage with no error
func FromJson(jsonStr []byte) (*Storage, error) {
	storage := NewEmpty()

	if len(jsonStr) == 0 {
		return storage, nil
	}

	err := storage.parseFromJson(jsonStr)
	if err != nil {
		return storage, fmt.Errorf("fail to parse JSON storage: %v", err);
	}

	return storage, nil
}

func (s *Storage) ToJson() ([]byte, error) {
	daysWithStrDate := map[string]dayJsons{}
	for date, day := range s.Days {
		daysWithStrDate[date.Format(time.DateOnly)] = day
	}
	result, err := json.MarshalIndent(daysWithStrDate, "", "\t")
	if err != nil { return []byte(""), err }
	return result, nil
}

type dayJsons struct {
	DailyJson json.RawMessage `json:"daily"`
	ConsumedJson json.RawMessage `json:"consumed"`
	GoalsJson json.RawMessage `json:"goals"`
	ExercisesJson json.RawMessage `json:"exercises"`
	WaterJson json.RawMessage `json:"water"`
}

func (s *Storage) parseFromJson(jsonStr []byte) error {
	daysWithStrDate := map[string]dayJsons{}
	err := json.Unmarshal(jsonStr, &daysWithStrDate)
	if err != nil { return err }

	for strDate, day := range daysWithStrDate {
		date, err := time.Parse(time.DateOnly, strDate)
		if err != nil {
			log.Printf(
				"[Storage Parse] Warning: unable to parse record with date %s, skip",
				strDate)
			continue
		}
		s.Days[date] = day
	}
	return nil
}