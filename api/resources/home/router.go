package home

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/")
	api.Get("/health", healthCheck)
	api.Get("/version", versionCheck)
}
