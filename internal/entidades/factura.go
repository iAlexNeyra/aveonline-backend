package entidades

import (
	"context"
	"time"
)

type Factura struct {
	ID           int
	FechaCrear   time.Time
	PagoTotal    float64
	Promocion    Promocion
	Medicamentos []FacturaItem
}

func NewFactura(fechaCrear time.Time, pagoTotal float64) Factura {
	return Factura{
		FechaCrear: fechaCrear,
		PagoTotal:  pagoTotal,
	}
}

func (f *Factura) AddItem(item FacturaItem) error {
	f.Medicamentos = append(f.Medicamentos, item)
	return nil
}

func (f *Factura) AddPromocion(promocion Promocion) {
	f.Promocion = promocion
}

type FacturaItem struct {
	Id          int
	Medicamento Medicamento
}

func NewItem(medicamento Medicamento) FacturaItem {
	return FacturaItem{
		Medicamento: medicamento,
	}
}

type FacturaRepository interface {
	Save(ctx context.Context, factura Factura) error
	FindBetweenDateCreated(ctx context.Context, dateFrom, dateTo time.Time) []Factura
	FindItemsByFacturaID(ctx context.Context, facturaID int) []FacturaItem
}
