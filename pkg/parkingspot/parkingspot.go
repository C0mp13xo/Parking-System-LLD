package Parkingspot

import (
	Vehicle "parking-system-lld/pkg/vehicle"
)

type ParkingSpotType int

const (
	Handicapped ParkingSpotType = iota
	Compact
	Large
	Motorbike
	Electric
)

type ParkingSpot struct {
	ID       string
	Occupied bool
	Number   string
	Vehicle  *Vehicle.Vehicle
	Type     ParkingSpotType
}

func (ps *ParkingSpot) IsFree() bool {
	return ps.Occupied
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
