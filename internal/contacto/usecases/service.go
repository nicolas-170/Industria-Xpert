package usecases

import (
	"github.com/nicolas-170/Industria-Xpert/internal/contacto/adapter/repository"
	"github.com/nicolas-170/Industria-Xpert/internal/contacto/domain/model"
)

type ContactoService struct {
	repository repository.ContactoRepository
}

func NewContactoService(repository repository.ContactoRepository) *ContactoService {
	return &ContactoService{repository}
}

func (s *ContactoService) Save(contacto *model.Contacto) error {
	return s.repository.Save(contacto)
}

func (s *ContactoService) Delete(contacto *model.Contacto) error {
	return s.repository.Delete(contacto)
}

func (s *ContactoService) Obtener(idContacto string) (*model.Contacto, error) {
	return s.repository.GetByID(idContacto)
}
