package products

import (
	"YazioExporter/cmd/utils"
	"YazioExporter/pkg/service"
	"YazioExporter/pkg/yzapi"
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
)

func getProducts(c *cli.Context, pe service.ProductsExporter) error {
	fileNameArg := c.String("from")
	if fileNameArg == "" {
		return errors.New("please specify a file to get products from")
	}

	token, err := utils.GetTokenFromArg(c)
	if err != nil {
		return fmt.Errorf("token err: %v", err)
	}

	outFileName := c.String("o")

	content, err := utils.CreateOrOpenFileAndRead(fileNameArg)
	if err != nil {
		return fmt.Errorf("could not open file %s: %v", fileNameArg, err)
	}

	exportsJson, err := pe.ExportProductsFromYazioToJson(string(content), yzapi.NewYzClientFactory(token))
	if err != nil {
		return err
	}

	err = utils.CreateOrOpenFileAndRewrite(outFileName, exportsJson)
	if err != nil {
		return err
	}

	return nil
}
