package utils

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"io"
	"os"
)

func GetTokenFromArg(c *cli.Context) (string, error) {
	tokenFile := c.Path("token")
	if tokenFile == "" {
		return "", fmt.Errorf("bad filename passed to -token")
	}
	file, err := os.Open(tokenFile)
	if err != nil {
		return "", fmt.Errorf("fail to open token file (%v): %v", tokenFile, err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("fail to read token file (%v): %v", tokenFile, err)
	}
	return string(content), nil
}
