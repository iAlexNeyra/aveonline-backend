package services

import (
	"aveonline/internal/entidades"
	"context"
	"time"
)

type FacturaService interface {
	RegistrarFactura(ctx context.Context, fechaCrear time.Time, pagoTotal float64, promocionID int, medicamentos []int) error
	ListarFacturas(ctx context.Context, fechaInicio, fechaFin time.Time) []Factura
	SimularFactura(ctx context.Context, fechaCompra time.Time, medicamentos []int) float64
}

type facturaService struct {
	facturas     entidades.FacturaRepository
	medicamentos entidades.MedicamentoRepository
	promociones  entidades.PromocionRepository
}

func NewFacturaService(facturas entidades.FacturaRepository, medicamentos entidades.MedicamentoRepository, promociones entidades.PromocionRepository) FacturaService {
	return &facturaService{
		facturas:     facturas,
		medicamentos: medicamentos,
		promociones:  promociones,
	}
}

func (s *facturaService) RegistrarFactura(ctx context.Context, fechaCrear time.Time, pagoTotal float64, promocionID int, medicamentoIDs []int) error {
	f := entidades.NewFactura(fechaCrear, pagoTotal)
	if promocionID > 0 {
		promo, err := s.promociones.Find(ctx, promocionID)
		if err != nil {
			return err
		}
		f.AddPromocion(promo)
	}

	for _, medicamentoID := range medicamentoIDs {
		m, err := s.medicamentos.Find(ctx, medicamentoID)
		if err != nil {
			return err
		}
		item := entidades.NewItem(m)
		f.AddItem(item)
	}

	err := s.facturas.Save(ctx, f)

	return err
}

func (s *facturaService) ListarFacturas(ctx context.Context, fechaInicio, fechFin time.Time) []Factura {
	var facturas []Factura
	fl := s.facturas.FindBetweenDateCreated(ctx, fechaInicio, fechFin)
	for _, factura := range fl {
		facturas = append(facturas, assambleFactura(&factura))
	}
	return facturas
}

func (s *facturaService) SimularFactura(ctx context.Context, fechaCompra time.Time, medicamentosIDs []int) float64 {
	medicamentos := s.medicamentos.FindByIds(ctx, medicamentosIDs)
	promos := s.promociones.FindAll(ctx)
	descuento := entidades.NewDescuento(promos)

	return descuento.CalcularDescuento(fechaCompra, medicamentos)
}

type Factura struct {
	ID           int           `json:"id,omitempty"`
	FechaCrear   JsonDate      `json:"fecha_crear,omitempty"`
	PagoTotal    float64       `json:"pago_total,omitempty"`
	Promocion    Promocion     `json:"promocion,omitempty"`
	Medicamentos []Medicamento `json:"medicamentos,omitempty"`
}

func assambleFactura(f *entidades.Factura) Factura {
	return Factura{
		ID:           f.ID,
		FechaCrear:   JsonDate(f.FechaCrear),
		PagoTotal:    f.PagoTotal,
		Promocion:    assamblePromocion(&f.Promocion),
		Medicamentos: assambleItems(f.Medicamentos),
	}
}

func assambleItems(items []entidades.FacturaItem) []Medicamento {
	var medicamentos []Medicamento
	for _, item := range items {
		medicamentos = append(medicamentos, assamble(&item.Medicamento))
	}
	return medicamentos
}
