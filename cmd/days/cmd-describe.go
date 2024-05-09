package days

import (
	"YazioExporter/pkg/service"
	"github.com/urfave/cli/v2"
	"time"
)

func NewCliCmd(dayExporter service.DaysExporter) *cli.Command {
	return &cli.Command{
		Name:        "days",
		Description: "Downloads all days for specific period and save in JSON file",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "what", Value: "all",
				Usage: "choose day results for export: consumed,goals,exercises,water\n" +
					"Examples:\n" +
					"\t-what all - export all\n" +
					"\t-what consumed,goals - export only consumed and goals"},
			&cli.TimestampFlag{Name: "from", Usage: "date in format yyyy-mm-dd. if omitted, then from first day of diary", Layout: time.DateOnly},
			&cli.TimestampFlag{Name: "to", Usage: "date in format yyyy-mm-dd. if omitted, then to last day of diary", Layout: time.DateOnly},
			&cli.PathFlag{Name: "out", Value: "days.json"},
			&cli.StringFlag{Name: "token", Required: true},
		},
		Usage: "Export all data from 04.05.2024 to 07.05.2024 and save to days.json:\n" +
			"\tdays -what=all -token token.txt -from 04.05.2024 -to 07.05.2024 -out days.json\n" +
			"Export all consumed and goals until 24.04.2024 and save to days.json:\n" +
			"\tdays -what=consumed,goals -token token.txt -to 24.04.2024\n" +
			"Export all consumed:\n" +
			"\t days -what=consumed, -token token.txt\n",
		Action: func(c *cli.Context) error { return daysCmdCtrl(c, dayExporter) },
	}
}
