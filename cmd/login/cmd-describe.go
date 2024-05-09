package login

import (
	"YazioExporter/pkg/service"
	"github.com/urfave/cli/v2"
)

func NewCliCmd(loginer service.Loginer) *cli.Command {
	return &cli.Command{
		Name:        "login",
		Description: "logins to Yazio, gets token and save it to file",
		Args:        true,
		ArgsUsage:   "mail pass",
		Flags: []cli.Flag{
			&cli.PathFlag{Name: "out", Value: "token.txt"},
		},
		Usage:  "yzexport login example@mail.com yourpassword -o token.txt",
		Action: func(c *cli.Context) error { return login(c, loginer) },
	}
}
