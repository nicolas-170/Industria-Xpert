package repository

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/nicolas-170/Industria-Xpert/internal/producto/domain/model"
)

type ProductoRepository interface {
	Save(producto *model.Producto) error
	GetByID(idProducto string) (*model.Producto, error)
}

type ProductoRepositoryDB struct {
	db *sql.DB
}

func NewProductoRepository(db *sql.DB) ProductoRepository {
	return &ProductoRepositoryDB{db: db}
}

func (r *ProductoRepositoryDB) Save(producto *model.Producto) error {
	if producto.IdProducto == "" {
		producto.IdProducto = uuid.NewString()
	}
	_, err := r.db.Exec("INSERT INTO producto (id_producto, nombre, talla, color, cantidad, tipo) VALUES (?, ?, ?, ?, ?, ?)",
		producto.IdProducto, producto.Nombre, producto.Talla, producto.Color, producto.Cantidad, producto.Tipo)
	return err
}

func (r *ProductoRepositoryDB) GetByID(idProducto string) (*model.Producto, error) {
	row := r.db.QueryRow("SELECT id_producto, nombre, talla, color, cantidad, tipo FROM producto WHERE id_producto = ?", idProducto)

	producto := &model.Producto{}
	err := row.Scan(&producto.IdProducto, &producto.Nombre, &producto.Talla, &producto.Color, &producto.Cantidad, &producto.Tipo)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return producto, err
}
