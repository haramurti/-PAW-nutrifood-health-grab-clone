package ai

// Schema prompt untuk Gemini — diisi via fmt.Sprintf sebelum dikirim

const IngredientNutritionSchema = `{
  "description": "buat list ingredient dasar untuk membuat sebuah raw %s dan berikan takaran pasti untuk setiap ingredient nya serta menghitung nutrition berdasarkan ingredient serta menentukan type food atau beverage",
  "type": "object",
  "properties": {
    "type": {
      "description": "%s termasuk dalam food atau beverage, hanya bisa memiliki value \"food\" dan \"beverage\" saja",
      "type": "string",
      "enum": ["food", "beverage"]
    },
    "ingredient": {
      "description": "buat list ingredient dasar untuk membuat sebuah raw %s dan berikan takaran pasti untuk setiap ingredient nya",
      "type": "array",
      "items": {
        "type": "object",
        "properties": {
          "name": { "type": "string" },
          "measurements": { "type": "string" }
        },
        "required": ["name", "measurements"]
      }
    },
    "nutrition": {
      "description": "jangan ada field tambahan",
      "type": "object",
      "properties": {
        "energy":                  { "description": "energy in KJ/100g",    "type": "number" },
        "saturated_fatty":         { "description": "saturated fatty g/100g", "type": "number" },
        "sugar":                   { "description": "sugar g/100g",          "type": "number" },
        "salt":                    { "description": "salt g/100g",            "type": "number" },
        "protein":                 { "description": "protein g/100g",         "type": "number" },
        "fibres":                  { "description": "fibres g/100g",          "type": "number" },
        "fruit_vegetable_legumes": { "description": "percentage of fruit, vegetables and legumes (%)", "type": "number" }
      },
      "required": ["energy","saturated_fatty","sugar","salt","protein","fibres","fruit_vegetable_legumes"]
    }
  },
  "required": ["type","ingredient","nutrition"]
}`

const NutritionSchema = `{
  "description": "berikan perkiraan nutrition berdasarkan %s",
  "type": "object",
  "properties": {
    "nutrition": {
      "description": "dont add field and answer with english",
      "type": "object",
      "properties": {
        "energy":                  { "description": "energy in KJ/100g",    "type": "number" },
        "saturated_fatty":         { "description": "saturated fatty g/100g", "type": "number" },
        "sugar":                   { "description": "sugar g/100g",          "type": "number" },
        "salt":                    { "description": "salt g/100g",            "type": "number" },
        "protein":                 { "description": "protein g/100g",         "type": "number" },
        "fibres":                  { "description": "fibres g/100g",          "type": "number" },
        "fruit_vegetable_legumes": { "description": "percentage of fruit, vegetables and legumes (%)", "type": "number" }
      },
      "required": ["energy","saturated_fatty","sugar","salt","protein","fibres","fruit_vegetable_legumes"]
    }
  },
  "required": ["nutrition"]
}`
