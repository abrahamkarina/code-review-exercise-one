package service

import (
	"errors"
	"github.com/abrahamkarina/code-review-exercise-one/internal/vehicle/repository"
)

func ErrorAdapter(err error) error {
	switch {
	case errors.Is(err, repository.ErrRepositoryVehicleNotFound):
		return ErrServiceVehicleNotFound
	case errors.Is(err, repository.ErrRepositoryIdInUse):
		return ErrServiceVIdInUse
	default:
		return ErrServiceVehicleInternal
	}

}
