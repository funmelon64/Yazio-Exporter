package service

import (
	"YazioExporter/internal/yazioscraper"
	"YazioExporter/pkg/yzapi"
	"encoding/json"
	"fmt"
	"regexp"
)

type prodExporter struct{}

func NewProductsExporter() *prodExporter {
	return &prodExporter{}
}

func (pe *prodExporter) ExportProductsFromYazioToJson(jsonStr string, yzFactory yzapi.ClientFactory) ([]byte, error) {
	prodIds := getUniqueProdIdsFromJson(jsonStr)

	prodsJsons := map[string]json.RawMessage{}

	const maxWorkers = 5
	scraper := yazioscraper.NewYazioJsonsScraper[string](prodIds, maxWorkers, yzFactory)

	scraper.Scrape(func(client yzapi.Client, task string) (string, error) {
		return client.GetProduct(task)
	}, func(results map[string]string) {
		for prodId, prodJson := range results {
			prodsJsons[prodId] = json.RawMessage(prodJson)
		}
	})

	resultJson, err := json.MarshalIndent(prodsJsons, "", "\t")
	if err != nil {
		return nil, fmt.Errorf("fail to marshal result: %v\n%v", err, prodsJsons)
	}

	return resultJson, nil
}

func getUniqueProdIdsFromJson(jsonStr string) (productIds []string) {
	alreadyFindedProdIds := make(map[string]bool)
	prodIdRe := regexp.MustCompile(`(?m)"product_id" *: *"(.*?)"`)
	for _, match := range prodIdRe.FindAllStringSubmatch(string(jsonStr), -1) {
		if match[1] != "" {
			if _, alreadyExists := alreadyFindedProdIds[match[1]]; !alreadyExists {
				productIds = append(productIds, match[1])
				alreadyFindedProdIds[match[1]] = true
			}
		}
	}
	return
}
