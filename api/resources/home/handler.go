package home

import (
	"go-service/config"

	"github.com/gofiber/fiber/v2"
)

func healthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("OK")
}

func versionCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString(config.Server.Version)
}
