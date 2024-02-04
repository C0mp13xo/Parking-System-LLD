package Parkingspot

import (
	"fmt"
	Vehicle "parking-system-lld/pkg/vehicle"
)

type ParkingSpotType int

const (
	Handicapped ParkingSpotType = iota
	Compact
	Large
	Motorbike
	Electric
	None
)

func (psottype ParkingSpotType) String() string {
	switch psottype {
	case Handicapped:
		return "Handicapped"
	case Compact:
		return "Compact"
	case Large:
		return "Large"
	case Motorbike:
		return "Motorbike"
	case Electric:
		return "Electric"
	default:
		return "Unknown"
	}
}

type ParkingSpot struct {
	ID       string
	Occupied bool
	Number   string
	Vehicle  *Vehicle.Vehicle
	Type     ParkingSpotType
}

func (ps *ParkingSpot) IsFree() bool {
	return !ps.Occupied
}

func NewParkingSpot(pt ParkingSpotType) *ParkingSpot {
	return &ParkingSpot{
		Occupied: false,
		Type:     pt,
	}
}

func (ps *ParkingSpot) AssignVehicleToSpot(vehicle *Vehicle.Vehicle) {
	ps.Occupied = true
	ps.Vehicle = vehicle
	fmt.Printf("vehicle with License %s is parked at spot number %s", vehicle.LicenseNumber, ps.Number)
}

func (ps *ParkingSpot) FreeVehicleFromSpot(vehicle *Vehicle.Vehicle) {
	ps.Occupied = false
	ps.Vehicle = nil
}

func NewParkingSpotWithNumber(pt ParkingSpotType, number string) *ParkingSpot {
	return &ParkingSpot{
		Occupied: false,
		Type:     pt,
		Number:   number,
	}
}
