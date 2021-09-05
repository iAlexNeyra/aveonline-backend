package entidades

import (
	"context"
	"errors"
)

type Medicamento struct {
	ID        int
	Nombre    string
	Precio    float64
	Ubicacion string
}

func NewMedicamento(nombre string, precio float64, ubicacion string) (Medicamento, error) {
	if nombre == "" {
		return Medicamento{}, errors.New("Ingrese nombre del medicamento")
	}
	return Medicamento{
		Nombre:    nombre,
		Precio:    precio,
		Ubicacion: ubicacion,
	}, nil
}

type MedicamentoRepository interface {
	Save(ctx context.Context, medicamento Medicamento) error
	FindAll(ctx context.Context) []Medicamento
	FindByIds(ctx context.Context, ids []int) []Medicamento
	Find(ctx context.Context, id int) (Medicamento, error)
}

//go:generate mockery --case=snake --outpkg=storagemock --output=../storagemock --name=MedicamentoRepository
