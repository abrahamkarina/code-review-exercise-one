package domain

// VehicleAttributes is an struct that represents the attributes of a vehicle.
type VehicleAttributes struct {
	// Brand is the brand of the vehicle.
	Brand 		 string
	// Model is the model of the vehicle.
	Model 		 string
	// Registration is the registration of the vehicle.
	Registration string
	// Year is the fabrication year of the vehicle.
	Year 		 int
	// Color is the color of the vehicle.
	Color 		 string

	// MaxSpeed is the maximum speed of the vehicle.
	MaxSpeed 	 int
	// FuelType is the fuel type of the vehicle.
	FuelType 	 string
	// Transmission is the transmission of the vehicle.
	Transmission string

	// Passengers is the capacity of passengers of the vehicle.
	Passengers 	 int

	// Height is the height of the vehicle.
	Height 		 float64
	// Width is the width of the vehicle.
	Width 		 float64

	// Weight is the weight of the vehicle.
	Weight 		 float64
}

// Vehicle is an struct that represents a vehicle.
type Vehicle struct {
	// ID is the unique identifier of the vehicle.
	Id 			 int
	
	// Attributes is the attributes of the vehicle.
	Attributes 	 VehicleAttributes
}