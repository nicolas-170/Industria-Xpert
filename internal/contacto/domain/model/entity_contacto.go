package model

type Contacto struct {
	IdContacto string `json:"id_contacto"`
	Whatsapp   string `json:"whatsapp"`
	Correo     string `json:"correo"`
	Direccion  string `json:"direccion"`
}
