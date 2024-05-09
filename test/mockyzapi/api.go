package mockyzapi

import (
	"YazioExporter/cmd/utils"
	"YazioExporter/pkg/yzapi"
	"fmt"
	"log"
	"time"
)

const base = "https://yzapi.yazio.com"
const apiConsumed = "/v9/user/consumed-items?date=%s"
const apiProduct = "/v9/products/%s"
const apiNutrsDaily = "/v9/user/consumed-items/nutrients-daily?start=%s&end=%s"
const apiRecipe = "/v9/recipes/%s"
const apiGoals = "/v9/goals?date=%s"
const apiExercises = "/v9/user/exercises?date=%s"
const apiWater = "/v9/user/water-intake?date=%s"

func rndSlp() {
	//time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
}

type mockSettings struct {
	DailyDayFrom time.Time
	DailyDayTo   time.Time
	Logger       *log.Logger
	Token        string
}

func NewMockSettings() mockSettings {
	const token = "d044bb9cd7ad7b6b97ba5d391c603b4a176d451dad3a18be63aa26c50e0aa961d1350bb3d4afc1339739a4"
	return mockSettings{DailyDayFrom: time.Date(2021, 8, 0, 0, 0, 0, 0, time.UTC),
		DailyDayTo: time.Now(), Logger: log.Default(), Token: token}
}

type clientFactory struct {
	settings mockSettings
	logger   *log.Logger
}

func NewMockClientFactory(settings mockSettings) yzapi.ClientFactory {
	return clientFactory{settings, settings.Logger}
}

func (f clientFactory) NewClient() yzapi.Client {
	return client{f.settings, f.logger}
}

type client struct {
	settings mockSettings
	logger   *log.Logger
}

func (api client) GetLoginToken(mail string, pass string) (string, error) {
	const answer = "{\"access_token\":\"%s\",\"expires_in\":172800,\"refresh_token\":\"fa0e8e2542693f7aa73ec62dfafd46f4800e1667e238724e8d87d90faf3000a490d52d3d6c4a8b13c9a351\",\"token_type\":\"bearer\"}"
	rndSlp()
	api.logger.Printf("[MockYzapi] GetConsumed(%v, %v)\n", mail, pass)
	return fmt.Sprintf(answer, api.settings.Token), nil
}

func (api client) GetConsumed(date time.Time) (string, error) {
	rndSlp()
	api.logger.Printf("[MockYzapi] GetConsumed(%v)\n", date)
	return "{\n    \"products\": [\n        {\n            \"id\": \"aa690403-06d8-47bd-b812-53b055a0aace\",\n            \"date\": \"2023-09-24 13:29:04\",\n            \"daytime\": \"lunch\",\n            \"type\": \"product\",\n            \"product_id\": \"9f89293c-becf-11e6-bcc7-e0071b8a8723\",\n            \"amount\": 50,\n            \"serving\": \"gram\",\n            \"serving_quantity\": 50\n        },\n        {\n            \"id\": \"c682224d-106f-4c98-bfe8-647797753f26\",\n            \"date\": \"2023-09-24 13:29:32\",\n            \"daytime\": \"lunch\",\n            \"type\": \"product\",\n            \"product_id\": \"0975b6b7-6268-4368-8de2-7b9a6e3e8741\",\n            \"amount\": 60,\n            \"serving\": \"gram\",\n            \"serving_quantity\": 60\n        },\n        {\n            \"id\": \"f94622e9-2905-41f8-b3fe-b158ff43a4aa\",\n            \"date\": \"2023-09-24 13:29:40\",\n            \"daytime\": \"lunch\",\n            \"type\": \"product\",\n            \"product_id\": \"a4351aa4-becf-11e6-b202-e0071b8a8723\",\n            \"amount\": 8,\n            \"serving\": \"gram\",\n            \"serving_quantity\": 8\n        },\n        {\n            \"id\": \"0d1d1a25-2d09-4036-9bea-32dae2d504c6\",\n            \"date\": \"2023-09-24 13:30:15\",\n            \"daytime\": \"dinner\",\n            \"type\": \"product\",\n            \"product_id\": \"a9add8f4-becf-11e6-a170-e0071b8a8723\",\n            \"amount\": 110,\n            \"serving\": \"gram\",\n            \"serving_quantity\": 110\n        },\n        {\n            \"id\": \"cfecfb90-2450-461e-8b83-083606b69dd2\",\n            \"date\": \"2023-09-24 17:22:08\",\n            \"daytime\": \"lunch\",\n            \"type\": \"product\",\n            \"product_id\": \"9f89971e-becf-11e6-9aab-e0071b8a8723\",\n            \"amount\": 60,\n            \"serving\": \"gram\",\n            \"serving_quantity\": 60\n        },\n        {\n            \"id\": \"fe208253-c4cf-4a7e-91d7-efdc8f5eed10\",\n            \"date\": \"2023-09-24 17:22:19\",\n            \"daytime\": \"lunch\",\n            \"type\": \"product\",\n            \"product_id\": \"181fa6e2-302a-4bad-bf75-c521c22cb812\",\n            \"amount\": 35,\n            \"serving\": \"gram\",\n            \"serving_quantity\": 35\n        },\n        {\n            \"id\": \"d7f47385-cac0-4c01-9615-890c8fc2178f\",\n            \"date\": \"2023-09-24 19:30:52\",\n            \"daytime\": \"breakfast\",\n            \"type\": \"product\",\n            \"product_id\": \"181fa6e2-302a-4bad-bf75-c521c22cb812\",\n            \"amount\": 30,\n            \"serving\": \"gram\",\n            \"serving_quantity\": 30\n        },\n        {\n            \"id\": \"d7be24ea-4679-4627-82c4-9ac91e0865a2\",\n            \"date\": \"2023-09-24 19:31:46\",\n            \"daytime\": \"breakfast\",\n            \"type\": \"product\",\n            \"product_id\": \"9f89ac7c-becf-11e6-b84b-e0071b8a8723\",\n            \"amount\": 100,\n            \"serving\": \"gram\",\n            \"serving_quantity\": 100\n        },\n        {\n            \"id\": \"cb0227c5-61de-4803-bbf2-c968dd34c5ea\",\n            \"date\": \"2023-09-24 20:08:36\",\n            \"daytime\": \"dinner\",\n            \"type\": \"product\",\n            \"product_id\": \"a9bcb333-1575-473a-aeb7-658dc17f24ae\",\n            \"amount\": 60,\n            \"serving\": \"gram\",\n            \"serving_quantity\": 60\n        }\n    ],\n    \"recipe_portions\": [\n        {\n            \"id\": \"edae5d06-4802-4be0-9c1d-fd196e9ea90e\",\n            \"date\": \"2023-09-24 12:18:57\",\n            \"daytime\": \"lunch\",\n            \"type\": \"recipe_portion\",\n            \"recipe_id\": \"27e0f3ef-44c9-4bee-9c78-136e697ec81e\",\n            \"portion_count\": 4\n        },\n        {\n            \"id\": \"a1435d81-29ec-48cf-b247-96d969a56eae\",\n            \"date\": \"2023-09-24 13:30:42\",\n            \"daytime\": \"dinner\",\n            \"type\": \"recipe_portion\",\n            \"recipe_id\": \"27e0f3ef-44c9-4bee-9c78-136e697ec81e\",\n            \"portion_count\": 2\n        }\n    ],\n    \"simple_products\": []\n}", nil
}

func (api client) GetProduct(prodId string) (string, error) {
	rndSlp()
	api.logger.Printf("[MockYzapi] GetProduct(%v)\n", prodId)
	return "{\n  \"name\": \"Рыба\",\n  \"is_verified\": false,\n  \"is_private\": false,\n  \"is_deleted\": false,\n  \"has_ean\": false,\n  \"category\": \"fish\",\n  \"image\": \"https://images.yazio.com/fish.jpg\",\n  \"producer\": \"Рыба филе жареная\",\n  \"nutrients\": {\n    \"energy.energy\": 1.9,\n    \"nutrient.carb\": 0.0616,\n    \"nutrient.fat\": 0.09910000000000001,\n    \"nutrient.protein\": 0.17309999999999998\n  },\n  \"updated_at\": null,\n  \"servings\": [],\n  \"base_unit\": \"g\"\n}", nil
}

func (api client) GetRecipe(recipeId string) (string, error) {
	rndSlp()
	api.logger.Printf("[MockYzapi] GetRecipe(%v)\n", recipeId)
	return "{\n    \"id\": \"27e0f3ef-44c9-4bee-9c78-136e697ec81e\",\n    \"yazio_id\": null,\n    \"locale\": \"ru\",\n    \"name\": \"Блинчики\",\n    \"portion_count\": 1,\n    \"nutrients\": {\n        \"energy.energy\": 119.34999999999999,\n        \"mineral.calcium\": 0.050200000000000002,\n        \"mineral.copper\": 1.22e-5,\n        \"mineral.iron\": 0.00066049999999999995,\n        \"mineral.magnesium\": 0.0074999999999999997,\n        \"mineral.manganese\": 1.5100000000000001e-5,\n        \"mineral.phosphorus\": 0.11147499999999999,\n        \"mineral.potassium\": 0.095825000000000007,\n        \"mineral.selenium\": 1.77e-5,\n        \"mineral.zinc\": 0.00065180000000000001,\n        \"nutrient.carb\": 1.5760000000000001,\n        \"nutrient.cholesterol\": 0.21355000000000002,\n        \"nutrient.dietaryfiber\": 0.0,\n        \"nutrient.fat\": 8.9725000000000001,\n        \"nutrient.monounsaturated\": 3.1231,\n        \"nutrient.polyunsaturated\": 0.90910000000000002,\n        \"nutrient.protein\": 7.556,\n        \"nutrient.saturated\": 3.718,\n        \"nutrient.sodium\": 0.076899999999999996,\n        \"nutrient.sugar\": 1.6259999999999999,\n        \"nutrient.water\": 58.673000000000002,\n        \"vitamin.a\": 0.000111,\n        \"vitamin.b1\": 4.5500000000000001e-5,\n        \"vitamin.b11\": 2.5199999999999999e-5,\n        \"vitamin.b12\": 6.9999999999999997e-7,\n        \"vitamin.b2\": 0.00031609999999999999,\n        \"vitamin.b3\": 5.3100000000000003e-5,\n        \"vitamin.b5\": 0.00084380000000000002,\n        \"vitamin.b6\": 7.3800000000000005e-5,\n        \"vitamin.c\": 0.0,\n        \"vitamin.d\": 1.5e-6,\n        \"vitamin.e\": 0.00065049999999999993,\n        \"vitamin.k\": 4.0000000000000003e-7\n    },\n    \"image\": null,\n    \"servings\": [\n        {\n            \"producer\": null,\n            \"name\": \"Яйцо, отварное\",\n            \"amount\": 55.0,\n            \"serving\": \"gram\",\n            \"serving_quantity\": 55.0,\n            \"base_unit\": \"g\",\n            \"note\": null,\n            \"product_id\": \"9f88ab74-becf-11e6-b54c-e0071b8a8723\"\n        },\n        {\n            \"producer\": null,\n            \"name\": \"Сливочное масло\",\n            \"amount\": 2.5,\n            \"serving\": \"gram\",\n            \"serving_quantity\": 2.5,\n            \"base_unit\": \"g\",\n            \"note\": null,\n            \"product_id\": \"9f87b764-becf-11e6-ace8-e0071b8a8723\"\n        },\n        {\n            \"producer\": null,\n            \"name\": \"Молоко\",\n            \"amount\": 20.0,\n            \"serving\": \"gram\",\n            \"serving_quantity\": 20.0,\n            \"base_unit\": \"g\",\n            \"note\": null,\n            \"product_id\": \"9f89ac7c-becf-11e6-b84b-e0071b8a8723\"\n        }\n    ],\n    \"instructions\": [],\n    \"is_yazio_recipe\": false,\n    \"available_since\": null,\n    \"is_pro_recipe\": true\n}", nil
}

func (api client) GetMonthDiary(month time.Time) (string, error) {
	month = utils.TruncToMonth(month)
	rndSlp()
	api.logger.Printf("[MockYzapi] GetMonthDiary(%v)\n", month)
	if month.After(utils.TruncToMonth(api.settings.DailyDayTo)) ||
		month.Before(utils.TruncToMonth(api.settings.DailyDayFrom)) {
		return "[]", nil
	}

	firstDay := 1
	lastDay := utils.LastMonthDay(month).Day()
	if month == utils.TruncToMonth(api.settings.DailyDayFrom) {
		firstDay = api.settings.DailyDayFrom.Day()
	}
	if month == utils.TruncToMonth(api.settings.DailyDayTo) {
		lastDay = api.settings.DailyDayTo.Day()
	}

	dailyJson := ""
	for day := firstDay; day <= lastDay; day++ {
		dailyJson += fmt.Sprintf("{\"date\":\"%s\"},", utils.Date(month.Year(), month.Month(), day).Format(time.DateOnly))
	}
	dailyJson = dailyJson[:len(dailyJson)-1] // remove last comma
	dailyJson = fmt.Sprintf("[%s]", dailyJson)

	return dailyJson, nil
}

func (api client) GetGoals(date time.Time) (string, error) {
	rndSlp()
	api.logger.Printf("[MockYzapi] GetGoals(%v)\n)", date)
	//return "", fmt.Errorf("Not implemented")
	return "{\n  \"energy.energy\": 2000,\n  \"nutrient.protein\": 156.0976,\n  \"nutrient.fat\": 113.9785,\n  \"nutrient.carb\": 73.1708,\n  \"activity.step\": 10000,\n  \"bodyvalue.weight\": 61,\n  \"water\": 2000\n}", nil
}

func (api client) GetExercises(date time.Time) (string, error) {
	rndSlp()
	api.logger.Printf("[MockYzapi] GetExercises(%v)\n)", date)
	return "{\n  \"training\":[],\n  \"custom_training\":[]\n}", nil
}

func (api client) GetWater(date time.Time) (string, error) {
	rndSlp()
	api.logger.Printf("[MockYzapi] GetWater(%v)\n)", date)
	return "{\n  \"water_intake\": 0,\n  \"gateway\": null,\n  \"source\": null\n}", nil
}
