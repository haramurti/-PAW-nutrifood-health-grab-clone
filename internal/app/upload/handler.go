package upload

import (
	"fmt"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/haramurti/-PAW-nutrifood-health-grab-clone/internal/app/helper"
)

type Handler struct {
	uploadDir string // path ke folder public/images
}

func NewHandler(uploadDir string) *Handler {
	return &Handler{uploadDir: uploadDir}
}

func (h *Handler) RegisterRoutes(router fiber.Router) {
	router.Post("/", h.Upload)
}

func (h *Handler) Upload(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return helper.InternalError(c, fmt.Errorf("failed to get file: %w", err))
	}

	filename := uuid.New().String() + ".png"
	savePath := filepath.Join(h.uploadDir, filename)

	if err := c.SaveFile(file, savePath); err != nil {
		return helper.InternalError(c, fmt.Errorf("failed to save file: %w", err))
	}

	return helper.OK(c, fiber.Map{
		"url": "/images/" + filename,
	})
}
