package cmd

import (
	"YazioExporter/cmd/days"
	"YazioExporter/cmd/login"
	"YazioExporter/cmd/products"
	"YazioExporter/test/mockservice"
	"github.com/urfave/cli/v2"
	"strings"
	"testing"
)

func testApp() *cli.App {
	app := &cli.App{
		Name:    "Yazio Parser",
		Authors: []*cli.Author{&cli.Author{Name: "Morph", Email: "morph@fm64.me"}},
		Commands: []*cli.Command{
			days.NewCliCmd(mockservice.NewDaysExporter()),
			products.NewCliCmd(mockservice.NewProductsExporter()),
			login.NewCliCmd(mockservice.NewLoginer()),
		},
	}

	return app
}

func TestGetDays(t *testing.T) {
	testTable := []string{
		"yzexport days -token token.txt -from 2024-03-02",
		"yzexport days",
		"yzexport days -token kek.txt",
		"yzexport days -token token.txt -to 2024-05-09",
		"yzexport days -token token.txt -from 09.05.2024 -to 2024-01-04",
		"yzexport days -token token.txt -what=consumed,goals",
	}

	for _, test := range testTable {
		t.Log(test, "\n")
		t.Log(testApp().Run(strings.Split(test, " ")), "\n")
	}

	t.Log("=== Please hand-check this test, automatic test not implemented!!! ===")
}
