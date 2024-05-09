package products

import (
	"YazioExporter/pkg/service"
	"github.com/urfave/cli/v2"
)

func NewCliCmd(prodExporter service.ProductsExporter) *cli.Command {
	return &cli.Command{
		Name:        "products",
		Description: "Downloads all products for any format JSON file (search keys product_id)",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "from", Required: true},
			&cli.StringFlag{Name: "o", Value: "products.json"},
			&cli.StringFlag{Name: "token", Required: true},
		},
		Usage:  "products -from days.json -o products.json -token token.txt",
		Action: func(c *cli.Context) error { return getProducts(c, prodExporter) },
	}
}
