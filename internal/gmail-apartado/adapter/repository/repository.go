package repository

import (
	"database/sql"

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
	m.SetHeader("Subject", "Correo enviado desde Magic Papers")
	if emailApartado.TipoMensaje == "text/html" {
		m.SetHeader("text/html", emailApartado.Mensaje)
	} else {
		m.SetHeader("text/plain", emailApartado.Mensaje)
	}

	connection := gomail.NewDialer("smtp.gmail.com", 587, g.emailSend, g.emailPassword)
	return connection.DialAndSend(m)
}
