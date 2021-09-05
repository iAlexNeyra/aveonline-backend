package main

import (
	handlers "aveonline/internal/server/handler"
	"aveonline/internal/services"
	"aveonline/internal/storage/postgres"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=test password=test dbname=avionlinedb port=5432 sslmode=disable"
	db, err := gorm.Open(pg.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error al conectarse a la base de datos: %v", err)
	}

	medicamentos := postgres.NewPGMedicamentoRepository(db)
	promociones := postgres.NewPGPromocionRepository(db)
	facturas := postgres.NewPGFacturaRepository(db)

	sm := services.NewMedicamentoService(medicamentos)
	sp := services.NewPromocionService(promociones)
	fs := services.NewFacturaService(facturas, medicamentos, promociones)

	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	r.Use(cors.New(config))
	r.POST("/v1/medicamento", handlers.RegistrarMedicamentoHandler(sm))
	r.GET("/v1/medicamento", handlers.ListarMedicamentoHandler(sm))
	r.POST("/v1/promocion", handlers.RegistrarPromocionHandler(sp))
	r.GET("/v1/promocion", handlers.ListarPromocionesHandler(sp))
	r.GET("/v1/factura", handlers.ListarFacturasHandler(fs))
	r.GET("/v1/factura/simular", handlers.SimularFacturaHandler(fs))
	r.POST("/v1/factura", handlers.RegistrarFacturaHandler(fs))
	r.Run(":9092")
}
