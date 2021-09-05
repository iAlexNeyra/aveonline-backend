package postgres

import (
	"aveonline/internal/entidades"
	"context"
	"errors"

	"gorm.io/gorm"
)

var ErrNotFound = errors.New("Medicamento no encontrado")

type PGMedicamentoRepository struct {
	db *gorm.DB
}

func NewPGMedicamentoRepository(db *gorm.DB) entidades.MedicamentoRepository {
	return &PGMedicamentoRepository{
		db: db,
	}
}

func (pg *PGMedicamentoRepository) FindByIds(ctx context.Context, ids []int) []entidades.Medicamento {
	var pgl []Medicamento
	pg.db.Find(&pgl, ids)

	var medicamentos []entidades.Medicamento
	for _, m := range pgl {
		medicamentos = append(medicamentos, entidades.Medicamento{
			ID:        m.ID,
			Nombre:    m.Nombre,
			Precio:    m.Precio,
			Ubicacion: m.Ubicacion,
		})
	}

	return medicamentos
}

func (pg *PGMedicamentoRepository) Save(ctx context.Context, medicamento entidades.Medicamento) error {
	m := Medicamento{
		Nombre:    medicamento.Nombre,
		Precio:    medicamento.Precio,
		Ubicacion: medicamento.Ubicacion,
	}
	result := pg.db.Create(&m)

	return result.Error

}

func (pg *PGMedicamentoRepository) FindAll(ctx context.Context) []entidades.Medicamento {
	var pgm []Medicamento
	pg.db.Find(&pgm)

	var medicamentos []entidades.Medicamento
	for _, m := range pgm {
		medicamentos = append(medicamentos, entidades.Medicamento{
			ID:        m.ID,
			Nombre:    m.Nombre,
			Precio:    m.Precio,
			Ubicacion: m.Ubicacion,
		})
	}

	return medicamentos
}

func (pg *PGMedicamentoRepository) Find(ctx context.Context, id int) (entidades.Medicamento, error) {
	var medicamento Medicamento
	pg.db.First(&medicamento, id)

	if medicamento.ID == 0 {
		return entidades.Medicamento{}, ErrNotFound
	}

	m := entidades.Medicamento{
		ID:        medicamento.ID,
		Nombre:    medicamento.Nombre,
		Precio:    medicamento.Precio,
		Ubicacion: medicamento.Ubicacion,
	}
	return m, nil
}
