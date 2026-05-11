package entity

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

// ── JSON field types ──────────────────────────────────────────────────────────

type Nutrition struct {
	Energy                float64 `json:"energy"`
	SaturatedFatty        float64 `json:"saturated_fatty"`
	Sugar                 float64 `json:"sugar"`
	Salt                  float64 `json:"salt"`
	Protein               float64 `json:"protein"`
	Fibres                float64 `json:"fibres"`
	FruitVegetableLegumes float64 `json:"fruit_vegetable_legumes"`
}

type Ingredient struct {
	Name         string `json:"name"`
	Measurements string `json:"measurements"`
}

// JSONB wrappers ── agar GORM bisa scan/value ke PostgreSQL JSONB

type NutritionJSON Nutrition

func (n NutritionJSON) Value() (driver.Value, error) {
	return json.Marshal(n)
}
func (n *NutritionJSON) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, n)
}

type IngredientsJSON []Ingredient

func (i IngredientsJSON) Value() (driver.Value, error) {
	return json.Marshal(i)
}
func (i *IngredientsJSON) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, i)
}

// ── Main entity ───────────────────────────────────────────────────────────────

type Product struct {
	ID          string          `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Name        string          `gorm:"not null" json:"name"`
	Type        string          `gorm:"not null" json:"type"` // "food" | "beverage"
	Price       int             `gorm:"not null" json:"price"`
	Image       string          `gorm:"not null" json:"image"`
	Description string          `gorm:"not null" json:"description"`
	Score       int             `gorm:"not null" json:"score"`
	Grade       string          `gorm:"not null;size:1" json:"grade"`
	Nutrition   NutritionJSON   `gorm:"type:jsonb;not null" json:"nutrition"`
	Ingredient  IngredientsJSON `gorm:"type:jsonb;not null" json:"ingredient"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

// ── DTOs ─────────────────────────────────────────────────────────────────────

type CreateProductRequest struct {
	Name        string       `json:"name"`
	Type        string       `json:"type"`
	Price       int          `json:"price"`
	Image       string       `json:"image"`
	Description string       `json:"description"`
	Score       int          `json:"score"`
	Grade       string       `json:"grade"`
	Nutrition   Nutrition    `json:"nutrition"`
	Ingredients []Ingredient `json:"ingredients"`
}

type UpdateProductRequest struct {
	Name        string       `json:"name"`
	Type        string       `json:"type"`
	Price       int          `json:"price"`
	Image       string       `json:"image"`
	Description string       `json:"description"`
	Score       int          `json:"score"`
	Grade       string       `json:"grade"`
	Nutrition   Nutrition    `json:"nutrition"`
	Ingredient  []Ingredient `json:"ingredient"`
}

type GenerateByNameRequest struct {
	Name string `json:"name"`
}

type GenerateByIngredientsRequest struct {
	Type        string       `json:"type"`
	Ingredients []Ingredient `json:"ingredients"`
}

// ── Response shapes ───────────────────────────────────────────────────────────

type GradeDetail struct {
	Grade       string `json:"grade"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Color       string `json:"color"`
	Image       string `json:"image"`
	URL         string `json:"url"`
}

type ProductResponse struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	Type        string          `json:"type"`
	Price       int             `json:"price"`
	Image       string          `json:"image"`
	Description string          `json:"description"`
	Score       int             `json:"score"`
	Grade       string          `json:"grade"`
	GradeDetail GradeDetail     `json:"grade_detail"`
	Nutrition   NutritionJSON   `json:"nutrition"`
	Ingredients IngredientsJSON `json:"ingredients"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

type GenerateResponse struct {
	Score       int          `json:"score"`
	Type        string       `json:"type"`
	Grade       string       `json:"grade"`
	GradeDetail GradeDetail  `json:"grade_detail"`
	Nutrition   Nutrition    `json:"nutrition"`
	Ingredients []Ingredient `json:"ingredients"`
}
