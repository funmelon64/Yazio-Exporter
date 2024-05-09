package yzapi

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

const base = "https://yzapi.yazio.com"
const apiLogin = "/v9/oauth/token"
const apiLoginBody = "{\"client_id\":\"1_4hiybetvfksgw40o0sog4s884kwc840wwso8go4k8c04goo4c\",\"client_secret\":\"6rok2m65xuskgkgogw40wkkk8sw0osg84s8cggsc4woos4s8o\",\"username\":\"%s\",\"password\":\"%s\",\"grant_type\":\"password\"}"
const apiConsumed = "/v9/user/consumed-items?date=%s"
const apiProduct = "/v9/products/%s"
const apiNutrsDaily = "/v9/user/consumed-items/nutrients-daily?start=%s&end=%s"
const apiRecipe = "/v9/recipes/%s"
const apiGoals = "/v9/user/goals?date=%s"
const apiExercises = "/v9/user/exercises?date=%s"
const apiWater = "/v9/user/water-intake?date=%s"

type Client interface {
	GetConsumed(date time.Time) (string, error)
	GetProduct(prodId string) (string, error)
	GetRecipe(recipeId string) (string, error)
	GetGoals(date time.Time) (string, error)
	GetExercises(date time.Time) (string, error)
	GetWater(date time.Time) (string, error)
	GetMonthDiary(month time.Time) (string, error)
	GetLoginToken(mail string, pass string) (string, error)
}

type ClientFactory interface {
	NewClient() Client
}

type client struct {
	httpCl *http.Client
	token  string
}

type clientFactory struct {
	token string
}

func NewYzClientFactory(token string) ClientFactory {
	return clientFactory{token: token}
}

func (f clientFactory) NewClient() Client {
	return client{httpCl: &http.Client{}, token: f.token}
}

func (api client) request(method string, apiPath string, reqBodyStr string) (string, error) {
	log.Printf("[yzapi] request(%v, %v, %v)\n", method, apiPath, reqBodyStr)

	var reqBody io.Reader = nil

	if reqBodyStr != "" {
		reqBody = strings.NewReader(reqBodyStr)
	}

	req, err := http.NewRequest(method, base+apiPath, reqBody)
	if err != nil {
		return "", fmt.Errorf("Construct Request Error: %v", err)
	}

	req.Header.Add("Authorization", "Bearer "+api.token)

	resp, err := api.httpCl.Do(req)
	if err != nil {
		return "", fmt.Errorf("Api Request Error: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Request Status Code: %v", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Reading Request Error: %v", err)
	}

	return string(body), nil
}

func (api client) GetConsumed(date time.Time) (string, error) {
	return api.request("GET", fmt.Sprintf(apiConsumed, date.Format(time.DateOnly)), "")
}

func (api client) GetProduct(prodId string) (string, error) {
	return api.request("GET", fmt.Sprintf(apiProduct, prodId), "")
}

func (api client) GetRecipe(recipeId string) (string, error) {
	return api.request("GET", fmt.Sprintf(apiRecipe, recipeId), "")
}

func (api client) GetMonthDiary(month time.Time) (string, error) {
	currentYear, currentMonth, _ := month.Date()
	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, month.Location())
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

	return api.request("GET", fmt.Sprintf(apiNutrsDaily,
		firstOfMonth.Format(time.DateOnly), lastOfMonth.Format(time.DateOnly)), "")
}

func (api client) GetGoals(date time.Time) (string, error) {
	return api.request("GET", fmt.Sprintf(apiGoals, date.Format(time.DateOnly)), "")
}

func (api client) GetExercises(date time.Time) (string, error) {
	return api.request("GET", fmt.Sprintf(apiExercises, date.Format(time.DateOnly)), "")
}

func (api client) GetWater(date time.Time) (string, error) {
	return api.request("GET", fmt.Sprintf(apiWater, date.Format(time.DateOnly)), "")
}

func (api client) GetLoginToken(mail string, pass string) (string, error) {
	return api.request("POST", apiLogin, fmt.Sprintf(apiLoginBody, mail, pass))
}
