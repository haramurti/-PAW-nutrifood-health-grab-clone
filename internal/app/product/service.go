package product

import (
	"context"
	"errors"

	"github.com/haramurti/-PAW-nutrifood-health-grab-clone/internal/app/ai"
	"github.com/haramurti/-PAW-nutrifood-health-grab-clone/internal/app/helper"
	"github.com/haramurti/-PAW-nutrifood-health-grab-clone/internal/app/helper/grade"
	"github.com/haramurti/-PAW-nutrifood-health-grab-clone/internal/app/product/contract"
	"github.com/haramurti/-PAW-nutrifood-health-grab-clone/internal/app/product/entity"
	"gorm.io/gorm"
)

type Service interface {
	GetAll(search string) ([]entity.ProductResponse, error)
	GetByID(id string) (*entity.ProductResponse, error)
	Create(req entity.CreateProductRequest) (string, error)
	Update(id string, req entity.UpdateProductRequest) (string, error)
	GenerateByName(ctx context.Context, name string) (*entity.GenerateResponse, error)
	GenerateByIngredients(ctx context.Context, foodType string, ingredients []entity.Ingredient) (*entity.GenerateResponse, error)
}

type serviceImpl struct {
	repo   contract.Repository
	gemini *ai.GeminiClient
}

func NewService(repo contract.Repository, gemini *ai.GeminiClient) Service {
	return &serviceImpl{repo: repo, gemini: gemini}
}

// ── helpers ───────────────────────────────────────────────────────────────────

func toResponse(p entity.Product) entity.ProductResponse {
	return entity.ProductResponse{
		ID:          p.ID,
		Name:        p.Name,
		Type:        p.Type,
		Price:       p.Price,
		Image:       p.Image,
		Description: p.Description,
		Score:       p.Score,
		Grade:       p.Grade,
		GradeDetail: helper.GetGrade(p.Grade, p.Type),
		Nutrition:   p.Nutrition,
		Ingredients: p.Ingredient,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

// ── implementations ───────────────────────────────────────────────────────────

func (s *serviceImpl) GetAll(search string) ([]entity.ProductResponse, error) {
	var products []entity.Product
	var err error

	if search != "" {
		products, err = s.repo.FindByName(search)
	} else {
		products, err = s.repo.FindAll()
	}
	if err != nil {
		return nil, err
	}

	result := make([]entity.ProductResponse, len(products))
	for i, p := range products {
		result[i] = toResponse(p)
	}
	return result, nil
}

func (s *serviceImpl) GetByID(id string) (*entity.ProductResponse, error) {
	p, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // caller handles 404
		}
		return nil, err
	}
	resp := toResponse(*p)
	return &resp, nil
}

func (s *serviceImpl) Create(req entity.CreateProductRequest) (string, error) {
	p := &entity.Product{
		Name:        req.Name,
		Type:        req.Type,
		Price:       req.Price,
		Image:       req.Image,
		Description: req.Description,
		Score:       req.Score,
		Grade:       req.Grade,
		Nutrition:   entity.NutritionJSON(req.Nutrition),
		Ingredient:  entity.IngredientsJSON(req.Ingredients),
	}
	return s.repo.Create(p)
}

func (s *serviceImpl) Update(id string, req entity.UpdateProductRequest) (string, error) {
	p := &entity.Product{
		Name:        req.Name,
		Type:        req.Type,
		Price:       req.Price,
		Image:       req.Image,
		Description: req.Description,
		Score:       req.Score,
		Grade:       req.Grade,
		Nutrition:   entity.NutritionJSON(req.Nutrition),
		Ingredient:  entity.IngredientsJSON(req.Ingredient),
	}
	return s.repo.Update(id, p)
}

func (s *serviceImpl) GenerateByName(ctx context.Context, name string) (*entity.GenerateResponse, error) {
	generated, err := s.gemini.GenerateIngredientNutrition(ctx, name)
	if err != nil {
		return nil, err
	}

	gradeResult, err := grade.GenerateContent(generated.Type, generated.Nutrition)
	if err != nil {
		return nil, err
	}

	return &entity.GenerateResponse{
		Score:       gradeResult.Score,
		Type:        generated.Type,
		Grade:       gradeResult.Grade,
		GradeDetail: helper.GetGrade(gradeResult.Grade, generated.Type),
		Nutrition:   generated.Nutrition,
		Ingredients: generated.Ingredient,
	}, nil
}

func (s *serviceImpl) GenerateByIngredients(ctx context.Context, foodType string, ingredients []entity.Ingredient) (*entity.GenerateResponse, error) {
	nutritionResp, err := s.gemini.GenerateNutrition(ctx, ingredients)
	if err != nil {
		return nil, err
	}

	gradeResult, err := grade.GenerateContent(foodType, nutritionResp.Nutrition)
	if err != nil {
		return nil, err
	}

	return &entity.GenerateResponse{
		Score:       gradeResult.Score,
		Type:        foodType,
		Grade:       gradeResult.Grade,
		GradeDetail: helper.GetGrade(gradeResult.Grade, foodType),
		Nutrition:   nutritionResp.Nutrition,
		Ingredients: ingredients,
	}, nil
}
