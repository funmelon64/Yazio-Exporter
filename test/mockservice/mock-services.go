package mockservice

import (
	"YazioExporter/cmd/utils"
	"YazioExporter/internal/daysstorage"
	"YazioExporter/pkg/service"
	"YazioExporter/pkg/yzapi"
	"log"
	"time"
)

type daysExporter struct{}

func NewDaysExporter() *daysExporter {
	return &daysExporter{}
}

func (de *daysExporter) ExportDaysFromYazioToStorage(dateFrom time.Time, dateTo time.Time, storage *daysstorage.Storage,
	whatRequesting service.JsonsForRequest, yzFactory yzapi.ClientFactory) {
	log.Printf("[MockService] ExportDaysFromYazioToStorage(%v, %v, %v, %v, %v)",
		dateFrom, dateTo, storage, whatRequesting, yzFactory)
	storage.AddJson(map[time.Time]string{
		utils.Date(2024, 05, 03): "\"GoalsJson1\"",
		utils.Date(2024, 05, 04): "\"GoalsJson2\"",
	}, daysstorage.GOALS)
}

type prodExporter struct{}

func NewProductsExporter() *prodExporter {
	return &prodExporter{}
}

func (pe *prodExporter) ExportProductsFromYazioToJson(jsonStr string, yzFactory yzapi.ClientFactory) ([]byte, error) {
	log.Printf("[MockService] ExportProductsFromYazioToJson(%v, %v)", jsonStr, yzFactory)
	return []byte("ProductsJson"), nil
}

type loginer struct{}

func NewLoginer() *loginer { return &loginer{} }

func (l *loginer) GetLoginToken(mail string, pass string, yzFactory yzapi.ClientFactory) (string, error) {
	log.Println("[MockService] GetLoginToken(%v, %v, %v)", mail, pass, yzFactory)
	return "Token", nil
}
