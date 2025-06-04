package model

type Cliente struct {
	IdCliente      string `json:"id_cliente"`
	Celular        string `json:"celular"`
	Nombre         string `json:"nombre"`
	Correo         string `json:"correo"`
	Identificacion string `json:"identificacion"`
	Direccion      string `json:"direccion"`
}
