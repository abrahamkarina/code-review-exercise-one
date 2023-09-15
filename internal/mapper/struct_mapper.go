package mapper

import (
	"github.com/abrahamkarina/code-review-exercise-one/cmd/web"
	"github.com/abrahamkarina/code-review-exercise-one/internal/domain"
)

type StructMapper interface {
	MapToVehicleHandlerGetByDimension(vehicle domain.Vehicle) web.VehicleHandlerGetByDimension
	MapFromModelVehicleHandlerGetByColorAndDate(vehicle domain.Vehicle) web.VehicleHandlerGetByColorAndDate
	MapToVehicleHandlerGetByWeight(vehicle domain.Vehicle) web.VehicleHandlerGetByWeight
	MapFromVehicleHandlerPutFuel(id int, vehicle web.VehicleHandlerPutFuel) *domain.Vehicle
	MapToVehicleHandlerGetByTransmission(vehicle domain.Vehicle) web.VehicleHandlerGetByTransmission
	MapToVehicleHandlerBatch(vehicles []web.VehicleHandlerPost) []*domain.Vehicle
	MapToVehicleHandlerPost(vehicles web.VehicleHandlerPost) *domain.Vehicle
}

type structMapper struct {
}

func NewStructMapper() StructMapper {
	return &structMapper{}
}

func (sm *structMapper) MapToVehicleHandlerGetByDimension(vehicle domain.Vehicle) web.VehicleHandlerGetByDimension {
	return web.VehicleHandlerGetByDimension{
		Id:           vehicle.Id,
		Brand:        vehicle.Attributes.Brand,
		Model:        vehicle.Attributes.Model,
		Registration: vehicle.Attributes.Registration,
		Year:         vehicle.Attributes.Year,
		Color:        vehicle.Attributes.Color,
		MaxSpeed:     vehicle.Attributes.MaxSpeed,
		FuelType:     vehicle.Attributes.FuelType,
		Transmission: vehicle.Attributes.Transmission,
		Passengers:   vehicle.Attributes.Passengers,
		Height:       vehicle.Attributes.Height,
		Width:        vehicle.Attributes.Width,
		Weight:       vehicle.Attributes.Weight,
	}
}

func (sm *structMapper) MapFromModelVehicleHandlerGetByColorAndDate(vehicle domain.Vehicle) web.VehicleHandlerGetByColorAndDate {
	return web.VehicleHandlerGetByColorAndDate{
		Id:           vehicle.Id,
		Brand:        vehicle.Attributes.Brand,
		Model:        vehicle.Attributes.Model,
		Registration: vehicle.Attributes.Registration,
		Year:         vehicle.Attributes.Year,
		Color:        vehicle.Attributes.Color,
		MaxSpeed:     vehicle.Attributes.MaxSpeed,
		FuelType:     vehicle.Attributes.FuelType,
		Transmission: vehicle.Attributes.Transmission,
		Passengers:   vehicle.Attributes.Passengers,
		Height:       vehicle.Attributes.Height,
		Width:        vehicle.Attributes.Width,
		Weight:       vehicle.Attributes.Weight,
	}
}
func (sm *structMapper) MapToVehicleHandlerGetByWeight(vehicle domain.Vehicle) web.VehicleHandlerGetByWeight {
	return web.VehicleHandlerGetByWeight{
		Id:           vehicle.Id,
		Brand:        vehicle.Attributes.Brand,
		Model:        vehicle.Attributes.Model,
		Registration: vehicle.Attributes.Registration,
		Year:         vehicle.Attributes.Year,
		Color:        vehicle.Attributes.Color,
		MaxSpeed:     vehicle.Attributes.MaxSpeed,
		FuelType:     vehicle.Attributes.FuelType,
		Transmission: vehicle.Attributes.Transmission,
		Passengers:   vehicle.Attributes.Passengers,
		Height:       vehicle.Attributes.Height,
		Width:        vehicle.Attributes.Width,
		Weight:       vehicle.Attributes.Weight,
	}
}

func (sm *structMapper) MapFromVehicleHandlerPutFuel(id int, vehicle web.VehicleHandlerPutFuel) *domain.Vehicle {
	return &domain.Vehicle{
		Id: id,
		Attributes: domain.VehicleAttributes{
			Brand:        vehicle.Brand,
			Model:        vehicle.Model,
			Registration: vehicle.Registration,
			Year:         vehicle.Year,
			Color:        vehicle.Color,
			MaxSpeed:     vehicle.MaxSpeed,
			FuelType:     vehicle.FuelType,
			Transmission: vehicle.Transmission,
			Passengers:   vehicle.Passengers,
			Height:       vehicle.Height,
			Width:        vehicle.Width,
			Weight:       vehicle.Weight,
		},
	}
}

func (sm *structMapper) MapToVehicleHandlerGetByTransmission(vehicle domain.Vehicle) web.VehicleHandlerGetByTransmission {
	return web.VehicleHandlerGetByTransmission{
		Id:           vehicle.Id,
		Brand:        vehicle.Attributes.Brand,
		Model:        vehicle.Attributes.Model,
		Registration: vehicle.Attributes.Registration,
		Year:         vehicle.Attributes.Year,
		Color:        vehicle.Attributes.Color,
		MaxSpeed:     vehicle.Attributes.MaxSpeed,
		FuelType:     vehicle.Attributes.FuelType,
		Transmission: vehicle.Attributes.Transmission,
		Passengers:   vehicle.Attributes.Passengers,
		Height:       vehicle.Attributes.Height,
		Width:        vehicle.Attributes.Width,
		Weight:       vehicle.Attributes.Weight,
	}
}

func (sm *structMapper) MapToVehicleHandlerBatch(vehicles []web.VehicleHandlerPost) []*domain.Vehicle {
	vs := make([]*domain.Vehicle, 0)
	for _, v := range vehicles {
		vs = append(vs, sm.MapToVehicleHandlerPost(v))
	}
	return vs
}
func (sm *structMapper) MapToVehicleHandlerPost(vehicle web.VehicleHandlerPost) *domain.Vehicle {
	return &domain.Vehicle{
		Id: vehicle.Id,
		Attributes: domain.VehicleAttributes{
			Brand:        vehicle.Brand,
			Model:        vehicle.Model,
			Registration: vehicle.Registration,
			Year:         vehicle.Year,
			Color:        vehicle.Color,
			MaxSpeed:     vehicle.MaxSpeed,
			FuelType:     vehicle.FuelType,
			Transmission: vehicle.Transmission,
			Passengers:   vehicle.Passengers,
			Height:       vehicle.Height,
			Width:        vehicle.Width,
			Weight:       vehicle.Weight,
		},
	}
}
