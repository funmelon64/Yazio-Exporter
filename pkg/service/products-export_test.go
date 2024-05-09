package service

import (
	"YazioExporter/test/mockyzapi"
	"encoding/json"
	"log"
	"testing"
)

func TestProductsExport(t *testing.T) {
	const cont = "{\n    \"products\": [\n        {\n            \"id\": \"aa690403-06d8-47bd-b812-53b055a0aace\",\n            \"date\": \"2023-09-24 13:29:04\",\n            \"daytime\": \"lunch\",\n            \"type\": \"product\",\n            \"product_id\": \"9f89293c-becf-11e6-bcc7-e0071b8a8723\",\n            \"amount\": 50,\n            \"serving\": \"gram\",\n            \"serving_quantity\": 50\n        },\n        {\n            \"id\": \"c682224d-106f-4c98-bfe8-647797753f26\",\n            \"date\": \"2023-09-24 13:29:32\",\n            \"daytime\": \"lunch\",\n            \"type\": \"product\",\n            \"product_id\": \"0975b6b7-6268-4368-8de2-7b9a6e3e8741\",\n            \"amount\": 60,\n            \"serving\": \"gram\",\n            \"serving_quantity\": 60\n        },\n        {\n            \"id\": \"f94622e9-2905-41f8-b3fe-b158ff43a4aa\",\n            \"date\": \"2023-09-24 13:29:40\",\n            \"daytime\": \"lunch\",\n            \"type\": \"product\",\n            \"product_id\": \"a4351aa4-becf-11e6-b202-e0071b8a8723\",\n            \"amount\": 8,\n            \"serving\": \"gram\",\n            \"serving_quantity\": 8\n        },\n        {\n            \"id\": \"0d1d1a25-2d09-4036-9bea-32dae2d504c6\",\n            \"date\": \"2023-09-24 13:30:15\",\n            \"daytime\": \"dinner\",\n            \"type\": \"product\",\n            \"product_id\": \"a9add8f4-becf-11e6-a170-e0071b8a8723\",\n            \"amount\": 110,\n            \"serving\": \"gram\",\n            \"serving_quantity\": 110\n        },\n        {\n            \"id\": \"cfecfb90-2450-461e-8b83-083606b69dd2\",\n            \"date\": \"2023-09-24 17:22:08\",\n            \"daytime\": \"lunch\",\n            \"type\": \"product\",\n            \"product_id\": \"9f89971e-becf-11e6-9aab-e0071b8a8723\",\n            \"amount\": 60,\n            \"serving\": \"gram\",\n            \"serving_quantity\": 60\n        },\n        {\n            \"id\": \"fe208253-c4cf-4a7e-91d7-efdc8f5eed10\",\n            \"date\": \"2023-09-24 17:22:19\",\n            \"daytime\": \"lunch\",\n            \"type\": \"product\",\n            \"product_id\": \"181fa6e2-302a-4bad-bf75-c521c22cb812\",\n            \"amount\": 35,\n            \"serving\": \"gram\",\n            \"serving_quantity\": 35\n        },\n        {\n            \"id\": \"d7f47385-cac0-4c01-9615-890c8fc2178f\",\n            \"date\": \"2023-09-24 19:30:52\",\n            \"daytime\": \"breakfast\",\n            \"type\": \"product\",\n            \"product_id\": \"181fa6e2-302a-4bad-bf75-c521c22cb812\",\n            \"amount\": 30,\n            \"serving\": \"gram\",\n            \"serving_quantity\": 30\n        },\n        {\n            \"id\": \"d7be24ea-4679-4627-82c4-9ac91e0865a2\",\n            \"date\": \"2023-09-24 19:31:46\",\n            \"daytime\": \"breakfast\",\n            \"type\": \"product\",\n            \"product_id\": \"9f89ac7c-becf-11e6-b84b-e0071b8a8723\",\n            \"amount\": 100,\n            \"serving\": \"gram\",\n            \"serving_quantity\": 100\n        },\n        {\n            \"id\": \"cb0227c5-61de-4803-bbf2-c968dd34c5ea\",\n            \"date\": \"2023-09-24 20:08:36\",\n            \"daytime\": \"dinner\",\n            \"type\": \"product\",\n            \"product_id\": \"a9bcb333-1575-473a-aeb7-658dc17f24ae\",\n            \"amount\": 60,\n            \"serving\": \"gram\",\n            \"serving_quantity\": 60\n        }\n    ],\n    \"recipe_portions\": [\n        {\n            \"id\": \"edae5d06-4802-4be0-9c1d-fd196e9ea90e\",\n            \"date\": \"2023-09-24 12:18:57\",\n            \"daytime\": \"lunch\",\n            \"type\": \"recipe_portion\",\n            \"recipe_id\": \"27e0f3ef-44c9-4bee-9c78-136e697ec81e\",\n            \"portion_count\": 4\n        },\n        {\n            \"id\": \"a1435d81-29ec-48cf-b247-96d969a56eae\",\n            \"date\": \"2023-09-24 13:30:42\",\n            \"daytime\": \"dinner\",\n            \"type\": \"recipe_portion\",\n            \"recipe_id\": \"27e0f3ef-44c9-4bee-9c78-136e697ec81e\",\n            \"portion_count\": 2\n        }\n    ],\n    \"simple_products\": []\n}"

	exportJson, err := NewProductsExporter().ExportProductsFromYazioToJson(cont, mockyzapi.NewMockClientFactory(mockyzapi.NewMockSettings()))
	if err != nil {
		log.Panicf("fail: %v", err)
	}

	m := map[string]json.RawMessage{}
	err = json.Unmarshal(exportJson, &m)
	if err != nil {
		log.Panicf("fail: %v", err)
	}

	if len(m) != 8 {
		log.Printf("fail, exported ids count is not 8 but %v", len(m))
	}
}
