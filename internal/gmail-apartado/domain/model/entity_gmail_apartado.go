package model

type EmailApartado struct {
	EmailTo     string `json:"emailTo,omitempty"`
	Codigo      string `json:"codigo ,omitempty"`
	Mensaje     string `json:"mensaje,omitempty"`
	TipoMensaje string `json:"tipoMensaje,omitempty"`
}
