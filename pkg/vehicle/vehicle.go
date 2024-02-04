package Vehicle

import "fmt"

type VehicleType int

const (
	CAR VehicleType = iota
	VAN
	TRUCK
	MOTORCYCLE
	ELECTRIC
)

func (status VehicleType) String() string {
	switch status {
	case CAR:
		return "Car"
	case VAN:
		return "Van"
	case TRUCK:
		return "Truck"
	case MOTORCYCLE:
		return "Motorcycle"
	case ELECTRIC:
		return "Electric"
	default:
		return "Unknown"
	}
}

type Vehicle struct {
	LicenseNumber string
	Type          VehicleType
}

func NewVehicle(vt VehicleType, licNumber string) *Vehicle {
	fmt.Printf("Vehicle of type %s has come to our parking \n", vt.String())
	return &Vehicle{
		Type:          vt,
		LicenseNumber: licNumber,
	}
}
