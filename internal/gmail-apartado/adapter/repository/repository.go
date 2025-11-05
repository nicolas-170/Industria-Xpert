package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/nicolas-170/Industria-Xpert/internal/gmail-apartado/domain/model"
)

type GmailApartadoRepository interface {
	SendEmail(emailApartado model.EmailApartado) error
}

type gmailApartadoRepositoryDB struct {
	emailSend     string
	emailPassword string
}

func NewGmailApartadoRepository(emailSend, emailPassword string) GmailApartadoRepository {
	return &gmailApartadoRepositoryDB{
		emailSend:     emailSend,
		emailPassword: emailPassword,
	}
}

type resendRequest struct {
	From    string   `json:"from"`
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	HTML    string   `json:"html,omitempty"`
	Text    string   `json:"text,omitempty"`
}

func (g *gmailApartadoRepositoryDB) SendEmail(emailApartado model.EmailApartado) error {
	apiKey := os.Getenv("RESEND_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("RESEND_API_KEY no configurada")
	}

	reqBody := resendRequest{
		From:    "Magic Papers <onboarding@resend.dev>",
		To:      []string{emailApartado.EmailTo},
		Subject: "Correo enviado de Magic Papers",
	}

	if emailApartado.TipoMensaje == "text/html" {
		reqBody.HTML = emailApartado.Mensaje
	} else {
		reqBody.Text = emailApartado.Mensaje
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "https://api.resend.com/emails", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		bodyString := string(bodyBytes)

		if len(bodyString) > 0 {
			return fmt.Errorf("error Resend API (%d): %s", resp.StatusCode, bodyString)
		}
		return fmt.Errorf("error Resend API (%d): %s", resp.StatusCode, resp.Status)
	}

	return nil
}
