package services

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/nicolas-170/Industria-Xpert/cmd/logger"
	"github.com/nicolas-170/Industria-Xpert/internal/cliente/adapter"
	"github.com/nicolas-170/Industria-Xpert/internal/cliente/adapter/repository"
	"github.com/nicolas-170/Industria-Xpert/internal/cliente/usecases"
)

func Init(baseDeDatos *sql.DB, router fiber.Router) {
	// Servicios de cliente
	initCliente(baseDeDatos, router)
}

func initCliente(baseDeDatos *sql.DB, router fiber.Router) {
	logger.Info("Se inician servicios de cliente...")
	clienteRepo := repository.NewClienteRepository(baseDeDatos)
	clienteService := usecases.NewClienteService(clienteRepo)
	clienteHandler := adapter.NewClienteHandler(clienteService)

	clienteGroup := router.Group("/clientes")

	clienteGroup.Post("/", clienteHandler.Save)
	clienteGroup.Get("/:id_cliente", clienteHandler.Obtener)
}
