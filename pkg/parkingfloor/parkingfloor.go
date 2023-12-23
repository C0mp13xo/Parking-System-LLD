package parkingfloor

import (
	"fmt"
)

type ParkingFloor struct {
	Name             string
	HandicappedSpots map[string]*HandicappedSpot
	CompactSpots     map[string]*CompactSpot
	LargeSpots       map[string]*LargeSpot
	MotorbikeSpots   map[string]*MotorbikeSpot
	ElectricSpots    map[string]*ElectricSpot
	InfoPortals      map[string]*CustomerInfoPortal
	DisplayBoard     *ParkingDisplayBoard
}

func NewParkingFloor(name string) *ParkingFloor {
	return &ParkingFloor{
		Name:             name,
		HandicappedSpots: make(map[string]*HandicappedSpot),
		CompactSpots:     make(map[string]*CompactSpot),
		LargeSpots:       make(map[string]*LargeSpot),
		MotorbikeSpots:   make(map[string]*MotorbikeSpot),
		ElectricSpots:    make(map[string]*ElectricSpot),
		InfoPortals:      make(map[string]*CustomerInfoPortal),
		DisplayBoard:     NewParkingDisplayBoard(),
	}
}

func (pf *ParkingFloor) AddParkingSpot(spot ParkingSpot) {
	switch spot.Type {
	case Handicapped:
		pf.HandicappedSpots[spot.Number] = spot.(*HandicappedSpot)
	case Compact:
		pf.CompactSpots[spot.Number] = spot.(*CompactSpot)
	case Large:
		pf.LargeSpots[spot.Number] = spot.(*LargeSpot)
	case Motorbike:
		pf.MotorbikeSpots[spot.Number] = spot.(*MotorbikeSpot)
	case Electric:
		pf.ElectricSpots[spot.Number] = spot.(*ElectricSpot)
	default:
		fmt.Println("Wrong parking spot type!")
	}
}

func (pf *ParkingFloor) AssignVehicleToSpot(vehicle *Vehicle, spot ParkingSpot) {
	spot.AssignVehicle(vehicle)
	switch spot.Type {
	case Handicapped:
		pf.UpdateDisplayBoardForHandicapped(spot)
	case Compact:
		pf.UpdateDisplayBoardForCompact(spot)
	case Large:
		pf.UpdateDisplayBoardForLarge(spot)
	case Motorbike:
		pf.UpdateDisplayBoardForMotorbike(spot)
	case Electric:
		pf.UpdateDisplayBoardForElectric(spot)
	default:
		fmt.Println("Wrong parking spot type!")
	}
}

func (pf *ParkingFloor) UpdateDisplayBoardForHandicapped(spot ParkingSpot) {
	if pf.DisplayBoard.HandicappedFreeSpot.Number == spot.Number {
		// Find another free handicapped parking and assign it to the displayBoard
		for key := range pf.HandicappedSpots {
			if pf.HandicappedSpots[key].IsFree() {
				pf.DisplayBoard.HandicappedFreeSpot = pf.HandicappedSpots[key]
			}
		}
		pf.DisplayBoard.ShowEmptySpotNumber()
	}
}

func (pf *ParkingFloor) UpdateDisplayBoardForCompact(spot ParkingSpot) {
	if pf.DisplayBoard.CompactFreeSpot.Number == spot.Number {
		// Find another free compact parking and assign it to the displayBoard
		for key := range pf.CompactSpots {
			if pf.CompactSpots[key].IsFree() {
				pf.DisplayBoard.CompactFreeSpot = pf.CompactSpots[key]
			}
		}
		pf.DisplayBoard.ShowEmptySpotNumber()
	}
}

func (pf *ParkingFloor) FreeSpot(spot ParkingSpot) {
	spot.RemoveVehicle()
	switch spot.Type {
	case Handicapped:
		// Increment freeHandicappedSpotCount or handle counts as needed
	case Compact:
		// Increment freeCompactSpotCount or handle counts as needed
	case Large:
		// Increment freeLargeSpotCount or handle counts as needed
	case Motorbike:
		// Increment freeMotorbikeSpotCount or handle counts as needed
	case Electric:
		// Increment freeElectricSpotCount or handle counts as needed
	default:
		fmt.Println("Wrong parking spot type!")
	}
}
