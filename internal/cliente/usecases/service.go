package usecases

import (
	"github.com/nicolas-170/Industria-Xpert/internal/cliente/adapter/repository"
	"github.com/nicolas-170/Industria-Xpert/internal/cliente/domain/model"
	"golang.org/x/crypto/bcrypt"
)

type ClienteService struct {
	repository repository.ClienteRepository
}

func NewClienteService(repository repository.ClienteRepository) *ClienteService {
	return &ClienteService{repository}
}

func (s *ClienteService) Save(cliente *model.Cliente) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(cliente.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// Pasamos como hash el password
	cliente.Password = string(hashedPassword)
	return s.repository.Save(cliente)
}

func (s *ClienteService) Delete(cliente *model.Cliente) error {
	return s.repository.Delete(cliente)
}

func (s *ClienteService) Obtener(idCliente string) (*model.Cliente, error) {
	return s.repository.GetByID(idCliente)
}
