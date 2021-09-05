package postgres

import (
	"aveonline/internal/entidades"
	"context"
	"errors"

	"gorm.io/gorm"
)

var ErrNotFoundPromocion = errors.New("No se encontro la promocion")

type PGPromocionRepository struct {
	db *gorm.DB
}

func NewPGPromocionRepository(db *gorm.DB) entidades.PromocionRepository {
	return &PGPromocionRepository{
		db: db,
	}
}

func (pg *PGPromocionRepository) Save(ctx context.Context, promocion entidades.Promocion) error {
	p := Promocion{
		Descripcion: promocion.Descripcion,
		Porcentaje:  promocion.Porcentaje,
		FechaInicio: promocion.FechaInicio,
		FechaFin:    promocion.FechaFin,
	}
	result := pg.db.Create(&p)

	return result.Error

}

func (pg *PGPromocionRepository) FindAll(ctx context.Context) []entidades.Promocion {
	var pgp []Promocion
	pg.db.Find(&pgp)

	var promociones []entidades.Promocion
	for _, p := range pgp {
		promociones = append(promociones, entidades.Promocion{
			ID:          p.ID,
			Descripcion: p.Descripcion,
			Porcentaje:  p.Porcentaje,
			FechaInicio: p.FechaInicio,
			FechaFin:    p.FechaFin,
		})
	}

	return promociones
}

func (pg *PGPromocionRepository) Find(ctx context.Context, id int) (entidades.Promocion, error) {
	var promo Promocion
	pg.db.First(&promo, id)

	if promo.ID == 0 {
		return entidades.Promocion{}, ErrNotFoundPromocion
	}

	return entidades.Promocion{
		ID:          promo.ID,
		Descripcion: promo.Descripcion,
		Porcentaje:  promo.Porcentaje,
		FechaInicio: promo.FechaInicio,
		FechaFin:    promo.FechaFin,
	}, nil
}
