package adapter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nicolas-170/Industria-Xpert/internal/gmail-apartado/domain/model"
	"github.com/nicolas-170/Industria-Xpert/internal/gmail-apartado/usecases"
)

type GmailApartadoHandler struct {
	service *usecases.GmailApartadoService
}

func NewGmailApartadoHandler(service *usecases.GmailApartadoService) *GmailApartadoHandler {
	return &GmailApartadoHandler{service: service}
}

func (h *GmailApartadoHandler) SendEmail(c *fiber.Ctx) error {
	var emailApartado model.EmailApartado
	if err := c.BodyParser(&emailApartado); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Datos inválidos, porque: " + err.Error()})
	}

	if err := h.service.SendEmail(emailApartado); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al crear, porque: " + err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(emailApartado)
}
