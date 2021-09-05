package handlers

import (
	"aveonline/internal/services"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
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
	return json.Marshal(j)
}

func (j JsonDate) Format(s string) string {
	t := time.Time(j)
	return t.Format(s)
}

func (j JsonDate) toTime() time.Time {
	return time.Time(j)
}

type registrarPromocionRequest struct {
	ID          int      `json:"id,omitempty"`
	Descripcion string   `json:"descripcion,omitempty"`
	Porcentaje  float32  `json:"porcentaje,omitempty"`
	FechaInicio JsonDate `json:"fecha_inicio,omitempty"`
	FechaFin    JsonDate `json:"fecha_fin,omitempty"`
}

func RegistrarPromocionHandler(service services.PromocionService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req registrarPromocionRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		err := service.RegistrarPromocion(
			ctx,
			req.Descripcion,
			req.Porcentaje, req.FechaInicio.toTime(),
			req.FechaFin.toTime(),
		)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, "promocion creado")
	}
}

func ListarPromocionesHandler(service services.PromocionService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		promociones := service.ListarPromociones(ctx)
		if len(promociones) == 0 {
			ctx.JSON(http.StatusOK, make([]string, 0))
			return
		}
		ctx.JSON(http.StatusOK, promociones)
	}
}
