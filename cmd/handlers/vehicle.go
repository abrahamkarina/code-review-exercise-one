package handlers

import (
	"errors"
	"fmt"
	"github.com/abrahamkarina/code-review-exercise-one/cmd/web"
	"github.com/abrahamkarina/code-review-exercise-one/internal/mapper"
	"github.com/abrahamkarina/code-review-exercise-one/internal/vehicle/service"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// NewControllerVehicle returns a new instance of a vehicle controller.
func NewControllerVehicle(st service.ServiceVehicle, apdater HandlerErrorAdapter, sm mapper.StructMapper) *ControllerVehicle {
	return &ControllerVehicle{st: st,
		errAdapter: apdater,
		sm:         sm,
	}
}

// ControllerVehicle is an struct that represents a vehicle controller.
type ControllerVehicle struct {
	// StorageVehicle is the storage of vehicles.
	st         service.ServiceVehicle
	sm         mapper.StructMapper
	errAdapter HandlerErrorAdapter
}

type HandlerErrorAdapter func(error) web.ResponseError

// GetAll returns all vehicles.
func (c *ControllerVehicle) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		// ...

		// process
		vehicles, err := c.st.GetAll()
		if err != nil {
			var code int
			var body web.ResponseBodyGetAll
			switch {
			case errors.Is(err, service.ErrServiceVehicleNotFound):
				code = http.StatusNotFound
				body = web.ResponseBodyGetAll{Message: "Not found", Error: true}
			default:
				code = http.StatusInternalServerError
				body = web.ResponseBodyGetAll{Message: "Internal server error", Error: true}
			}

			ctx.JSON(code, body)
			return
		}

		// response
		code := http.StatusOK
		body := web.ResponseBodyGetAll{Message: "Success", Data: make([]*web.VehicleHandlerGetAll, 0, len(vehicles)), Error: false}
		for _, vehicle := range vehicles {
			body.Data = append(body.Data, &web.VehicleHandlerGetAll{
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
			})
		}

		ctx.JSON(code, body)
	}
}

func (c *ControllerVehicle) GetByDimensions() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		lengthParam := ctx.Query("length")
		widthParam := ctx.Query("width")

		length := strings.Split(lengthParam, "-")
		width := strings.Split(widthParam, "-")

		if len(length) != 2 || len(width) != 2 {
			httpError := web.ResponseError{
				Code:    http.StatusBadRequest,
				Message: "Wrong param format",
			}
			ctx.JSON(httpError.Code, httpError)
			return
		}

		minLength, err := strconv.ParseFloat(length[0], 64)
		if err != nil {
			httpError := web.ResponseError{
				Code:    http.StatusBadRequest,
				Message: "Wrong param format",
			}
			ctx.JSON(httpError.Code, httpError)
			return

		}
		maxLength, err := strconv.ParseFloat(length[1], 64)
		if err != nil {
			httpError := web.ResponseError{
				Code:    http.StatusBadRequest,
				Message: "Wrong param format",
			}
			ctx.JSON(httpError.Code, httpError)
			return

		}

		minWidth, err := strconv.ParseFloat(width[0], 64)
		if err != nil {
			httpError := web.ResponseError{
				Code:    http.StatusBadRequest,
				Message: "Wrong param format",
			}
			ctx.JSON(httpError.Code, httpError)
			return

		}
		maxWidth, err := strconv.ParseFloat(width[1], 64)
		if err != nil {
			httpError := web.ResponseError{
				Code:    http.StatusBadRequest,
				Message: "Wrong param format",
			}
			ctx.JSON(httpError.Code, httpError)
			return

		}

		vehicles, err := c.st.GetByDimensions(minLength, maxLength, minWidth, maxWidth)
		if err != nil {
			httpError := c.errAdapter(err)
			ctx.JSON(httpError.Code, httpError)
			return
		}
		var response web.ResponseBodyGetByDimension
		for _, v := range vehicles {
			response.Data = append(response.Data, c.sm.MapToVehicleHandlerGetByDimension(*v))
		}
		ctx.JSON(http.StatusOK, response)
		return

	}
}

func (c *ControllerVehicle) GetByWeight() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		minParam := ctx.Query("min")
		maxParam := ctx.Query("max")

		if minParam == "" || maxParam == "" {
			httpError := web.ResponseError{
				Code:    http.StatusBadRequest,
				Message: "Wrong param format",
			}
			ctx.JSON(httpError.Code, httpError)
			return
		}

		minWeight, err := strconv.ParseFloat(minParam, 64)
		if err != nil {
			httpError := web.ResponseError{
				Code:    http.StatusBadRequest,
				Message: "Wrong param format",
			}
			ctx.JSON(httpError.Code, httpError)
			return

		}
		maxWeight, err := strconv.ParseFloat(maxParam, 64)
		if err != nil {
			httpError := web.ResponseError{
				Code:    http.StatusBadRequest,
				Message: "Wrong param format",
			}
			ctx.JSON(httpError.Code, httpError)
			return

		}

		vehicles, err := c.st.GetByWeight(minWeight, maxWeight)
		if err != nil {
			httpError := c.errAdapter(err)
			ctx.JSON(httpError.Code, httpError)
			return
		}
		var response web.ResponseBodyGetByWeight
		for _, v := range vehicles {
			response.Data = append(response.Data, c.sm.MapToVehicleHandlerGetByWeight(*v))
		}
		ctx.JSON(http.StatusOK, response)
		return

	}
}

func (c *ControllerVehicle) GetByColorAndYear() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		color := ctx.Param("color")
		yearParam := ctx.Param("year")
		layout := "2006"
		_, err := time.Parse(layout, yearParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.ResponseError{
				Code:    http.StatusBadRequest,
				Message: "Bad Request: the year must be in yyyy format",
			})
			return
		}
		year, _ := strconv.Atoi(yearParam)

		vehicles, err := c.st.SearchByColorAndYear(color, year)
		if err != nil {
			httpErr := c.errAdapter(err)
			ctx.JSON(httpErr.Code, httpErr)
			return
		}
		var response web.ResponseBodyGetByYearAndColor
		for _, v := range vehicles {
			response.Data = append(response.Data, c.sm.MapFromModelVehicleHandlerGetByColorAndDate(*v))
		}
		ctx.JSON(http.StatusOK, response)

	}
}

func (c *ControllerVehicle) GetAverageCapacityByBrand() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		brand := ctx.Param("brand")
		average, err := c.st.GetAverageCapacityByBrand(brand)
		if err != nil {
			httpErr := c.errAdapter(err)
			ctx.JSON(httpErr.Code, httpErr.Message)
			return
		}
		respond := fmt.Sprintf("The average capacity for the brand %s is %f", brand, average)
		ctx.JSON(http.StatusOK, respond)
		return

	}
}
func (c *ControllerVehicle) PatchFuel() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			httpErr := web.ResponseError{Code: http.StatusBadRequest, Message: "Bad id format"}
			ctx.JSON(httpErr.Code, httpErr)
			return
		}
		var vehicle web.VehicleHandlerPatchFuel
		err = ctx.ShouldBindJSON(&vehicle)
		if err != nil {
			httpErr := web.ResponseError{Code: http.StatusBadRequest, Message: "Bad request format"}
			ctx.JSON(httpErr.Code, httpErr)
			return
		}
		if isAllowed := validateFuel(vehicle.FuelType); !isAllowed {
			httpErr := web.ResponseError{Code: http.StatusBadRequest, Message: "Not allowed fuel"}
			ctx.JSON(httpErr.Code, httpErr)
			return
		}
		err = c.st.PatchFuel(id, vehicle.FuelType)
		if err != nil {
			httpErr := c.errAdapter(err)
			ctx.JSON(httpErr.Code, httpErr)
			return
		}
		ctx.JSON(http.StatusOK, web.ResponseUpdateFuel{
			Message: "Fuel updated successfully",
		})
		return

	}
}

func (c *ControllerVehicle) Batch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var vehicles []web.VehicleHandlerPost
		err := ctx.ShouldBindJSON(&vehicles)
		if err != nil {
			httpErr := web.ResponseError{Code: http.StatusBadRequest, Message: "Bad Request"}
			ctx.JSON(httpErr.Code, httpErr)
			return
		}
		err = c.st.Batch(c.sm.MapToVehicleHandlerBatch(vehicles))
		if err != nil {
			httpErr := c.errAdapter(err)
			ctx.JSON(httpErr.Code, httpErr)
			return
		}
		ctx.JSON(http.StatusCreated, web.ResponsePost{
			Message: "Vehicles created successfully",
		})
		return
	}
}

func (c *ControllerVehicle) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var vehicle web.VehicleHandlerPost
		err := ctx.ShouldBindJSON(&vehicle)
		if err != nil {
			httpErr := web.ResponseError{Code: http.StatusBadRequest, Message: "Bad Request"}
			ctx.JSON(httpErr.Code, httpErr)
			return
		}
		err = c.st.Post(c.sm.MapToVehicleHandlerPost(vehicle))
		if err != nil {
			httpErr := c.errAdapter(err)
			ctx.JSON(httpErr.Code, httpErr)
			return
		}
		ctx.JSON(http.StatusCreated, web.ResponsePost{
			Message: "Vehicle created successfully",
		})
		return
	}
}

func (c *ControllerVehicle) PutFuel() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			httpErr := web.ResponseError{Code: http.StatusBadRequest, Message: "Bad id format"}
			ctx.JSON(httpErr.Code, httpErr)
			return
		}
		var vehicle web.VehicleHandlerPutFuel
		err = ctx.ShouldBindJSON(&vehicle)
		if err != nil {
			httpErr := web.ResponseError{Code: http.StatusBadRequest, Message: "Bad request format"}
			ctx.JSON(httpErr.Code, httpErr)
			return
		}
		if isAllowed := validateFuel(vehicle.FuelType); !isAllowed {
			httpErr := web.ResponseError{Code: http.StatusBadRequest, Message: "Not allowed fuel"}
			ctx.JSON(httpErr.Code, httpErr)
			return
		}
		err = c.st.Put(c.sm.MapFromVehicleHandlerPutFuel(id, vehicle))
		if err != nil {
			httpErr := c.errAdapter(err)
			ctx.JSON(httpErr.Code, httpErr)
			return
		}
		ctx.JSON(http.StatusOK, web.ResponseUpdateFuel{
			Message: "Fuel updated successfully",
		})
		return

	}
}

func (c *ControllerVehicle) GetByTransmission() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		transmission := ctx.Param("type")
		vehicles, err := c.st.GetByTransmission(transmission)
		if err != nil {
			httpErr := c.errAdapter(err)
			ctx.JSON(httpErr.Code, httpErr)
			return
		}
		var response web.ResponseBodyGetByTransmission
		for _, v := range vehicles {
			response.Data = append(response.Data, c.sm.MapToVehicleHandlerGetByTransmission(*v))
		}
		ctx.JSON(http.StatusOK, response)
	}
}

func (c *ControllerVehicle) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			httpErr := web.ResponseError{Code: http.StatusBadRequest, Message: "Bad id format"}
			ctx.JSON(httpErr.Code, httpErr)
			return
		}
		err = c.st.Delete(id)
		if err != nil {
			httpErr := c.errAdapter(err)
			ctx.JSON(httpErr.Code, httpErr)
			return
		}
		//TODO: NoContent doesn't show
		ctx.JSON(http.StatusNoContent, web.ResponseDelete{
			Message: "Vehicle has been successfully deleted",
		})
		return

	}

}
func validateFuel(fuel string) bool {
	allowedFuels := []string{"diesel", "biodisel", "gas", "gasoline"}
	for _, allowed := range allowedFuels {
		if strings.ToLower(fuel) == allowed {
			return true
		}
	}
	return false
}
