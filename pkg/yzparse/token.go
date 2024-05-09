package yzparse

import (
	"fmt"
	"regexp"
)

var tokenRe = regexp.MustCompile(`(?m)"access_token" *: *"(.*?)"`)

func ParseTokenJson(jsonStr string) (string, error) {
	match := tokenRe.FindStringSubmatch(jsonStr)
	if match[1] == "" {
		return "", fmt.Errorf("json not contain token")
	}
	return match[1], nil
}
