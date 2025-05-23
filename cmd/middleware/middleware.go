package middleware

import (
	"github.com/gofiber/fiber/v2"
	middlewareLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/nicolas-170/Industria-Xpert/cmd/logger"
)

// zerologWriter adapta zerolog para ser usado como io.Writer
type zerologWriter struct {
}

func (w *zerologWriter) Write(p []byte) (n int, err error) {
	if err := logger.Info(string(p)); err != nil {
		return 0, err
	}
	return len(p), nil
}

// SetupMiddlewares añade los middlewares a Fiber
func SetupMiddlewares(app *fiber.App) {
	// Crea nuestro logger
	logger.NewLogger()

	// Añade el middleware de logger de Fiber
	app.Use(middlewareLogger.New(middlewareLogger.Config{
		Output:     &zerologWriter{},
		TimeFormat: "2006-01-02 15:04:05",
		Format:     "${time} | ${status} | ${latency} | ${ip} | ${method} | ${path} | ${error}\n",
	}))
}
