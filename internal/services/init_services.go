package services

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/nicolas-170/Industria-Xpert/cmd/logger"
	clienteAdapter "github.com/nicolas-170/Industria-Xpert/internal/cliente/adapter"
	clienteRepo "github.com/nicolas-170/Industria-Xpert/internal/cliente/adapter/repository"
	clienteUsecase "github.com/nicolas-170/Industria-Xpert/internal/cliente/usecases"

	productoAdapter "github.com/nicolas-170/Industria-Xpert/internal/producto/adapter"
	productoRepo "github.com/nicolas-170/Industria-Xpert/internal/producto/adapter/repository"
	productoUsecase "github.com/nicolas-170/Industria-Xpert/internal/producto/usecases"
)

func Init(baseDeDatos *sql.DB, router fiber.Router) {
	// Servicios de cliente
	initCliente(baseDeDatos, router)
	// Servicios producto
	initProducto(baseDeDatos, router)
}

func initCliente(baseDeDatos *sql.DB, router fiber.Router) {
	logger.Info("Se inician servicios de cliente...")
	clienteRepo := clienteRepo.NewClienteRepository(baseDeDatos)
	clienteService := clienteUsecase.NewClienteService(clienteRepo)
	clienteHandler := clienteAdapter.NewClienteHandler(clienteService)

	clienteGroup := router.Group("/clientes")

	clienteGroup.Post("/", clienteHandler.Save)
	clienteGroup.Get("/:id_cliente", clienteHandler.Obtener)
}

func initProducto(baseDeDatos *sql.DB, router fiber.Router) {
	logger.Info("Se inician servicios de producto...")
	productoRepo := productoRepo.NewProductoRepository(baseDeDatos)
	productoService := productoUsecase.NewProductoService(productoRepo)
	productoHandler := productoAdapter.NewProductoHandler(productoService)

	productoGroup := router.Group("/productos")

	productoGroup.Post("/", productoHandler.Save)
	productoGroup.Get("/:id_producto", productoHandler.Obtener)
}
