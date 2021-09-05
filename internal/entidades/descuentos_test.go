package entidades

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCalcularDescuentoZero(t *testing.T) {
	fechaInicio, _ := time.Parse("2006-01-02", "2021-09-01")
	fechaFin, _ := time.Parse("2006-01-02", "2021-09-01")
	fechaCompra := time.Now()
	promo, _ := NewPromocion("Descuentos Septiembre", 10, fechaInicio, fechaFin)
	promos := []Promocion{promo}
	descuento := NewDescuento(promos)
	medicamento, _ := NewMedicamento("Tesst", 100, "D")
	medicamentos := []Medicamento{medicamento}

	result := descuento.CalcularDescuento(fechaCompra, medicamentos)

	assert.Equal(t, result, 0.0)
}

func TestCalcularDescuento(t *testing.T) {
	fechaInicio, _ := time.Parse("2006-01-02", "2021-09-01")
	fechaFin, _ := time.Parse("2006-01-02", "2021-09-05")
	fechaCompra := time.Now()
	promo, _ := NewPromocion("Descuentos Septiembre", 10, fechaInicio, fechaFin)
	promos := []Promocion{promo}
	descuento := NewDescuento(promos)
	medicamento, _ := NewMedicamento("Tesst", 100, "D")
	medicamentos := []Medicamento{medicamento}

	result := descuento.CalcularDescuento(fechaCompra, medicamentos)

	assert.Equal(t, result, 10.0)
}
