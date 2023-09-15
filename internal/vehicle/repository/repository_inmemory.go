package repository

import (
	"github.com/abrahamkarina/code-review-exercise-one/internal/domain"
	"strings"
)

func NewRepositoryVehicleInMemory(db map[int]*domain.VehicleAttributes) *RepositoryVehicleInMemory {
	return &RepositoryVehicleInMemory{db: db}
}

// RepositoryVehicleInMemory is an struct that represents a vehicle storage in memory.
type RepositoryVehicleInMemory struct {
	// db is the database of vehicles.
	db map[int]*domain.VehicleAttributes
}

// GetAll returns all vehicles
func (s *RepositoryVehicleInMemory) GetAll() (v []*domain.Vehicle, err error) {
	// check if the database is empty
	if len(s.db) == 0 {
		err = ErrRepositoryVehicleNotFound
		return
	}

	// get all vehicles from the database
	v = make([]*domain.Vehicle, 0, len(s.db))
	for key, value := range s.db {
		v = append(v, &domain.Vehicle{
			Id:         key,
			Attributes: *value,
		})
	}

	return
}

func (s *RepositoryVehicleInMemory) GetByDimensions(minHeight float64, maxHeight float64,
	minWidth float64, maxWidth float64) ([]*domain.Vehicle, error) {
	if len(s.db) == 0 {
		err := ErrRepositoryVehicleNotFound
		return nil, err
	}
	v := make([]*domain.Vehicle, 0)
	for key, value := range s.db {
		if minWidth < value.Width && value.Width < maxWidth && minHeight < value.Height && value.Height < maxHeight {
			v = append(v, &domain.Vehicle{
				Id:         key,
				Attributes: *value,
			})
		}
	}

	if len(v) == 0 {
		return nil, ErrRepositoryVehicleNotFound
	}
	return v, nil

}
func (s *RepositoryVehicleInMemory) GetByWeight(min float64, max float64) ([]*domain.Vehicle, error) {
	if len(s.db) == 0 {
		err := ErrRepositoryVehicleNotFound
		return nil, err
	}
	v := make([]*domain.Vehicle, 0)
	for key, value := range s.db {
		if min < value.Weight && value.Weight < max {
			v = append(v, &domain.Vehicle{
				Id:         key,
				Attributes: *value,
			})
		}
	}

	if len(v) == 0 {
		return nil, ErrRepositoryVehicleNotFound
	}
	return v, nil
}

func (s *RepositoryVehicleInMemory) GetByColorAndYear(color string, year int) ([]*domain.Vehicle, error) {
	if len(s.db) == 0 {
		err := ErrRepositoryVehicleNotFound
		return nil, err
	}
	vehicles := make([]*domain.Vehicle, 0)
	for k, v := range s.db {
		if v.Year == year && strings.ToLower(v.Color) == strings.ToLower(color) {
			vehicles = append(vehicles, &domain.Vehicle{
				Id:         k,
				Attributes: *v,
			})
		}
	}
	if len(vehicles) != 0 {
		return vehicles, nil
	}
	return nil, ErrRepositoryVehicleNotFound
}

func (s *RepositoryVehicleInMemory) GetByBrand(brand string) ([]*domain.Vehicle, error) {
	if len(s.db) == 0 {
		err := ErrRepositoryVehicleNotFound
		return nil, err
	}
	vehicles := make([]*domain.Vehicle, 0)
	for k, v := range s.db {
		if strings.ToLower(v.Brand) == strings.ToLower(brand) {
			vehicles = append(vehicles, &domain.Vehicle{
				Id:         k,
				Attributes: *v,
			})
		}
	}
	if len(vehicles) != 0 {
		return vehicles, nil
	}
	return nil, ErrRepositoryVehicleNotFound
}

func (s *RepositoryVehicleInMemory) PatchFuel(id int, fuelType string) error {
	val, ok := s.db[id]
	if !ok {
		return ErrRepositoryVehicleNotFound
	}
	val.FuelType = fuelType
	return nil
}
func (s *RepositoryVehicleInMemory) Put(vehicle *domain.Vehicle) error {
	_, ok := s.db[vehicle.Id]
	if !ok {
		return ErrRepositoryVehicleNotFound
	}
	s.db[vehicle.Id] = &vehicle.Attributes
	return nil
}

func (s *RepositoryVehicleInMemory) GetByTransmission(transmission string) ([]*domain.Vehicle, error) {
	if len(s.db) == 0 {
		err := ErrRepositoryVehicleNotFound
		return nil, err
	}
	vehicles := make([]*domain.Vehicle, 0)
	for k, v := range s.db {
		if strings.ToLower(transmission) == strings.ToLower(v.Transmission) {
			vehicles = append(vehicles, &domain.Vehicle{
				Id:         k,
				Attributes: *v,
			})
		}
	}
	if len(vehicles) != 0 {
		return vehicles, nil
	}
	return nil, ErrRepositoryVehicleNotFound
}
func (s *RepositoryVehicleInMemory) Delete(id int) error {
	if _, found := s.db[id]; !found {
		return ErrRepositoryVehicleNotFound
	}
	delete(s.db, id)
	return nil
}

func (s *RepositoryVehicleInMemory) Post(vehicle *domain.Vehicle) error {
	if _, found := s.db[vehicle.Id]; found {
		return ErrRepositoryIdInUse
	}
	s.db[vehicle.Id] = &vehicle.Attributes
	return nil
}
