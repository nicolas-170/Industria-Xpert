package repository

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"io"

	"gopkg.in/gomail.v2"

	"github.com/nicolas-170/Industria-Xpert/internal/gmail-apartado/domain/model"
)

type GmailApartadoRepository interface {
	SendEmail(emailApartado model.EmailApartado) error
}

type gmailApartadoRepositoryDB struct {
	db            *sql.DB
	emailSend     string
	emailPassword string
}

func NewGmailApartadoRepository(db *sql.DB, emailSend, emailPassword string) GmailApartadoRepository {
	return &gmailApartadoRepositoryDB{
		db:            db,
		emailSend:     emailSend,
		emailPassword: emailPassword,
	}
}

func (g *gmailApartadoRepositoryDB) SendEmail(emailApartado model.EmailApartado) error {
	m := gomail.NewMessage()
	m.SetHeader("From", g.emailSend)
	m.SetHeader("To", emailApartado.EmailTo)
	m.SetHeader("Subject", "Correo enviado de Magic Papers")
	fmt.Println("Mensaje", emailApartado.Mensaje)
	if emailApartado.TipoMensaje == "text/html" {
		m.SetBody("text/html", emailApartado.Mensaje)
	} else {
		m.SetBody("text/plain", emailApartado.Mensaje)
	}

	// Si hay imagen en Base64 → la agregamos como adjunto
	if emailApartado.Imagen.ImagenBase64 != "" {
		data, err := base64.StdEncoding.DecodeString(emailApartado.Imagen.ImagenBase64)
		if err != nil {
			return fmt.Errorf("error decodificando la imagen: %w", err)
		}
		nombre := emailApartado.Imagen.ImagenNombre
		if nombre == "" {
			nombre = "archivo.png"
		}
		m.Attach(nombre, gomail.SetCopyFunc(func(w io.Writer) error {
			_, err := w.Write(data)
			return err
		}))
	}

	connection := gomail.NewDialer("smtp.gmail.com", 587, g.emailSend, g.emailPassword)
	return connection.DialAndSend(m)
}
