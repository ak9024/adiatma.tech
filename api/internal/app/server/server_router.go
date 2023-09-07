package server

import (
	"github.com/ak9024/adiatma.tech/api/internal/hadith"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

var (
	app *fiber.App
)

func Router() *fiber.App {
	app = fiber.New()

	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	app.Get("/metrics", monitor.New())

	h := hadith.New()
	app.Get("/hadith/:hadith", h.GetHadith)

	return app
}
