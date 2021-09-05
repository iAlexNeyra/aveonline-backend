package postgres

import "time"

type Factura struct {
	ID          int `gorm:"primaryKey"`
	FechaCrear  time.Time
	PagoTotal   float64
	PromocionID int
}

type FacturaItem struct {
	ID            int `gorm:"primaryKey"`
	MedicamentoID int
	FacturaID     int
}
