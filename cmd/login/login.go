package login

import (
	"YazioExporter/cmd/utils"
	"YazioExporter/pkg/service"
	"YazioExporter/pkg/yzapi"
	"fmt"
	"github.com/urfave/cli/v2"
)

func login(c *cli.Context, l service.Loginer) error {
	mail := c.Args().First()
	pass := c.Args().Get(1)

	if mail == "" || pass == "" {
		return fmt.Errorf("login requires mail and pass arguments")
	}

	tokenPath := c.Path("out")

	token, err := l.GetLoginToken(mail, pass, yzapi.NewYzClientFactory(""))
	if err != nil {
		return fmt.Errorf("login error: %v", err)
	}

	err = utils.CreateOrOpenFileAndRewrite(tokenPath, []byte(token))
	if err != nil {
		return fmt.Errorf("fail to write token file (%v): %v", tokenPath, err)
	}

	return nil
}
