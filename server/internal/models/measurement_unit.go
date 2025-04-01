package models

import "gorm.io/gorm"

// MeasurementUnit representa uma unidade de medida
type MeasurementUnit struct {
	gorm.Model

	Name         string    `gorm:"size:50;not null;unique" json:"name"`
	Abbreviation string    `gorm:"size:10;not null;unique" json:"abbreviation"`
	Products     []Product `gorm:"foreignKey:UnitID" json:"-"`
}

// TableName especifica o nome da tabela
func (MeasurementUnit) TableName() string {
	return "measurement_units"
}
