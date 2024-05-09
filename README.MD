# Yazio Exporter

Application for export your diary to JSON

## Usage

**First**, you need login to your Yazio account and get authorization token:

```bash
./YazioExport login yourmail yourpass -out token.txt
```

**Next**, you can export various result of your days: consumed products, goals, exercises, water intake (daily data are exporting by default, no need to specify)

Example: export all results, for period of 7 april to 7 may:

```bash
./YazioExport days -token token.txt -what all -from 2024-04-07 -to 2024-05-07 -out days.json
```

Example: export only consumed and goals, for all days

```bash
./YazioExport days -token token.txt -what consumed,goals -out days.json
```

**Next** you can export products names and nutrients for data that you exported previously

```bash
./YazioExport products -token token.txt -from days.json
```

## Export format

### Days data

All day results will be saved to JSON file of next structure:

```json
{
    "yyyy-mm-dd": {
        "daily": {...},
        "consumed": {...}, 
        "goals": {...},
        "exercises": {...},
        "water": {...}
    },
    ...
}
```

**Where each have next structures**:

`daily`:

```json
  {
    "date": "2024-03-14",
    "energy": 1636.31,
    "carb": 51.7914,
    "protein": 129.6612,
    "fat": 96.1132,
    "energy_goal": 2000
  },
```

`consumed`:

```json
{
    "products": [
        {
            "id": "aa690403-06d8-47bd-b812-53b055a0aace",
            "date": "2023-09-24 13:29:04",
            "daytime": "lunch",
            "type": "product",
            "product_id": "9f89293c-becf-11e6-bcc7-e0071b8a8723",
            "amount": 50,
            "serving": "gram",
            "serving_quantity": 50
        },
        ...
    ],
    "recipe_portions": [
        {
            "id": "edae5d06-4802-4be0-9c1d-fd196e9ea90e",
            "date": "2023-09-24 12:18:57",
            "daytime": "lunch",
            "type": "recipe_portion",
            "recipe_id": "27e0f3ef-44c9-4bee-9c78-136e697ec81e",
            "portion_count": 4
        },
        ...
    ],
    "simple_products": []
}
```

`goals`

```json
{
  "energy.energy": 2000,
  "nutrient.protein": 156.0976,
  "nutrient.fat": 113.9785,
  "nutrient.carb": 73.1708,
  "activity.step": 10000,
  "bodyvalue.weight": 61,
  "water": 2000
}
```

`exercises`:

```json
{
  "training":[],
  "custom_training":[]
}
```

`water`:

```json
{
  "water_intake": 50,
  "gateway": null,
  "source": null
}
```

### Products

Products are resolved for each occurance of "product_id" key in JSON of your exported days, and have next format (`nutrients` field can be more detailed for various products):

```json
{
  "051bb345-12b3-471e-83f3-ce82af3faf89": {
    "name": "Potato",
    "is_verified": false,
    "is_private": false,
    "is_deleted": false,
    "has_ean": false,
    "category": "potatoproducts",
    "image": "https://images.yazio.com/potatoproducts.jpg",
    "producer": null,
    "nutrients": {
      "energy.energy": 0.857143,
      "nutrient.carb": 0.171429,
      "nutrient.fat": 0.014286,
      "nutrient.protein": 0.028571
    },
    "updated_at": null,
    "servings": [],
    "base_unit": "g"
  },
  ...
}
```