package api

import (
	"go-service/api/resources/home"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	home.SetupRoutes(app)
}
