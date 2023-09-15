package repository

import (
	"errors"
	"github.com/abrahamkarina/code-review-exercise-one/internal/domain"
)

// RepositoryVehicle is the interface that wraps the basic methods for a vehicle repository.
type RepositoryVehicle interface {
	// GetAll returns all vehicles
	GetAll() (v []*domain.Vehicle, err error)
	GetByDimensions(float64, float64, float64, float64) ([]*domain.Vehicle, error)
	GetByColorAndYear(string, int) ([]*domain.Vehicle, error)
	GetByWeight(float64, float64) ([]*domain.Vehicle, error)
	GetByBrand(brand string) ([]*domain.Vehicle, error)
	PatchFuel(id int, fuelType string) error
	Put(vehicle *domain.Vehicle) error
	GetByTransmission(transmission string) ([]*domain.Vehicle, error)
	Delete(id int) error
	Post(vehicle *domain.Vehicle) error
}

var (
	// ErrRepositoryVehicleInternal is returned when an internal error occurs.
	ErrRepositoryVehicleInternal = errors.New("repository: internal error")

	// ErrRepositoryVehicleNotFound is returned when a vehicle is not found.
	ErrRepositoryVehicleNotFound = errors.New("repository: vehicle not found")
	ErrRepositoryIdInUse         = errors.New("repository: identifier already in use")
)
