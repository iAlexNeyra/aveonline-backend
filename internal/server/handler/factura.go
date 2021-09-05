package handlers

import (
	"aveonline/internal/services"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type registrarFacturaRequest struct {
	ID             int      `json:"id,omitempty"`
	PagoTotal      float64  `json:"pago_total,omitempty"`
	FechaCrear     JsonDate `json:"fecha_crear,omitempty"`
	IDMedicamentos []int    `json:"id_medicamentos,omiatempty,omitempty"`
	IDPromocion    int      `json:"id_promocion,omitempty"`
}

func RegistrarFacturaHandler(service services.FacturaService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req registrarFacturaRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		err := service.RegistrarFactura(
			ctx,
			req.FechaCrear.toTime(),
			req.PagoTotal,
			req.IDPromocion,
			req.IDMedicamentos,
		)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, "factura creada")
	}
}

type listFacturaRequest struct {
	DateFrom string `form:"fecha_inicio"`
	DateTo   string `form:"fecha_fin"`
}

func ListarFacturasHandler(service services.FacturaService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req listFacturaRequest
		ctx.Bind(&req)
		fechaInicio, err := parseDate(req.DateFrom)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "formato de fecha incorrecto")
			return
		}
		fechaFin, err := parseDate(req.DateTo)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "formato de fecha incorrecto")
			return
		}

		facturas := service.ListarFacturas(ctx, fechaInicio, fechaFin)

		if len(facturas) == 0 {
			ctx.JSON(http.StatusOK, make([]string, 0))
			return
		}
		ctx.JSON(http.StatusOK, facturas)
	}
}

type simularRequest struct {
	FechaCompra    string `form:"fecha_compra"`
	IDMedicamentos []int  `form:"id_medicamentos"`
}

func SimularFacturaHandler(service services.FacturaService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req simularRequest
		ctx.Bind(&req)
		fechaCompra, err := parseDate(req.FechaCompra)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "formato de fecha incorrecto")
			return
		}

		descuento := service.SimularFactura(ctx, fechaCompra, req.IDMedicamentos)
		ctx.JSON(http.StatusOK, descuento)
	}
}

func parseDate(value string) (time.Time, error) {
	s := strings.Trim(value, "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}
