package model

import "time"

type Carrito struct {
	IdCarrito  string    `json:"id_carrito"`
	IdProducto string    `json:"id_producto"`
	IdCliente  string    `json:"id_cliente"`
	Fecha      time.Time `json:"fecha"`
}
