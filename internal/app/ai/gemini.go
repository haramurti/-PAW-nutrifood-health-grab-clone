package ai

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/haramurti/-PAW-nutrifood-health-grab-clone/internal/app/product/entity"

	"google.golang.org/genai"
)

type GeminiClient struct {
	client *genai.Client
}

func NewGeminiClient(apiKey string) (*GeminiClient, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create gemini client: %w", err)
	}
	return &GeminiClient{client: client}, nil
}

// ── Response shapes dari Gemini ───────────────────────────────────────────────

type IngredientNutritionResponse struct {
	Type       string              `json:"type"`
	Ingredient []entity.Ingredient `json:"ingredient"`
	Nutrition  entity.Nutrition    `json:"nutrition"`
}

type NutritionResponse struct {
	Nutrition entity.Nutrition `json:"nutrition"`
}

// ── Helpers ───────────────────────────────────────────────────────────────────

func cleanJSON(raw string) string {
	s := strings.TrimSpace(raw)
	s = strings.TrimPrefix(s, "```json")
	s = strings.TrimPrefix(s, "```")
	s = strings.TrimSuffix(s, "```")
	return strings.TrimSpace(s)
}

func (g *GeminiClient) sendPrompt(ctx context.Context, prompt string) (string, error) {
	result, err := g.client.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash",
		genai.Text(prompt),
		&genai.GenerateContentConfig{
			Temperature: genai.Ptr[float32](0),
			TopP:        genai.Ptr[float32](0.1),
			TopK:        genai.Ptr[float32](1),
		},
	)
	if err != nil {
		return "", fmt.Errorf("gemini generate error: %w", err)
	}
	if len(result.Candidates) == 0 || len(result.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("gemini returned empty response")
	}
	return result.Candidates[0].Content.Parts[0].Text, nil
}

// ── Public methods ────────────────────────────────────────────────────────────

func (g *GeminiClient) GenerateIngredientNutrition(ctx context.Context, name string) (*IngredientNutritionResponse, error) {
	schema := fmt.Sprintf(IngredientNutritionSchema, name, name, name)
	prompt := fmt.Sprintf("Follow JSON schema.<JSONSchema>%s</JSONSchema>", schema)

	raw, err := g.sendPrompt(ctx, prompt)
	if err != nil {
		return nil, err
	}

	var result IngredientNutritionResponse
	if err := json.Unmarshal([]byte(cleanJSON(raw)), &result); err != nil {
		return nil, fmt.Errorf("failed to parse gemini response: %w", err)
	}
	return &result, nil
}

func (g *GeminiClient) GenerateNutrition(ctx context.Context, ingredients []entity.Ingredient) (*NutritionResponse, error) {
	// format ingredients jadi "name measurements, name measurements, ..."
	parts := make([]string, len(ingredients))
	for i, ing := range ingredients {
		parts[i] = fmt.Sprintf("%s %s", ing.Name, ing.Measurements)
	}
	ingredientsStr := strings.Join(parts, ", ")

	schema := fmt.Sprintf(NutritionSchema, ingredientsStr)
	prompt := fmt.Sprintf("Follow JSON schema.<JSONSchema>%s</JSONSchema>", schema)

	raw, err := g.sendPrompt(ctx, prompt)
	if err != nil {
		return nil, err
	}

	var result NutritionResponse
	if err := json.Unmarshal([]byte(cleanJSON(raw)), &result); err != nil {
		return nil, fmt.Errorf("failed to parse gemini response: %w", err)
	}
	return &result, nil
}
