package service

import (
	"YazioExporter/pkg/yzapi"
	"YazioExporter/pkg/yzparse"
	"fmt"
)

type loginer struct{}

func NewLoginer() *loginer { return &loginer{} }

func (l *loginer) GetLoginToken(mail string, pass string, yzFactory yzapi.ClientFactory) (string, error) {
	yazio := yzFactory.NewClient()

	result, err := yazio.GetLoginToken(mail, pass)
	if err != nil {
		return "", fmt.Errorf("login request failed: %v", err)
	}

	token, err := yzparse.ParseTokenJson(result)
	if err != nil {
		return "", fmt.Errorf("parse login response failed: %v\n\t%v", result, err)
	}

	return token, nil
}
