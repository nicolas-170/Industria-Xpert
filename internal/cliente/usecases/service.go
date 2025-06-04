package usecases

import (
	"github.com/nicolas-170/Industria-Xpert/internal/cliente/adapter/repository"
	"github.com/nicolas-170/Industria-Xpert/internal/cliente/domain/model"
)

type ClienteService struct {
	repository repository.ClienteRepository
}

func NewClienteService(repository repository.ClienteRepository) *ClienteService {
	return &ClienteService{repository}
}

func (s *ClienteService) Save(cliente *model.Cliente) error {
	return s.repository.Save(cliente)
}

func (s *ClienteService) Obtener(idCliente string) (*model.Cliente, error) {
	return s.repository.GetByID(idCliente)
}
