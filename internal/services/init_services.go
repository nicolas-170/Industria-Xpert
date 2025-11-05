package services

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/nicolas-170/Industria-Xpert/cmd/logger"
	"github.com/nicolas-170/Industria-Xpert/config"
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

	gmailApartadoAdapter "github.com/nicolas-170/Industria-Xpert/internal/gmail-apartado/adapter"
	gmailApartadoRepo "github.com/nicolas-170/Industria-Xpert/internal/gmail-apartado/adapter/repository"
	gmailApartadoUsecase "github.com/nicolas-170/Industria-Xpert/internal/gmail-apartado/usecases"
)

func Init(baseDeDatos *sql.DB, router fiber.Router, cfg *config.Config) {
	// Se inicializan repositorios, de uso por varios casos de uso
	clienteRepo := clienteRepo.NewClienteRepository(baseDeDatos)
	contactoRepo := contactoRepo.NewContactoRepository(baseDeDatos)
	productoRepo := productoRepo.NewProductoRepository(baseDeDatos)
	carritoRepo := carritoRepo.NewCarritoRepository(baseDeDatos)
	gmailApartadoRepo := gmailApartadoRepo.NewGmailApartadoRepository(cfg.Email.EmailSend, cfg.Email.EmailPassword)

	// Servicios de cliente
	initCliente(router, clienteRepo)
	// Servicios producto
	initProducto(router, productoRepo)
	// Servicios de contacto
	initContacto(router, contactoRepo)
	// Servicios de carrito
	initCarrito(router, carritoRepo, clienteRepo, productoRepo)

	// Servicios de gmail apartado
	initGmailApartado(router, gmailApartadoRepo)
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

	productoGroup.Get("/obtener-todos", productoHandler.ObtenerTodos)
	productoGroup.Get("/:id_producto", productoHandler.Obtener)

	productoGroup.Post("/", productoHandler.Save)
	productoGroup.Delete("/", productoHandler.Delete)
}

func initContacto(router fiber.Router, contactoRepo contactoRepo.ContactoRepository) {
	logger.Info("Se inician servicios de contacto...")
	contactoService := contactoUsecase.NewContactoService(contactoRepo)
	contactoHandler := contactoAdapter.NewContactoHandler(contactoService)

	// Ruta para procesar las peticiones de contactos
	contactoGroup := router.Group("/contactos")

	contactoGroup.Post("/", contactoHandler.Save)
	contactoGroup.Get("/:id_contacto", contactoHandler.Obtener)
	contactoGroup.Delete("/", contactoHandler.Delete)
}

func initCarrito(router fiber.Router, carritoRepo carritoRepo.CarritoRepository, clienteRepo clienteRepo.ClienteRepository, productoRepo productoRepo.ProductoRepository) {
	logger.Info("Se inician servicios de carrito...")

	carritoService := carritoUsecase.NewCarritoService(carritoRepo, productoRepo, clienteRepo)
	carritoHandler := carritoAdapter.NewCarritoHandler(carritoService)

	// Ruta para procesar las peticiones de carritos
	carritoGroup := router.Group("/carritos")

	carritoGroup.Post("/", carritoHandler.Save)
	carritoGroup.Get("/:id_carrito", carritoHandler.Obtener)
	carritoGroup.Delete("/", carritoHandler.Delete)
}

func initGmailApartado(router fiber.Router, gmailApartadoRepo gmailApartadoRepo.GmailApartadoRepository) {
	logger.Info("Se inicia servicios de gmail apartado...")
	gmailApartadoService := gmailApartadoUsecase.NewGmailApartadoService(gmailApartadoRepo)
	gmailApartadoHandler := gmailApartadoAdapter.NewGmailApartadoHandler(gmailApartadoService)

	gmailApartadoGroup := router.Group("/gmail-apartado")

	gmailApartadoGroup.Post("/", gmailApartadoHandler.SendEmail)
}
