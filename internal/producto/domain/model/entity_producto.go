package model

type Producto struct {
	IdProducto string `json:"id_producto"`
	Nombre     string `json:"nombre"`
	Talla      string `json:"talla"`
	Color      string `json:"color"`
	Cantidad   string `json:"cantidad"`
	Tipo       string `json:"tipo"`
}
