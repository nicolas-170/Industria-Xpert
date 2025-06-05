package usecases

import (
	"errors"

	"github.com/nicolas-170/Industria-Xpert/internal/carrito/adapter/repository"
	"github.com/nicolas-170/Industria-Xpert/internal/carrito/domain/model"
	clienteRepo "github.com/nicolas-170/Industria-Xpert/internal/cliente/adapter/repository"
	productoRepo "github.com/nicolas-170/Industria-Xpert/internal/producto/adapter/repository"
)

type CarritoService struct {
	carritoRepository  repository.CarritoRepository
	productoRepository productoRepo.ProductoRepository
	clienteRepository  clienteRepo.ClienteRepository
}

func NewCarritoService(
	carritoRepository repository.CarritoRepository,
	productoRepository productoRepo.ProductoRepository,
	clienteRepository clienteRepo.ClienteRepository,
) *CarritoService {
	return &CarritoService{carritoRepository, productoRepository, clienteRepository}
}

func (s *CarritoService) Save(carrito *model.Carrito) error {
	cliente, err := s.clienteRepository.GetByID(carrito.IdCliente)
	if err != nil {
		return err
	}
	if cliente == nil {
		return errors.New("el id del cliente: " + carrito.IdCliente + " no se encuentra registrado.")
	}
	producto, err := s.productoRepository.GetByID(carrito.IdProducto)
	if err != nil {
		return err
	}
	if producto == nil {
		return errors.New("el id del producto: " + carrito.IdProducto + " no se encuentra registrado.")
	}
	return s.carritoRepository.Save(carrito)
}

func (s *CarritoService) Obtener(idCarrito string) (*model.Carrito, error) {
	return s.carritoRepository.GetByID(idCarrito)
}
