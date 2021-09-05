package services

import (
	"aveonline/internal/entidades"
	"context"
	"fmt"
	"strings"
	"time"
)

type JsonDate time.Time

func (j *JsonDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = JsonDate(t)
	return nil
}

func (j JsonDate) MarshalJSON() ([]byte, error) {
	return []byte(j.Format("2006-01-02")), nil
}

func (j JsonDate) Format(s string) string {
	t := time.Time(j)
	return fmt.Sprintf("%q", t.Format(s))
}

type PromocionService interface {
	RegistrarPromocion(ctx context.Context, descripcion string, porcentaje float32, fechaInicio, fechaFin time.Time) error
	ListarPromociones(ctx context.Context) []Promocion
}

type promocionService struct {
	promociones entidades.PromocionRepository
}

func NewPromocionService(promociones entidades.PromocionRepository) PromocionService {
	return &promocionService{
		promociones: promociones,
	}
}

func (s *promocionService) RegistrarPromocion(ctx context.Context, descripcion string, porcentaje float32, fechaInicio, fechaFin time.Time) error {
	p, err := entidades.NewPromocion(
		descripcion,
		porcentaje,
		fechaInicio,
		fechaFin,
	)

	if err != nil {
		return err
	}

	promociones := s.promociones.FindAll(ctx)
	err = p.ValidaExistePromo(promociones)

	if err != nil {
		return err
	}

	err = s.promociones.Save(ctx, p)

	return err
}

func (s *promocionService) ListarPromociones(ctx context.Context) []Promocion {
	var promociones []Promocion
	pl := s.promociones.FindAll(ctx)
	for _, promocion := range pl {
		promociones = append(promociones, assamblePromocion(&promocion))
	}
	return promociones
}

type Promocion struct {
	ID          int      `json:"id"`
	Descripcion string   `json:"descripcion,omitempty"`
	Porcentaje  float32  `json:"porcentaje,omitempty"`
	FechaInicio JsonDate `json:"fecha_inicio,omitempty"`
	FechaFin    JsonDate `json:"fecha_fin,omitempty"`
}

func assamblePromocion(p *entidades.Promocion) Promocion {
	return Promocion{
		ID:          p.ID,
		Porcentaje:  p.Porcentaje,
		Descripcion: p.Descripcion,
		FechaInicio: JsonDate(p.FechaInicio),
		FechaFin:    JsonDate(p.FechaFin),
	}
}
