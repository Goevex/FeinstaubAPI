package feinstaub

import (
	"fmt"
	"time"

	"app/shared/database"
)

// Name of the table
const tableName = "messdaten"

// MeasuredData information
type MeasuredData struct {
	ID            string    `db:"ID" json:"id"`
	Time          time.Time `db:"DATETIME" json:"datetime" require:"true"`
	Latitude      string    `db:"LAT" json:"lat" require:"true"`
	Longitude     string    `db:"LNG" json:"long" require:"true"`
	AccelerationX float32   `db:"ACCEL_X" json:"accel_x" require:"true"`
	AccelerationY float32   `db:"ACCEL_Y" json:"accel_y" require:"true"`
	AccelerationZ float32   `db:"ACCEL_Z" json:"accel_z" require:"true"`
	Sats          int       `db:"SATS" json:"sats" require:"true"`
	Speed         float32   `db:"SPEED" json:"speed" require:"true"`
	PM25          float32   `db:"PM25" json:"PM25" require:"true"`
	PM10          float32   `db:"PM10" json:"PM10" require:"true"`
	GasNH3        float32   `db:"GAS_NH3" json:"GAS_NH3" require:"true"`
	GasCO         float32   `db:"GAS_CO" json:"GAS_CO" require:"true"`
	GasNO2        float32   `db:"GAS_NO2" json:"GAS_NO2" require:"true"`
	GasC3H8       float32   `db:"GAS_C3H8" json:"GAS_C3H8" require:"true"`
	GasC4H10      float32   `db:"GAS_C4H10" json:"GAS_C4H10" require:"true"`
	GasCH4        float32   `db:"GAS_CH4" json:"GAS_CH4" require:"true"`
	GasH2         float32   `db:"GAS_H2" json:"GAS_H2" require:"true"`
	GasC2H5OH     float32   `db:"GAS_C2H5OH" json:"GAS_C2H5OH" require:"true"`
}

// ParticulateData of entities
type ParticulateData []MeasuredData

// *****************************************************************************
// Read
// *****************************************************************************

// ReadMeasurements returns all entities of ParticulateMatter
func ReadMeasurements(condition string) (ParticulateData, error) {
	var result ParticulateData
	if len(condition) > 0 {
		condition = " WHERE " + condition
	}
	err := database.SQL.Select(&result, fmt.Sprintf("SELECT * FROM %v%v", tableName, condition))
	return result, err
}
