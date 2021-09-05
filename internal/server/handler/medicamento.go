package handlers

import (
	"aveonline/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type registrarRequest struct {
	ID        int     `json:"id"`
	Nombre    string  `json:"nombre"`
	Precio    float64 `json:"precio"`
	Ubicacion string  `json:"ubicacion"`
}

func RegistrarMedicamentoHandler(service services.MedicamentoService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req registrarRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, "medicamento no creado")
			return
		}

		err := service.RegistrarMedicamento(
			ctx,
			req.Nombre,
			req.Precio,
			req.Ubicacion,
		)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, "medicamento creado")
	}
}

func ListarMedicamentoHandler(service services.MedicamentoService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		medicamentos := service.ListarMedicamentos(ctx)
		if len(medicamentos) == 0 {
			ctx.JSON(http.StatusOK, make([]string, 0))
			return
		}
		ctx.JSON(http.StatusOK, medicamentos)
	}
}
