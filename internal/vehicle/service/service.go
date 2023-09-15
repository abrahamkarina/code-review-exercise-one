package service

import (
	"errors"
	"github.com/abrahamkarina/code-review-exercise-one/internal/domain"
)

// ServiceVehicle is the interface that wraps the basic methods for a vehicle service.
// - conections with external apis
// - business logic
type ServiceVehicle interface {
	// GetAll returns all vehicles
	GetAll() (v []*domain.Vehicle, err error)
	GetByDimensions(minLength float64, maxLength float64, minWidth float64, maxWidth float64) ([]*domain.Vehicle, error)
	SearchByColorAndYear(color string, year int) (v []*domain.Vehicle, err error)
	GetByWeight(weight float64, weight2 float64) ([]*domain.Vehicle, error)
	GetAverageCapacityByBrand(brand string) (float64, error)
	PatchFuel(id int, fuelType string) error
	Put(vehicle *domain.Vehicle) error
	GetByTransmission(transmission string) ([]*domain.Vehicle, error)
	Delete(id int) error
	Batch([]*domain.Vehicle) error
	Post(*domain.Vehicle) error
}

var (
	// ErrServiceVehicleInternal is returned when an internal error occurs.
	ErrServiceVehicleInternal = errors.New("service: internal error")

	// ErrServiceVehicleNotFound is returned when no vehicle is found.
	ErrServiceVehicleNotFound = errors.New("service: vehicle not found")
	ErrServiceVIdInUse        = errors.New("service: identifier already in use")
)
