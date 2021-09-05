package services

import (
	"aveonline/internal/storagemock"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegistarMedicamentoFailed(t *testing.T) {
	var medicamentos storagemock.MedicamentoRepository
	medicamentos.On("Save", mock.Anything, mock.AnythingOfType("entidades.Medicamento")).Return(errors.New("Occurrio un error al registrar"))
	s := NewMedicamentoService(&medicamentos)
	err := s.RegistrarMedicamento(context.TODO(), "medicamento test", 1.2, "test ubicacion")
	assert.Error(t, err)
}

func TestRegistarMedicamentoSucceed(t *testing.T) {
	var medicamentos storagemock.MedicamentoRepository
	medicamentos.On("Save", mock.Anything, mock.AnythingOfType("entidades.Medicamento")).Return(nil)
	s := NewMedicamentoService(&medicamentos)
	err := s.RegistrarMedicamento(context.TODO(), "medicamento test", 1.2, "test ubicacion")
	assert.NoError(t, err)
}
