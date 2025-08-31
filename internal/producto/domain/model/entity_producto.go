package model

type Producto struct {
	IdProducto string  `json:"id_producto,omitempty"`
	Nombre     string  `json:"nombre,omitempty"`
	Talla      string  `json:"talla,omitempty"`
	Color      string  `json:"color,omitempty"`
	Cantidad   string  `json:"cantidad,omitempty"`
	Tipo       string  `json:"tipo,omitempty"`
	Precio     *float64 `json:"precio,omitempty"`
}
