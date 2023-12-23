package Parkingspot

import Vehicle "parking-system-lld/pkg/vehicle"

type ParkingSpotType int

const (
	Handicapped ParkingSpotType = iota
	Compact
	Large
	Motorbike
	Electric
)

type ParkingSpot struct {
	Number  string
	Free    bool
	Vehicle Vehicle.Vehicle
	Type    ParkingSpotType
}

func (ps *ParkingSpot) IsFree() bool {
	return ps.Free
}

func NewParkingSpot(pt ParkingSpotType) *ParkingSpot {
	return &ParkingSpot{
		Free: true,
		Type: pt,
	}
}

func (ps *ParkingSpot) AssignVehicle(vehicle Vehicle.Vehicle) {
	ps.Vehicle = vehicle
	ps.Free = false
}

func (ps *ParkingSpot) RemoveVehicle() {
	ps.Vehicle = nil
	ps.Free = true
}

type HandicappedSpot struct {
	ParkingSpot
}

func NewHandicappedSpot() *HandicappedSpot {
	return &HandicappedSpot{
		ParkingSpot: *NewParkingSpot(Handicapped),
	}
}

type CompactSpot struct {
	ParkingSpot
}

func NewCompactSpot() *CompactSpot {
	return &CompactSpot{
		ParkingSpot: *NewParkingSpot(Compact),
	}
}

type LargeSpot struct {
	ParkingSpot
}

func NewLargeSpot() *LargeSpot {
	return &LargeSpot{
		ParkingSpot: *NewParkingSpot(Large),
	}
}

type MotorbikeSpot struct {
	ParkingSpot
}

func NewMotorbikeSpot() *MotorbikeSpot {
	return &MotorbikeSpot{
		ParkingSpot: *NewParkingSpot(Motorbike),
	}
}

type ElectricSpot struct {
	ParkingSpot
}

func NewElectricSpot() *ElectricSpot {
	return &ElectricSpot{
		ParkingSpot: *NewParkingSpot(Electric),
	}
}
