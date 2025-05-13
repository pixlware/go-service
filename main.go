package main

import (
	"go-api-bp/api"
	"go-api-bp/config"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

const Version = "1.0.0"

func main() {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: config.Server.Env != "local" && config.Server.Env != "default",
	})

	app.Use(cors.New(cors.Config{
		AllowMethods: config.Server.CorsMethods,
		AllowOrigins: config.Server.CorsOrigins,
	}))

	api.RegisterRoutes(app)

	log.Println("Running '" + config.Server.Env + "' environment on port: " + config.Server.Port)
	app.Listen(":" + config.Server.Port)
}
