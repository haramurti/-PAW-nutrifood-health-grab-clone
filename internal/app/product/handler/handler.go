package handler

import (
	"github.com/haramurti/-PAW-nutrifood-health-grab-clone/internal/app/helper"
	"github.com/haramurti/-PAW-nutrifood-health-grab-clone/internal/app/product/entity"
	"github.com/haramurti/-PAW-nutrifood-health-grab-clone/internal/app/product/service"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service service.Service
}

func NewHandler(service service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(router fiber.Router) {
	router.Get("/", h.GetAll)
	router.Get("/:id", h.GetByID)
	router.Post("/", h.Create)
	router.Put("/:id", h.Update)
	router.Post("/generate/ingredient-nutriscore", h.GenerateByName)
	router.Post("/generate/nutriscore", h.GenerateByIngredients)
}

func (h *Handler) GetAll(c *fiber.Ctx) error {
	search := c.Query("search")
	products, err := h.service.GetAll(search)
	if err != nil {
		return helper.InternalError(c, err)
	}
	return helper.OK(c, products)
}

func (h *Handler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	product, err := h.service.GetByID(id)
	if err != nil {
		return helper.InternalError(c, err)
	}
	if product == nil {
		return helper.NotFound(c, "Product Not Found")
	}
	return helper.OK(c, product)
}

func (h *Handler) Create(c *fiber.Ctx) error {
	var req entity.CreateProductRequest
	if err := c.BodyParser(&req); err != nil {
		return helper.InternalError(c, err)
	}
	id, err := h.service.Create(req)
	if err != nil {
		return helper.InternalError(c, err)
	}
	return helper.Created(c, fiber.Map{"id": id})
}

func (h *Handler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var req entity.UpdateProductRequest
	if err := c.BodyParser(&req); err != nil {
		return helper.InternalError(c, err)
	}
	updatedID, err := h.service.Update(id, req)
	if err != nil {
		return helper.InternalError(c, err)
	}
	return helper.Created(c, fiber.Map{"id": updatedID})
}

func (h *Handler) GenerateByName(c *fiber.Ctx) error {
	var req entity.GenerateByNameRequest
	if err := c.BodyParser(&req); err != nil {
		return helper.InternalError(c, err)
	}
	result, err := h.service.GenerateByName(c.Context(), req.Name)
	if err != nil {
		return helper.InternalError(c, err)
	}
	return helper.OK(c, result)
}

func (h *Handler) GenerateByIngredients(c *fiber.Ctx) error {
	var req entity.GenerateByIngredientsRequest
	if err := c.BodyParser(&req); err != nil {
		return helper.InternalError(c, err)
	}
	result, err := h.service.GenerateByIngredients(c.Context(), req.Type, req.Ingredients)
	if err != nil {
		return helper.InternalError(c, err)
	}
	return helper.OK(c, result)
}
