package entidades

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPromocionValidationFailed(t *testing.T) {
	start := time.Date(2021, time.September, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2021, time.September, 30, 0, 0, 0, 0, time.UTC)
	_, err := NewPromocion("Descuentos Septiembre", 71, start, end)
	assert.Error(t, err)

	promo, err := NewPromocion("Descuentos Septiembre", 70, start, end)
	assert.NoError(t, err)
	promo1, err := NewPromocion("Descuentos Septiembre", 50, start, end)
	promos := []Promocion{promo1}
	err = promo.ValidaExistePromo(promos)
	assert.Error(t, err)
}

func TestPromocionValidationSucceed(t *testing.T) {
	start := time.Date(2021, time.September, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2021, time.September, 30, 0, 0, 0, 0, time.UTC)

	start2 := time.Date(2021, time.November, 10, 0, 0, 0, 0, time.UTC)
	end2 := time.Date(2021, time.November, 30, 0, 0, 0, 0, time.UTC)

	promo, err := NewPromocion("Descuentos Septiembre", 70, start2, end2)
	assert.NoError(t, err)
	promo1, err := NewPromocion("Descuentos Septiembre", 50, start, end)
	promos := []Promocion{promo1}
	err = promo.ValidaExistePromo(promos)
	assert.NoError(t, err)
}
