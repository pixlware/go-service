package api

import (
	"go-api-bp/api/resources/home"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	home.SetupRoutes(app)
}
