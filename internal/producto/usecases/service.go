package usecases

import (
	"github.com/nicolas-170/Industria-Xpert/internal/producto/adapter/repository"
	"github.com/nicolas-170/Industria-Xpert/internal/producto/domain/model"
)

type ProductoService struct {
	repository repository.ProductoRepository
}

func NewProductoService(repository repository.ProductoRepository) *ProductoService {
	return &ProductoService{repository}
}

func (s *ProductoService) Save(cliente *model.Producto) error {
	return s.repository.Save(cliente)
}

func (s *ProductoService) Obtener(idProducto string) (*model.Producto, error) {
	return s.repository.GetByID(idProducto)
}
