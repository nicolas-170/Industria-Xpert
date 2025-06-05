package repository

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/nicolas-170/Industria-Xpert/internal/cliente/domain/model"
)

type ClienteRepository interface {
	Save(cliente *model.Cliente) error
	GetByID(idCliente string) (*model.Cliente, error)
}

type clienteRepositoryDB struct {
	db *sql.DB
}

func NewClienteRepository(db *sql.DB) ClienteRepository {
	return &clienteRepositoryDB{db: db}
}

func (r *clienteRepositoryDB) Save(cliente *model.Cliente) error {
	if cliente.IdCliente == "" {
		cliente.IdCliente = uuid.NewString()
	}
	_, err := r.db.Exec("INSERT INTO cliente (id_cliente, celular, nombre, correo, identificacion, direccion) VALUES (?, ?, ?, ?, ?, ?)",
		cliente.IdCliente, cliente.Celular, cliente.Nombre, cliente.Correo, cliente.Identificacion, cliente.Direccion)
	return err
}

func (r *clienteRepositoryDB) GetByID(idCliente string) (*model.Cliente, error) {
	row := r.db.QueryRow("SELECT id_cliente, celular, nombre, correo, identificacion, direccion FROM cliente WHERE id_cliente = ?", idCliente)

	cliente := &model.Cliente{}
	err := row.Scan(&cliente.IdCliente, &cliente.Celular, &cliente.Nombre, &cliente.Correo, &cliente.Identificacion, &cliente.Direccion)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return cliente, err
}
