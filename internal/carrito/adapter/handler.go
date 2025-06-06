package adapter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nicolas-170/Industria-Xpert/internal/carrito/domain/model"
	"github.com/nicolas-170/Industria-Xpert/internal/carrito/usecases"
)

type CarritoHandler struct {
	service *usecases.CarritoService
}

func NewCarritoHandler(service *usecases.CarritoService) *CarritoHandler {
	return &CarritoHandler{service: service}
}

func (h *CarritoHandler) Save(c *fiber.Ctx) error {
	var carrito model.Carrito
	if err := c.BodyParser(&carrito); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Datos inválidos, porque: " + err.Error()})
	}
	if err := h.service.Save(&carrito); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al crear, porque: " + err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(carrito)
}

func (h *CarritoHandler) Delete(c *fiber.Ctx) error {
	var carrito model.Carrito
	if err := c.BodyParser(&carrito); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Datos inválidos, porque: " + err.Error()})
	}
	if err := h.service.Delete(&carrito); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al eliminar: " + err.Error()})
	}
	return c.JSON(carrito)
}

func (h *CarritoHandler) Obtener(c *fiber.Ctx) error {
	idCarrito := c.Params("id_carrito")
	carrito, err := h.service.Obtener(idCarrito)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al obtener: " + err.Error()})
	}
	if carrito == nil {
		return c.Status(404).JSON(fiber.Map{"error": "No encontrado"})
	}
	return c.JSON(carrito)
}
