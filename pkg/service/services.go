package service

import (
	"YazioExporter/internal/daysstorage"
	"YazioExporter/pkg/yzapi"
	"time"
)

type DaysExporter interface {
	// ExportDaysFromYazioToStorage
	//
	// if dateFrom.IsZero() then exports days from first day in diary.
	//
	// if dateTo.IsZero() then exports days until last day in diary
	ExportDaysFromYazioToStorage(dateFrom time.Time, dateTo time.Time, storage *daysstorage.Storage,
		whatRequesting JsonsForRequest, yzFactory yzapi.ClientFactory)
}

type ProductsExporter interface {
	ExportProductsFromYazioToJson(jsonStr string, yzFactory yzapi.ClientFactory) ([]byte, error)
}

type Loginer interface {
	GetLoginToken(mail string, pass string, yzFactory yzapi.ClientFactory) (string, error)
}
