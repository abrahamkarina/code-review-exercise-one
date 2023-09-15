package service

import (
	"github.com/abrahamkarina/code-review-exercise-one/internal/domain"
	"github.com/abrahamkarina/code-review-exercise-one/internal/vehicle/repository"
)

// ServiceVehicleDefault is an struct that represents a vehicle service.
type ServiceVehicleDefault struct {
	rp         repository.RepositoryVehicle
	errAdapter ServiceErrorAdapter
}

type ServiceErrorAdapter func(error) error

// NewServiceVehicleDefault returns a new instance of a vehicle service.
func NewServiceVehicleDefault(rp repository.RepositoryVehicle, adapter ServiceErrorAdapter) *ServiceVehicleDefault {
	return &ServiceVehicleDefault{rp: rp,
		errAdapter: adapter,
	}
}

// GetAll returns all vehicles.
func (s *ServiceVehicleDefault) GetAll() ([]*domain.Vehicle, error) {
	v, err := s.rp.GetAll()
	if err != nil {
		return nil, s.errAdapter(err)

	}

	return v, err
}

func (s *ServiceVehicleDefault) GetByDimensions(minHeight float64, maxHeight float64, minWidth float64, maxWidth float64) ([]*domain.Vehicle, error) {
	v, err := s.rp.GetByDimensions(minHeight, maxHeight, minWidth, maxWidth)
	if err != nil {
		return nil, s.errAdapter(err)
	}

	return v, err
}

func (s *ServiceVehicleDefault) SearchByColorAndYear(color string, year int) (v []*domain.Vehicle, err error) {
	v, err = s.rp.GetByColorAndYear(color, year)
	if err != nil {
		return nil, s.errAdapter(err)
	}

	return
}
func (s *ServiceVehicleDefault) GetByWeight(min float64, max float64) ([]*domain.Vehicle, error) {
	v, err := s.rp.GetByWeight(min, max)
	if err != nil {
		return nil, s.errAdapter(err)
	}

	return v, err
}

func (s *ServiceVehicleDefault) GetAverageCapacityByBrand(brand string) (float64, error) {
	vehicles, err := s.rp.GetByBrand(brand)
	if err != nil {
		return 0, s.errAdapter(err)
	}
	if len(vehicles) == 0 {
		return 0, ErrServiceVehicleNotFound
	}

	count := 0
	for _, v := range vehicles {
		count += v.Attributes.Passengers
	}
	return float64(count) / float64(len(vehicles)), nil
}

func (s *ServiceVehicleDefault) PatchFuel(id int, fuelType string) error {
	err := s.rp.PatchFuel(id, fuelType)
	if err != nil {
		return s.errAdapter(err)
	}
	return nil
}

func (s *ServiceVehicleDefault) Put(vehicle *domain.Vehicle) error {
	err := s.rp.Put(vehicle)
	if err != nil {
		return s.errAdapter(err)
	}
	return nil
}

func (s *ServiceVehicleDefault) GetByTransmission(transmission string) ([]*domain.Vehicle, error) {
	v, err := s.rp.GetByTransmission(transmission)
	if err != nil {
		return nil, s.errAdapter(err)
	}

	return v, err
}

func (s *ServiceVehicleDefault) Delete(id int) error {
	err := s.rp.Delete(id)
	if err != nil {
		return s.errAdapter(err)
	}
	return nil
}

func (s *ServiceVehicleDefault) Batch(vehicles []*domain.Vehicle) error {
	for _, v := range vehicles {
		err := s.Post(v)
		if err != nil {
			return s.errAdapter(err)
		}
	}
	return nil
}
func (s *ServiceVehicleDefault) Post(vehicle *domain.Vehicle) error {
	return s.rp.Post(vehicle)
}
