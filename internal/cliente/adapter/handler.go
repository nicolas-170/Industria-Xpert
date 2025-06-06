package adapter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nicolas-170/Industria-Xpert/internal/cliente/domain/model"
	"github.com/nicolas-170/Industria-Xpert/internal/cliente/usecases"
)

type ClienteHandler struct {
	service *usecases.ClienteService
}

func NewClienteHandler(service *usecases.ClienteService) *ClienteHandler {
	return &ClienteHandler{service: service}
}

func (h *ClienteHandler) Save(c *fiber.Ctx) error {
	var cliente model.Cliente
	if err := c.BodyParser(&cliente); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Datos inválidos, porque: " + err.Error()})
	}
	if err := h.service.Save(&cliente); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al crear, porque: " + err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(cliente)
}

func (h *ClienteHandler) Delete(c *fiber.Ctx) error {
	var cliente model.Cliente
	if err := c.BodyParser(&cliente); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Datos inválidos, porque: " + err.Error()})
	}
	if err := h.service.Delete(&cliente); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al eliminar, porque: " + err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(cliente)
}

func (h *ClienteHandler) Obtener(c *fiber.Ctx) error {
	idCliente := c.Params("id_cliente")
	cliente, err := h.service.Obtener(idCliente)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al obtener: " + err.Error()})
	}
	if cliente == nil {
		return c.Status(404).JSON(fiber.Map{"error": "No encontrado"})
	}
	return c.JSON(cliente)
}
