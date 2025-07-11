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

	contactoAdapter "github.com/nicolas-170/Industria-Xpert/internal/contacto/adapter"
	contactoRepo "github.com/nicolas-170/Industria-Xpert/internal/contacto/adapter/repository"
	contactoUsecase "github.com/nicolas-170/Industria-Xpert/internal/contacto/usecases"

	carritoAdapter "github.com/nicolas-170/Industria-Xpert/internal/carrito/adapter"
	carritoRepo "github.com/nicolas-170/Industria-Xpert/internal/carrito/adapter/repository"
	carritoUsecase "github.com/nicolas-170/Industria-Xpert/internal/carrito/usecases"
)

func Init(baseDeDatos *sql.DB, router fiber.Router) {
	// Se inicializan repositorios, de uso por varios casos de uso
	clienteRepo := clienteRepo.NewClienteRepository(baseDeDatos)
	productoRepo := productoRepo.NewProductoRepository(baseDeDatos)

	// Servicios de cliente
	initCliente(router, clienteRepo)
	// Servicios producto
	initProducto(router, productoRepo)
	// Servicios de contacto
	initContacto(baseDeDatos, router)
	// Servicios de carrito
	initCarrito(baseDeDatos, router, clienteRepo, productoRepo)
}

func initCliente(router fiber.Router, clienteRepo clienteRepo.ClienteRepository) {
	logger.Info("Se inician servicios de cliente...")

	clienteService := clienteUsecase.NewClienteService(clienteRepo)
	clienteHandler := clienteAdapter.NewClienteHandler(clienteService)

	// Ruta para procesar las peticiones de clientes
	clienteGroup := router.Group("/clientes")

	clienteGroup.Post("/", clienteHandler.Save)
	clienteGroup.Get("/:id_cliente", clienteHandler.Obtener)
	clienteGroup.Delete("/", clienteHandler.Delete)

	// Autenticarse
	clienteGroup.Post("/autenticarse", clienteHandler.Autenticarse)
}

func initProducto(router fiber.Router, productoRepo productoRepo.ProductoRepository) {
	logger.Info("Se inician servicios de producto...")

	productoService := productoUsecase.NewProductoService(productoRepo)
	productoHandler := productoAdapter.NewProductoHandler(productoService)

	// Ruta para procesar las peticiones de productos
	productoGroup := router.Group("/productos")

	productoGroup.Post("/", productoHandler.Save)
	productoGroup.Get("/:id_producto", productoHandler.Obtener)
	productoGroup.Delete("/", productoHandler.Delete)
}

func initContacto(baseDeDatos *sql.DB, router fiber.Router) {
	logger.Info("Se inician servicios de contacto...")
	contactoRepo := contactoRepo.NewContactoRepository(baseDeDatos)
	contactoService := contactoUsecase.NewContactoService(contactoRepo)
	contactoHandler := contactoAdapter.NewContactoHandler(contactoService)

	// Ruta para procesar las peticiones de contactos
	contactoGroup := router.Group("/contactos")

	contactoGroup.Post("/", contactoHandler.Save)
	contactoGroup.Get("/:id_contacto", contactoHandler.Obtener)
	contactoGroup.Delete("/", contactoHandler.Delete)
}

func initCarrito(baseDeDatos *sql.DB, router fiber.Router, clienteRepo clienteRepo.ClienteRepository, productoRepo productoRepo.ProductoRepository) {
	logger.Info("Se inician servicios de carrito...")
	carritoRepo := carritoRepo.NewCarritoRepository(baseDeDatos)
	carritoService := carritoUsecase.NewCarritoService(carritoRepo, productoRepo, clienteRepo)
	carritoHandler := carritoAdapter.NewCarritoHandler(carritoService)

	// Ruta para procesar las peticiones de carritos
	carritoGroup := router.Group("/carritos")

	carritoGroup.Post("/", carritoHandler.Save)
	carritoGroup.Get("/:id_carrito", carritoHandler.Obtener)
	carritoGroup.Delete("/", carritoHandler.Delete)
}
