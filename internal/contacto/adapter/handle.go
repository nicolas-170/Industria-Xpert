package adapter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nicolas-170/Industria-Xpert/internal/contacto/domain/model"
	"github.com/nicolas-170/Industria-Xpert/internal/contacto/usecases"
)

type ContactoHandler struct {
	service *usecases.ContactoService
}

func NewContactoHandler(service *usecases.ContactoService) *ContactoHandler {
	return &ContactoHandler{service: service}
}

func (h *ContactoHandler) Save(c *fiber.Ctx) error {
	var contacto model.Contacto
	if err := c.BodyParser(&contacto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Datos inválidos"})
	}

	if err := h.service.Save(&contacto); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al crear, porque: " + err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(contacto)
}

func (h *ContactoHandler) Obtener(c *fiber.Ctx) error {
	idContacto := c.Params("id_contacto")
	contacto, err := h.service.Obtener(idContacto)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error al obtener: " + err.Error()})
	}
	if contacto == nil {
		return c.Status(404).JSON(fiber.Map{"error": "No encontrado"})
	}
	return c.JSON(contacto)
}
