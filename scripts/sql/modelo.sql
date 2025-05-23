CREATE TABLE cliente (
  id_cliente VARCHAR(50) PRIMARY KEY,
  celular INT,
  nombre VARCHAR(100),
  correo VARCHAR(256),
  identificacion VARCHAR(100),
  direccion VARCHAR(150)
);

CREATE TABLE producto (
  id_producto VARCHAR(50) PRIMARY KEY,
  nombre VARCHAR(100),
  talla VARCHAR(10), 
  color VARCHAR(30),
  cantidad INT,
  tipo VARCHAR(50)
);

CREATE TABLE contacto (
  id_contacto VARCHAR(50) PRIMARY KEY,
  whatsapp VARCHAR(50),
  correo VARCHAR(256),
  direccion VARCHAR(150)
);

CREATE TABLE carrito (
  id_carrito VARCHAR(50) PRIMARY KEY,
  id_producto VARCHAR(50),
  id_cliente VARCHAR(50) UNIQUE,
  fecha DATE,
  CONSTRAINT fk_carrito_producto FOREIGN KEY (id_producto) REFERENCES producto(id_producto),
  CONSTRAINT fk_carrito_cliente FOREIGN KEY (id_cliente) REFERENCES cliente(id_cliente)
);
