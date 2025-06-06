package model

import "time"

type Carrito struct {
	IdCarrito  string    `json:"id_carrito,omitempty"`
	IdProducto string    `json:"id_producto,omitempty"`
	IdCliente  string    `json:"id_cliente,omitempty"`
	Fecha      time.Time `json:"fecha,omitempty"`
}
