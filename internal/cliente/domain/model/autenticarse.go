package model

type Autenticarse struct {
	Correo   string `json:"correo,omitempty"`
	Password string `json:"password,omitempty"`
}
