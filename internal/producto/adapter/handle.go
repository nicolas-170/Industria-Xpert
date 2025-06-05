package adapter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nicolas-170/Industria-Xpert/internal/producto/domain/model"
	"github.com/nicolas-170/Industria-Xpert/internal/producto/usecases"
)

type ProductoHandler struct {
	service *usecases.ProductoService
}

func NewProductoHandler(service *usecases.ProductoService) *ProductoHandler {
	return &ProductoHandler{service: service}
}

func (h *ProductoHandler) Save(c *fiber.Ctx) error {
	var producto model.Producto
	if err := c.BodyParser(&producto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Datos inválidos, porque: " + err.Error()})
	}

	if err := h.service.Save(&producto); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al crear, porque: " + err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(producto)
}

func (h *ProductoHandler) Obtener(c *fiber.Ctx) error {
	idProducto := c.Params("id_producto")
	producto, err := h.service.Obtener(idProducto)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al obtener: " + err.Error()})
	}
	if producto == nil {
		return c.Status(404).JSON(fiber.Map{"error": "No encontrado"})
	}
	return c.JSON(producto)
}
