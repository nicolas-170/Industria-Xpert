package model

type Cliente struct {
	IdCliente      string `json:"id_cliente,omitempty"`
	Celular        string `json:"celular,omitempty"`
	Nombre         string `json:"nombre,omitempty"`
	Correo         string `json:"correo,omitempty"`
	Identificacion string `json:"identificacion,omitempty"`
	Direccion      string `json:"direccion,omitempty"`
}
