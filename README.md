# Industria-Xpert

Desarrollo del proyecto industria Xpert


## 🌐 Variables de entorno

A continuación, se listan las variables de entorno necesarias para la correcta ejecución del servicio:

- `PORT`: Indica el puerto en el que se va a exponer el servicio (por ejemplo, `7777`), por defecto esta en el 7777

## 🧩 Modelo de datos (SQL)

El modelado de la base de datos se encuentra en el archivo: scripts/sql/modelo.sql
Este contiene las entidades necesarias para el proyecto.

## ▶️ Para correr el servicio

### 1. Requisitos

- Tener instalado [Go](https://golang.org/doc/install)
- Tener acceso a una base de datos PostgreSQL (u otra compatible)
- Configurar las variables de entorno mencionadas anteriormente

### 2. Ejecutar en modo desarrollo

- En terminal:

```bash
go run ./cmd/server/main.go
```