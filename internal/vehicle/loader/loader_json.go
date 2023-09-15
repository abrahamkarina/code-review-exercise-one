package loader

import (
	"encoding/json"
	"fmt"
	"github.com/abrahamkarina/code-review-exercise-one/internal/domain"
	"os"
)

// NewLoaderVehicleJSON returns a new instance of a vehicle loader.
func NewLoaderVehicleJSON(path string) *LoaderVehicleJSON {
	return &LoaderVehicleJSON{Path: path}
}

// LoaderVehicleJSON is an struct that implements the LoaderVehicle interface.
type LoaderVehicleJSON struct {
	Path string
}

// Load returns all vehicles.
type VehicleJSON struct {
	ID           int     `json:"id"`
	Brand        string  `json:"brand"`
	Model        string  `json:"model"`
	Registration string  `json:"registration"`
	Year         int     `json:"year"`
	Color        string  `json:"color"`
	MaxSpeed     int     `json:"max_speed"`
	FuelType     string  `json:"fuel_type"`
	Transmission string  `json:"transmission"`
	Passengers   int     `json:"passengers"`
	Height       float64 `json:"height"`
	Width        float64 `json:"width"`
	Weight       float64 `json:"weight"`
}

// Load returns all vehicles.
func (l *LoaderVehicleJSON) Load() (v map[int]*domain.VehicleAttributes, err error) {
	// open file
	f, err := os.Open(l.Path)
	if err != nil {
		err = fmt.Errorf("%w. %v", ErrLoaderVehicleInternal, err)
		return
	}
	defer f.Close()

	// read file
	var vehiclesJSON []*VehicleJSON
	err = json.NewDecoder(f).Decode(&vehiclesJSON)
	if err != nil {
		err = fmt.Errorf("%w. %v", ErrLoaderVehicleInternal, err)
		return
	}

	// serialize vehicles
	v = make(map[int]*domain.VehicleAttributes)
	for _, vehicleJSON := range vehiclesJSON {
		v[vehicleJSON.ID] = &domain.VehicleAttributes{
			Brand:        vehicleJSON.Brand,
			Model:        vehicleJSON.Model,
			Registration: vehicleJSON.Registration,
			Year:         vehicleJSON.Year,
			Color:        vehicleJSON.Color,
			MaxSpeed:     vehicleJSON.MaxSpeed,
			FuelType:     vehicleJSON.FuelType,
			Transmission: vehicleJSON.Transmission,
			Passengers:   vehicleJSON.Passengers,
			Height:       vehicleJSON.Height,
			Width:        vehicleJSON.Width,
			Weight:       vehicleJSON.Weight,
		}
	}

	return
}
