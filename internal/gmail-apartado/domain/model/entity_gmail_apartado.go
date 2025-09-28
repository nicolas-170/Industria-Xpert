package model

type EmailApartado struct {
	EmailTo     string `json:"emailTo,omitempty"`
	Mensaje     string `json:"mensaje,omitempty"`
	TipoMensaje string `json:"tipoMensaje,omitempty"`
	Imagen      Imagen `json:"imagen,omitempty"`
}

type Imagen struct {
	ImagenBase64 string `json:"imagenBase64,omitempty"`
	ImagenNombre string `json:"imagenNombre,omitempty"`
}
