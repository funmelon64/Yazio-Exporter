package cmd

import (
	"YazioExporter/cmd/days"
	"YazioExporter/cmd/login"
	"YazioExporter/cmd/products"
	service2 "YazioExporter/pkg/service"
	"github.com/urfave/cli/v2"
)

func Init() *cli.App {
	app := &cli.App{
		Name:    "Yazio Exporter",
		Usage:   "Application for exporting Yazio diary",
		Version: "0.0.1",
		Authors: []*cli.Author{&cli.Author{Name: "Morph", Email: "morph@fm64.me"}},
		Commands: []*cli.Command{
			days.NewCliCmd(service2.NewDaysExporter()),
			products.NewCliCmd(service2.NewProductsExporter()),
			login.NewCliCmd(service2.NewLoginer()),
		},
	}

	return app
}
