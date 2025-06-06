package model

type Contacto struct {
	IdContacto string `json:"id_contacto,omitempty"`
	Whatsapp   string `json:"whatsapp,omitempty"`
	Correo     string `json:"correo,omitempty"`
	Direccion  string `json:"direccion,omitempty"`
}
