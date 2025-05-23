package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nicolas-170/Industria-Xpert/cmd/logger"
	"github.com/nicolas-170/Industria-Xpert/cmd/middleware"
	"github.com/nicolas-170/Industria-Xpert/config"
)

func main() {
	// Iniciamos servidor con fiber
	app := fiber.New()
	// Añadirmos middlewares
	middleware.SetupMiddlewares(app)
	logger.Info("Iniciando servidor...")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	cfg := config.InitConfig()
	// Añadinos en escucha el puerto definido en las ENV
	err := app.Listen(":" + cfg.Port)
	logger.Fatal(err)
}
