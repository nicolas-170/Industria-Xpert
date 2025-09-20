package repository

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/nicolas-170/Industria-Xpert/internal/contacto/domain/model"
)

type ContactoRepository interface {
	Save(contacto *model.Contacto) error
	Delete(carrito *model.Contacto) error

	GetByID(idContacto string) (*model.Contacto, error)
}

type contactoRepositoryDB struct {
	db *sql.DB
}

func NewContactoRepository(db *sql.DB) ContactoRepository {
	return &contactoRepositoryDB{db: db}
}

func (r *contactoRepositoryDB) Save(contacto *model.Contacto) error {
	if contacto.IdContacto == "" {
		contacto.IdContacto = uuid.NewString()
	}
	_, err := r.db.Exec("INSERT INTO contacto (id_contacto, whatsapp, correo, direccion) VALUES (?, ?, ?, ?)",
		contacto.IdContacto, contacto.Whatsapp, contacto.Correo, contacto.Direccion)
	return err
}
func (r *contactoRepositoryDB) Delete(carrito *model.Contacto) error {
	_, err := r.db.Exec("DELETE FROM contacto WHERE id_contacto = ?", carrito.IdContacto)
	return err
}

func (r *contactoRepositoryDB) GetByID(idContacto string) (*model.Contacto, error) {
	row := r.db.QueryRow("SELECT id_contacto, whatsapp, correo, direccion FROM contacto WHERE id_contacto = ?", idContacto)

	contacto := &model.Contacto{}
	err := row.Scan(&contacto.IdContacto, &contacto.Whatsapp, &contacto.Correo, &contacto.Direccion)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return contacto, err
}
