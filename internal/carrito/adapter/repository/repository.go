package repository

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/nicolas-170/Industria-Xpert/internal/carrito/domain/model"
)

type CarritoRepository interface {
	Save(carrito *model.Carrito) error
	Delete(carrito *model.Carrito) error

	GetByID(idCarrito string) (*model.Carrito, error)
}

type carritoRepositoryDB struct {
	db *sql.DB
}

func NewCarritoRepository(db *sql.DB) CarritoRepository {
	return &carritoRepositoryDB{db: db}
}

func (r *carritoRepositoryDB) Save(carrito *model.Carrito) error {
	if carrito.IdCarrito == "" {
		carrito.IdCarrito = uuid.NewString()
	}
	_, err := r.db.Exec("INSERT INTO carrito (id_carrito, id_producto, id_cliente, fecha) VALUES (?, ?, ?, ?)",
		carrito.IdCarrito, carrito.IdProducto, carrito.IdCliente, carrito.Fecha)
	return err
}

func (r *carritoRepositoryDB) Delete(carrito *model.Carrito) error {
	_, err := r.db.Exec("DELETE FROM carrito WHERE id_carrito = ?", carrito.IdCarrito)
	return err
}

func (r *carritoRepositoryDB) GetByID(idCarrito string) (*model.Carrito, error) {
	row := r.db.QueryRow("SELECT id_carrito, id_producto, id_cliente, fecha FROM carrito WHERE id_carrito = ?", idCarrito)

	carrito := &model.Carrito{}
	err := row.Scan(&carrito.IdCarrito, &carrito.IdProducto, &carrito.IdCliente, &carrito.Fecha)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return carrito, err
}
