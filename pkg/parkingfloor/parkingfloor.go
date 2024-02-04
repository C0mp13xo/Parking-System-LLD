package parkingfloor

import (
	"fmt"
	Parkingspot "parking-system-lld/pkg/parkingspot"
	"parking-system-lld/pkg/parkingticket"
	Vehicle "parking-system-lld/pkg/vehicle"
	"strconv"
)

type ParkingFloor struct {
	Name                string
	Spots               []*Parkingspot.ParkingSpot
	Capacity            int
	CustomersInfoPortal []*parkingticket.ParkingTicket
}

func NewParkingFloor(name string, capacity int) *ParkingFloor {
	return &ParkingFloor{
		Name:                name,
		Spots:               make([]*Parkingspot.ParkingSpot, 0),
		Capacity:            capacity,
		CustomersInfoPortal: make([]*parkingticket.ParkingTicket, 0),
	}
}

func (pf *ParkingFloor) AddParkingSpot(spot Parkingspot.ParkingSpotType) (bool, error) {
	n := len(pf.Spots)
	if n == pf.Capacity {
		return false, fmt.Errorf("capacity of %s has reached", pf.Name)
	}
	spotNumber := strconv.Itoa(n + 1)
	switch spot {
	case Parkingspot.Handicapped:
		pf.Spots = append(pf.Spots, Parkingspot.NewParkingSpotWithNumber(Parkingspot.Handicapped, spotNumber))
	case Parkingspot.Compact:
		pf.Spots = append(pf.Spots, Parkingspot.NewParkingSpotWithNumber(Parkingspot.Compact, spotNumber))
	case Parkingspot.Large:
		pf.Spots = append(pf.Spots, Parkingspot.NewParkingSpotWithNumber(Parkingspot.Large, spotNumber))
	case Parkingspot.Motorbike:
		pf.Spots = append(pf.Spots, Parkingspot.NewParkingSpotWithNumber(Parkingspot.Motorbike, spotNumber))
	case Parkingspot.Electric:
		pf.Spots = append(pf.Spots, Parkingspot.NewParkingSpotWithNumber(Parkingspot.Electric, spotNumber))
	default:
		fmt.Println("Wrong parking spot type!")
	}
	return true, nil
}

func (pf *ParkingFloor) FindParkingSpot(v Vehicle.Vehicle) (*Parkingspot.ParkingSpot, error) {
	for _, s := range pf.Spots {
		if s.Type == Parkingspot.ParkingSpotType(v.Type) && !s.Occupied {
			return s, nil
		}
	}
	return nil, fmt.Errorf("spot not Found")
}

func (pf *ParkingFloor) AssignVehicleToSpot(vehicle *Vehicle.Vehicle, spot *Parkingspot.ParkingSpot) {
	spot.AssignVehicleToSpot(vehicle)
}

func (pf *ParkingFloor) DisplayBoard() {
	fmt.Println(" **************** welcome to floor ", pf.Name, " ****************")
	var largespot *Parkingspot.ParkingSpot
	var motorspot *Parkingspot.ParkingSpot
	var electricspot *Parkingspot.ParkingSpot
	var compactspot *Parkingspot.ParkingSpot
	var handicappedspot *Parkingspot.ParkingSpot
	for _, spot := range pf.Spots {
		if spot.Type == Parkingspot.Large && spot.IsFree() {
			largespot = spot
			break
		}
	}

	for _, spot := range pf.Spots {
		if spot.Type == Parkingspot.Compact && spot.IsFree() {
			compactspot = spot
			break
		}
	}

	for _, spot := range pf.Spots {
		if spot.Type == Parkingspot.Motorbike && spot.IsFree() {
			motorspot = spot
			break
		}
	}

	for _, spot := range pf.Spots {
		if spot.Type == Parkingspot.Handicapped && spot.IsFree() {
			handicappedspot = spot
			break
		}
	}

	for _, spot := range pf.Spots {
		if spot.Type == Parkingspot.Electric && spot.IsFree() {
			electricspot = spot
			break
		}
	}
	if largespot != nil {
		fmt.Println(" Large spot available at ", largespot.Number)
	} else {
		fmt.Println(" Large spot are full at this moment ")
	}

	if motorspot != nil {
		fmt.Println(" Motor spot available at ", motorspot.Number)
	} else {
		fmt.Println(" Motor spot not available at this moment ")
	}

	if compactspot != nil {
		fmt.Println(" Compact spot available at ", compactspot.Number)
	} else {
		fmt.Println(" Compact spot not available at this moment ")
	}

	if handicappedspot != nil {
		fmt.Println(" Handicapped spot available at ", handicappedspot.Number)
	} else {
		fmt.Println(" Handicapped spot not available at this moment ")
	}

	if electricspot != nil {
		fmt.Println(" Electric spot available at ", electricspot.Number)
	} else {
		fmt.Println(" Electric spot not available at this moment ")
	}
}

//method to pay at customer info portal
func (pf *ParkingFloor) PayAtCustomerInfoPortal(ticket *parkingticket.ParkingTicket) {

}
