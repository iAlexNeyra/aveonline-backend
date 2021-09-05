CREATE TABLE IF NOT EXISTS medicamentos(
  id serial primary key
  ,nombre varchar(50)
  ,precio  numeric(14,2)
  ,ubicacion varchar(50)
);

CREATE TABLE IF NOT EXISTS promociones(
  id serial primary key
  ,descripcion varchar(100)
  ,porcentaje numeric(5,2)
  ,fecha_inicio date
  ,fecha_fin date
);

CREATE TABLE IF NOT EXISTS facturas(
  id serial primary key
  ,fecha_crear date
  ,promocion_id int
  ,pago_total numeric(12, 2)
);

CREATE TABLE IF NOT EXISTS factura_items(
  id serial primary key
  ,factura_id int
  ,medicamento_id int
)
