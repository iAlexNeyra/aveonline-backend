package postgres

import (
	"aveonline/internal/entidades"
	"context"
	"time"

	"gorm.io/gorm"
)

type PGFacturaRepository struct {
	db *gorm.DB
}

func NewPGFacturaRepository(db *gorm.DB) entidades.FacturaRepository {
	return &PGFacturaRepository{
		db: db,
	}
}

func (pg *PGFacturaRepository) FindItemsByFacturaID(ctx context.Context, id int) []entidades.FacturaItem {
	var items []entidades.FacturaItem
	var pgfil []FacturaItem
	pg.db.Find(&pgfil, "factura_id = ?", id)
	for _, item := range pgfil {
		items = append(items, entidades.FacturaItem{
			Id:          item.ID,
			Medicamento: entidades.Medicamento{ID: item.MedicamentoID},
		})
	}

	return items
}

func (pg *PGFacturaRepository) Save(ctx context.Context, factura entidades.Factura) error {
	f := Factura{
		FechaCrear:  factura.FechaCrear,
		PagoTotal:   factura.PagoTotal,
		PromocionID: factura.Promocion.ID,
	}
	result := pg.db.Create(&f)

	if result.Error != nil {
		return result.Error
	}

	var items []FacturaItem
	for _, item := range factura.Medicamentos {
		items = append(items, FacturaItem{
			MedicamentoID: item.Medicamento.ID,
			FacturaID:     f.ID,
		})
	}

	result = pg.db.Create(&items)

	return result.Error
}

type facturaResult struct {
	ID            int
	FechaCrear    time.Time
	PromocionID   int
	PagoTotal     float64
	Descripcion   string
	Porcentaje    float32
	FechaInicio   time.Time
	FechaFin      time.Time
	MedicamentoID int
	Nombre        string
	Precio        float64
	Ubicacion     string
}

func (pg *PGFacturaRepository) FindBetweenDateCreated(ctx context.Context, dateFrom, dateTo time.Time) []entidades.Factura {

	var result []facturaResult
	sql := `
	select 
	f.*
	,p.descripcion
	,p.porcentaje 
	,p.fecha_inicio 
	,p.fecha_fin 
	,fi.medicamento_id 
	,m.nombre
	,m.precio 
	,m.ubicacion
from 
	facturas f 
inner join promociones p on p.id  = f.promocion_id
inner  join factura_items fi on fi.factura_id = f.id
inner  join medicamentos m on m.id = fi.medicamento_id where
	f.fecha_crear  between  ? and  ?
	`
	pg.db.Raw(sql, dateFrom, dateTo).Scan(&result)
	var facturas []entidades.Factura
	for _, f := range result {
		if checkFacturaID(facturas, f.ID) {
			continue
		}
		ef := entidades.Factura{
			ID:         f.ID,
			FechaCrear: f.FechaCrear,
			PagoTotal:  f.PagoTotal,
			Promocion: entidades.Promocion{
				ID:          f.PromocionID,
				Descripcion: f.Descripcion,
				Porcentaje:  f.Porcentaje,
				FechaInicio: f.FechaInicio,
				FechaFin:    f.FechaFin,
			},
		}

		for _, fi := range result {
			if f.ID == fi.ID {
				item := entidades.NewItem(entidades.Medicamento{
					ID:        fi.MedicamentoID,
					Nombre:    fi.Nombre,
					Precio:    fi.Precio,
					Ubicacion: fi.Ubicacion,
				})
				ef.AddItem(item)
			}
		}
		facturas = append(facturas, ef)
	}

	return facturas
}

func checkFacturaID(facturas []entidades.Factura, id int) bool {
	for _, factura := range facturas {
		if factura.ID == id {
			return true
		}
	}
	return false
}
