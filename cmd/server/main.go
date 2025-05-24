package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nicolas-170/Industria-Xpert/cmd/logger"
	"github.com/nicolas-170/Industria-Xpert/cmd/middleware"
	"github.com/nicolas-170/Industria-Xpert/config"
)

func main() {
	// Iniciamos la configuraciob
	cfg := config.InitConfig()

	// Iniciamos servidor con fiber
	appFiber := fiber.New(fiber.Config{
		AppName: cfg.AppName,
	})

	// Añadirmos middlewares
	middleware.SetupMiddlewares(appFiber)
	logger.Info("Iniciando servidor...")

	// Agrupammos las rutas con el nombre de la aplicacion en el inicio de cada ruta
	router := appFiber.Group("/" + cfg.AppName)

	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Inicio de la aplicacion " + cfg.AppName)
	})

	// Añadinos en escucha el puerto definido en las ENV
	logger.Fatal(appFiber.Listen(":" + cfg.Port))
}
