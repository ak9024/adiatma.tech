package server

import (
	"os"

	"github.com/ak9024/adiatma.tech/api/docs"
	"github.com/ak9024/adiatma.tech/api/internal/hadith"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/swagger"
)

var (
	app *fiber.App
)

func Router() *fiber.App {
	app = fiber.New()

	docs.SwaggerInfo.BasePath = os.Getenv("BASE_PATH")
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	// GET /metrics
	app.Get("/metrics", monitor.New())
	// GET /swagger/*
	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL: "../swagger/doc.json",
	}))

	h := hadith.New()

	// GET /hadith/:author
	app.Get("/hadith/:author", h.GetHadith)
	// GET /hadith/list/authors
	app.Get("/hadith/list/authors", h.GetAuthor)

	return app
}
