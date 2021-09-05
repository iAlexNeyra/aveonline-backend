package postgres

import "time"

type Promocion struct {
	ID          int `gorm:"primaryKey"`
	Descripcion string
	Porcentaje  float32
	FechaInicio time.Time
	FechaFin    time.Time
}

func (p *Promocion) TableName() string {
	return "promociones"
}
