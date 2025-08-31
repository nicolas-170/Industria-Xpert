package repository

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/nicolas-170/Industria-Xpert/internal/producto/domain/model"
)

type ProductoRepository interface {
	Save(producto *model.Producto) error
	Delete(cliente *model.Producto) error

	GetByID(idProducto string) (*model.Producto, error)
	GetAll() ([]model.Producto, error)
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

func (r *ProductoRepositoryDB) Delete(cliente *model.Producto) error {
	_, err := r.db.Exec("DELETE FROM producto WHERE id_producto = ?", cliente.IdProducto)
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

func (r *ProductoRepositoryDB) GetAll() ([]model.Producto, error) {
	rows, err := r.db.Query("SELECT id_producto, nombre, talla, color, cantidad, tipo, precio FROM producto")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	productos := []model.Producto{}
	for rows.Next() {
		var p model.Producto
		if err = rows.Scan(&p.IdProducto, &p.Nombre, &p.Talla, &p.Color, &p.Cantidad, &p.Tipo, &p.Precio); err != nil {
			rows.Close()
			return nil, err
		}
		productos = append(productos, p)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return productos, nil
}
