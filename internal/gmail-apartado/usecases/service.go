package usecases

import (
	"github.com/nicolas-170/Industria-Xpert/internal/gmail-apartado/adapter/repository"
	"github.com/nicolas-170/Industria-Xpert/internal/gmail-apartado/domain/model"
)

type GmailApartadoService struct {
	repository repository.GmailApartadoRepository
}

func NewGmailApartadoService(repository repository.GmailApartadoRepository) *GmailApartadoService {
	return &GmailApartadoService{repository}
}

func (s *GmailApartadoService) SendEmail(emailApartado model.EmailApartado) error {
	return s.repository.SendEmail(emailApartado)
}
