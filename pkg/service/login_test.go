package service

import (
	"YazioExporter/test/mockyzapi"
	"testing"
)

func TestLogin(t *testing.T) {
	mockSets := mockyzapi.NewMockSettings()
	token, err := NewLoginer().GetLoginToken("mambich933@mail.ru", "ssuper228",
		mockyzapi.NewMockClientFactory(mockSets))

	if err != nil {
		panic(err)
	}

	if token != mockSets.Token {
		t.Errorf("token does not match")
	}
}
