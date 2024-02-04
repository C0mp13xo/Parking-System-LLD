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
	IsFull              bool
}

func NewParkingFloor(name string, capacity int) *ParkingFloor {
	return &ParkingFloor{
		Name:                name,
		Spots:               make([]*Parkingspot.ParkingSpot, 0),
		Capacity:            capacity,
		CustomersInfoPortal: make([]*parkingticket.ParkingTicket, 0),
	}
}

func (pf *ParkingFloor) AddParkingSpot(spot Parkingspot.ParkingSpotType) error {
	n := len(pf.Spots)
	if n == pf.Capacity {
		return fmt.Errorf("capacity of %s has reached", pf.Name)
	}
	spotNumber := strconv.Itoa(n + 1)
	switch spot {
	case Parkingspot.Handicapped:
		pf.Spots = append(pf.Spots, Parkingspot.NewParkingSpotWithNumber(Parkingspot.Handicapped, spotNumber))
		fmt.Printf("Parking spot of type %s with spot number %s was added \n", Parkingspot.Handicapped.String(), spotNumber)
	case Parkingspot.Compact:
		pf.Spots = append(pf.Spots, Parkingspot.NewParkingSpotWithNumber(Parkingspot.Compact, spotNumber))
		fmt.Printf("Parking spot of type %s with spot number %s was added \n", Parkingspot.Compact.String(), spotNumber)
	case Parkingspot.Large:
		pf.Spots = append(pf.Spots, Parkingspot.NewParkingSpotWithNumber(Parkingspot.Large, spotNumber))
		fmt.Printf("Parking spot of type %s with spot number %s was added \n", Parkingspot.Large.String(), spotNumber)
	case Parkingspot.Motorbike:
		pf.Spots = append(pf.Spots, Parkingspot.NewParkingSpotWithNumber(Parkingspot.Motorbike, spotNumber))
		fmt.Printf("Parking spot of type %s with spot number %s was added \n", Parkingspot.Motorbike.String(), spotNumber)
	case Parkingspot.Electric:
		pf.Spots = append(pf.Spots, Parkingspot.NewParkingSpotWithNumber(Parkingspot.Electric, spotNumber))
		fmt.Printf("Parking spot of type %s with spot number %s was added \n", Parkingspot.Electric.String(), spotNumber)
	default:
		return fmt.Errorf("invalid spot type provided")
	}
	return nil
}

func (pf *ParkingFloor) FindParkingSpot(v Vehicle.Vehicle) (*Parkingspot.ParkingSpot, error) {
	sptType, err := findSpotTypeForVehicle(v.Type)
	if err != nil {
		return nil, fmt.Errorf("we do not accept this vehicle type : %s", v.Type.String())
	}
	for _, s := range pf.Spots {
		if s.Type == Parkingspot.ParkingSpotType(sptType) && !s.Occupied {
			return s, nil
		}
	}
	return nil, fmt.Errorf("spot not Found for vehicle type : %s ", v.Type.String())
}

func findSpotTypeForVehicle(v Vehicle.VehicleType) (Parkingspot.ParkingSpotType, error) {
	switch v {
	case Vehicle.CAR:
		return Parkingspot.Compact, nil
	case Vehicle.ELECTRIC:
		return Parkingspot.Electric, nil
	case Vehicle.TRUCK:
		return Parkingspot.Large, nil
	case Vehicle.VAN:
		return Parkingspot.Large, nil
	case Vehicle.MOTORCYCLE:
		return Parkingspot.Motorbike, nil
	default:
		return Parkingspot.None, fmt.Errorf("invalid Vehicle type")
	}
}

func (pf *ParkingFloor) AssignVehicleToSpot(vehicle *Vehicle.Vehicle, spot *Parkingspot.ParkingSpot) {
	spot.AssignVehicleToSpot(vehicle)
}

func (pf *ParkingFloor) DisplayBoard() {
	lrCount, cmCount, moCount, elCount, hanCount := 0, 0, 0, 0, 0
	for _, ps := range pf.Spots {
		switch ps.Type {
		case Parkingspot.Compact:
			if ps.IsFree() {
				cmCount++
			}
		case Parkingspot.Electric:
			if ps.IsFree() {
				elCount++
			}
		case Parkingspot.Large:
			if ps.IsFree() {
				lrCount++
			}
		case Parkingspot.Handicapped:
			if ps.IsFree() {
				hanCount++
			}
		case Parkingspot.Motorbike:
			if ps.IsFree() {
				moCount++
			}
		default:
		}
	}
	fmt.Println(" **************** welcome to floor ", pf.Name, " ****************")
	fmt.Println(" Total spots present at this floor ", len(pf.Spots))
	fmt.Println(" Total Available spots for Large Vehicle ", lrCount)
	fmt.Println(" Total Available spots for Compact Vehicle ", cmCount)
	fmt.Println(" Total Available spots for Motorbike Vehicle ", moCount)
	fmt.Println(" Total Available spots for Electric Vehicle ", elCount)
	fmt.Println(" Total Available spots for Handicapped  ", hanCount)
	fmt.Println("*******************************************************************")

}

// method to pay at customer info portal
func (pf *ParkingFloor) PayAtCustomerInfoPortal(ticket *parkingticket.ParkingTicket) {

}
