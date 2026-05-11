package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"

	"github.com/haramurti/-PAW-nutrifood-health-grab-clone/config"
	"github.com/haramurti/-PAW-nutrifood-health-grab-clone/internal/app/ai"
	"github.com/haramurti/-PAW-nutrifood-health-grab-clone/internal/app/product/entity"
	"github.com/haramurti/-PAW-nutrifood-health-grab-clone/internal/app/product/handler"
	"github.com/haramurti/-PAW-nutrifood-health-grab-clone/internal/app/product/repository"
	"github.com/haramurti/-PAW-nutrifood-health-grab-clone/internal/app/product/service"
	"github.com/haramurti/-PAW-nutrifood-health-grab-clone/internal/app/upload"
)

func main() {
	// ── Load .env ─────────────────────────────────────────────────────────────
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, reading from environment")
	}

	// ── Database ──────────────────────────────────────────────────────────────
	config.InitDB()

	// Auto migrate
	if err := config.DB.AutoMigrate(&entity.Product{}); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}

	// ── Gemini ────────────────────────────────────────────────────────────────
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("GEMINI_API_KEY is not set")
	}
	geminiClient, err := ai.NewGeminiClient(apiKey)
	if err != nil {
		log.Fatalf("failed to init gemini: %v", err)
	}

	// ── Fiber app ─────────────────────────────────────────────────────────────
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Use(logger.New())

	// Serve static files (uploaded images)
	publicDir := filepath.Join(".", "internal", "app", "public")
	app.Static("/images", filepath.Join(publicDir, "images"))

	// ── Wire dependencies ─────────────────────────────────────────────────────
	productRepo := repository.NewRepository(config.DB)
	productService := service.NewService(productRepo, geminiClient)
	productHandler := handler.NewHandler(productService)

	uploadDir := filepath.Join(publicDir, "images")
	uploadHandler := upload.NewHandler(uploadDir)

	// ── Routes ────────────────────────────────────────────────────────────────
	api := app.Group("/api")
	productHandler.RegisterRoutes(api.Group("/product"))
	uploadHandler.RegisterRoutes(api.Group("/upload"))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello from health-grab-clone 🚀")
	})

	// ── Start server ──────────────────────────────────────────────────────────
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Printf("Server running on port %s", port)
	if err := app.Listen("0.0.0.0:" + port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
