package services

import (
	"aveonline/internal/entidades"
	"context"
)

type MedicamentoService interface {
	RegistrarMedicamento(ctx context.Context, nombre string, precio float64, ubicacion string) error
	ListarMedicamentos(ctx context.Context) []Medicamento
}

type medicamentoService struct {
	medicamentos entidades.MedicamentoRepository
}

func NewMedicamentoService(medicamentos entidades.MedicamentoRepository) MedicamentoService {
	return &medicamentoService{
		medicamentos: medicamentos,
	}
}

func (s *medicamentoService) RegistrarMedicamento(ctx context.Context, nombre string, precio float64, ubicacion string) error {
	m, err := entidades.NewMedicamento(
		nombre,
		precio,
		ubicacion,
	)

	if err != nil {
		return err
	}

	err = s.medicamentos.Save(ctx, m)

	return err
}

func (s *medicamentoService) ListarMedicamentos(ctx context.Context) []Medicamento {
	var medicamentos []Medicamento
	ml := s.medicamentos.FindAll(ctx)
	for _, medicamento := range ml {
		medicamentos = append(medicamentos, assamble(&medicamento))
	}
	return medicamentos
}

type Medicamento struct {
	ID        int     `json:"id,omitempty"`
	Nombre    string  `json:"nombre,omitempty"`
	Precio    float64 `json:"precio,omitempty"`
	Ubicacion string  `json:"ubicacionomitempty"`
}

func assamble(m *entidades.Medicamento) Medicamento {
	return Medicamento{
		ID:        m.ID,
		Nombre:    m.Nombre,
		Precio:    m.Precio,
		Ubicacion: m.Ubicacion,
	}
}
