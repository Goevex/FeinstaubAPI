package feinstaub

import (
	"fmt"
	"time"

	"app/shared/database"
)

// Name of the table
const tableName = "messdaten"

// Acceleration information
type Acceleration struct {
	ID            string    `db:"ID" json:"id"`
	Time          time.Time `db:"DATETIME" json:"datetime" require:"true"`
	Latitude      string    `db:"LAT" json:"lat" require:"true"`
	Longitude     string    `db:"LNG" json:"long" require:"true"`
	AccelerationX float32   `db:"ACCEL_X" json:"accel_x" require:"true"`
	AccelerationY float32   `db:"ACCEL_Y" json:"accel_y" require:"true"`
	AccelerationZ float32   `db:"ACCEL_Z" json:"accel_z" require:"true"`
}

// Acceleration information
type ParticulateMatter struct {
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
	GAS_NH3       float32   `db:"GAS_NH3" json:"GAS_NH3" require:"true"`
	GAS_CO        float32   `db:"GAS_CO" json:"GAS_CO" require:"true"`
	GAS_NO2       float32   `db:"GAS_NO2" json:"GAS_NO2" require:"true"`
	GAS_C3H8      float32   `db:"GAS_C3H8" json:"GAS_C3H8" require:"true"`
	GAS_C4H10     float32   `db:"GAS_C4H10" json:"GAS_C4H10" require:"true"`
	GAS_CH4       float32   `db:"GAS_CH4" json:"GAS_CH4" require:"true"`
	GAS_H2        float32   `db:"GAS_H2" json:"GAS_H2" require:"true"`
	GAS_C2H5OH    float32   `db:"GAS_C2H5OH" json:"GAS_C2H5OH" require:"true"`
}

// AccelData of entities
type AccelData []Acceleration

// AccelData of entities
type ParticulateData []ParticulateMatter

// New Acceleration
func NewAccel() (*Acceleration, error) {
	var err error
	Acceleration := &Acceleration{}

	// Set the default parameters
	Acceleration.ID, err = database.UUID()
	// If error on UUID generation
	if err != nil {
		return Acceleration, err
	}

	return Acceleration, nil
}

// *****************************************************************************
// Read
// *****************************************************************************

// ReadAll returns all entities of Acceleration
func ReadAllAccel() (AccelData, error) {
	var result AccelData
	err := database.SQL.Select(&result, fmt.Sprintf("SELECT * FROM %v", tableName))
	return result, err
}

// ReadAll returns all entities of ParticulateMatter
func ReadAllParticulate() (ParticulateData, error) {
	var result ParticulateData
	err := database.SQL.Select(&result, fmt.Sprintf("SELECT * FROM %v", tableName))
	return result, err
}
