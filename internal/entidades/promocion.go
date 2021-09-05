package entidades

import (
	"context"
	"errors"
	"time"
)

const PORCENTAJE_MAXIMO = 70.0

var ErrPorcentajeDescuentoSuperado = errors.New("El porcentaje descuento no debe superar el 70%")
var ErrPromocionYaExiste = errors.New("Ya existe una promocion en ese periodo")

type Promocion struct {
	ID          int
	Descripcion string
	Porcentaje  float32
	FechaInicio time.Time
	FechaFin    time.Time
}

func NewPromocion(description string, porcentaje float32, fechaInicio time.Time, fechaFin time.Time) (Promocion, error) {
	if porcentaje > PORCENTAJE_MAXIMO {
		return Promocion{}, ErrPorcentajeDescuentoSuperado
	}
	return Promocion{
		Descripcion: description,
		Porcentaje:  porcentaje,
		FechaInicio: fechaInicio,
		FechaFin:    fechaFin,
	}, nil
}

func (p *Promocion) ValidaExistePromo(promos []Promocion) error {
	for date := p.FechaInicio; date.Before(p.FechaFin.Add(24 * time.Hour)); date = date.Add(24 * time.Hour) {
		for _, promo := range promos {
			startDate := promo.FechaInicio.Add(-24 * time.Hour)
			endDate := promo.FechaFin.Add(24 * time.Hour)
			if date.After(startDate) && date.Before(endDate) {
				return ErrPromocionYaExiste
			}
		}
	}
	return nil
}

type PromocionRepository interface {
	Save(ctx context.Context, promocion Promocion) error
	FindAll(ctx context.Context) []Promocion
	Find(ctx context.Context, id int) (Promocion, error)
}
