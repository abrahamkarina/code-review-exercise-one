package http_error

import (
	"errors"
	"github.com/abrahamkarina/code-review-exercise-one/cmd/web"
	"github.com/abrahamkarina/code-review-exercise-one/internal/vehicle/service"
	"net/http"
)

func ErrorAdapter(err error) web.ResponseError {
	switch {
	case errors.Is(err, service.ErrServiceVehicleNotFound):
		return web.ResponseError{
			Code:    http.StatusNotFound,
			Message: "Not Found: Vehicle not found",
		}
	case errors.Is(err, service.ErrServiceVIdInUse):
		return web.ResponseError{
			Code:    http.StatusNotFound,
			Message: "Id already in use",
		}

	default:
		return web.ResponseError{
			Code:    http.StatusInternalServerError,
			Message: "internal server error",
		}
	}

}
