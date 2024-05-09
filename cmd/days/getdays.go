package days

import (
	"YazioExporter/cmd/utils"
	"YazioExporter/internal/daysstorage"
	service2 "YazioExporter/pkg/service"
	"YazioExporter/pkg/yzapi"
	"fmt"
	"github.com/urfave/cli/v2"
	"strings"
	"time"
)

func daysCmdCtrl(c *cli.Context, de service2.DaysExporter) error {
	token, err := utils.GetTokenFromArg(c)
	if err != nil {
		return fmt.Errorf("fail to get token: %v", err)
	}

	var dateFrom, dateTo time.Time = time.Time{}, time.Time{}
	from := c.Timestamp("from")
	to := c.Timestamp("to")
	if from != nil {
		dateFrom = *from
	}
	if to != nil {
		dateTo = *to
	}

	storageFileName := c.String("out")

	whatArg := c.String("what")
	whatRequesting := service2.JsonsForRequest{}
	if whatArg == "all" {
		whatRequesting = service2.JsonsForRequest{Consumed: true, Goals: true, Exercises: true, Water: true}
	} else {
		whatRequesting = service2.JsonsForRequest{}
		if strings.Contains(whatArg, "consumed") {
			whatRequesting.Consumed = true
		}
		if strings.Contains(whatArg, "goals") {
			whatRequesting.Goals = true
		}
		if strings.Contains(whatArg, "exercises") {
			whatRequesting.Exercises = true
		}
		if strings.Contains(whatArg, "water") {
			whatRequesting.Water = true
		}
	}

	storageFile, err := utils.CreateOrOpenFileAndRead(storageFileName)
	if err != nil {
		return fmt.Errorf("fail to open or create file %s: %v", storageFileName, err)
	}

	storage, err := daysstorage.FromJson(storageFile)
	if err != nil {
		return fmt.Errorf("fail to create or open json storage file %v: %v", storageFileName, err)
	}

	de.ExportDaysFromYazioToStorage(dateFrom, dateTo, storage, whatRequesting, yzapi.NewYzClientFactory(token))

	storageJson, err := storage.ToJson()
	if err != nil {
		return fmt.Errorf("fail to convert storage to json: %v", err)
	}

	if err := utils.CreateOrOpenFileAndRewrite(storageFileName, storageJson); err != nil {
		return fmt.Errorf("fail to save json storage: %v", err)
	}

	return nil
}
