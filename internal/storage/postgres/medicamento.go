package postgres

type Medicamento struct {
	ID        int `gorm:"primaryKey"`
	Nombre    string
	Precio    float64
	Ubicacion string
}
