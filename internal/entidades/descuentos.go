package entidades

import (
	"time"
)

type Descuento struct {
	Promociones []Promocion
}

func NewDescuento(promociones []Promocion) Descuento {
	return Descuento{
		Promociones: promociones,
	}
}

func (d *Descuento) CalcularDescuento(fechaCompra time.Time, medicamentos []Medicamento) float64 {
	var porcentajeDescuento float32
	porcentajeDescuento = 0.0
	for _, promo := range d.Promociones {
		startDate := promo.FechaInicio.Add(-24 * time.Hour)
		endDate := promo.FechaFin.Add(24 * time.Hour)
		if fechaCompra.After(startDate) && fechaCompra.Before(endDate) {
			porcentajeDescuento = promo.Porcentaje
			break
		}
	}

	var pagoTotal float64
	pagoTotal = 0
	for _, medicamento := range medicamentos {
		pagoTotal = pagoTotal + medicamento.Precio
	}

	return (float64(porcentajeDescuento) * pagoTotal) / 100
}
