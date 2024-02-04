package parkinglot

import (
	"fmt"
	"parking-system-lld/pkg/account"
	"parking-system-lld/pkg/parkingfloor"
	"parking-system-lld/pkg/parkingticket"
	Vehicle "parking-system-lld/pkg/vehicle"
)

type EntrancePanel struct {
	ID   string
	Name string
}

type ExitPanel struct {
	ID   string
	Name string
}

type Location struct {
	State        string
	City         string
	ZipCode      string
	StreetAdress string
	Country      string
}

type ParkingLot struct {
	Name    string
	Address Location
	// ParkingRate        ParkingRate
	// CompactSpotCount   int
	// LargeSpotCount     int
	// MotorbikeSpotCount int
	// ElectricSpotCount  int
	// MaxCompactCount    int
	// MaxLargeCount      int
	// MaxMotorbikeCount  int
	// MaxElectricCount   int
	EntrancePanels map[string]*EntrancePanel
	ExitPanels     map[string]*ExitPanel
	ParkingFloors  map[string]*parkingfloor.ParkingFloor
	// ActiveTickets      map[string]*ParkingTicket
}

// getInstance returns the singleton instance of ParkingLot
func CreateInstance(l Location, name string) *ParkingLot {

	parkingLot := &ParkingLot{
		Address: l,
		Name:    name,
	}

	return parkingLot
}

// getNewParkingTicket issues a new parking ticket for the given vehicle
// func (pl *ParkingLot) getNewParkingTicket(vehicle *Vehicle) (*ParkingTicket, error) {
// 	pl.mu.Lock()
// 	defer pl.mu.Unlock()

// 	if pl.isFull(vehicle.Type) {
// 		return nil, fmt.Errorf("Parking full for vehicle type %v", vehicle.Type)
// 	}

// 	ticket := &ParkingTicket{}
// 	vehicle.AssignTicket(ticket)
// 	ticket.SaveInDB()

// 	// If the ticket is successfully saved in the database, increment the parking spot count
// 	pl.incrementSpotCount(vehicle.Type)

// 	pl.ActiveTickets[ticket.TicketNumber] = ticket
// 	return ticket, nil
// }

// isFull checks if the parking lot is full for the given vehicle type
// func (pl *ParkingLot) isFull(vehicleType VehicleType) bool {
// 	pl.mu.Lock()
// 	defer pl.mu.Unlock()

// 	// Trucks and vans can only be parked in LargeSpot
// 	if vehicleType == TruckType || vehicleType == VanType {
// 		return pl.LargeSpotCount >= pl.MaxLargeCount
// 	}

// 	// Motorbikes can only be parked at motorbike spots
// 	if vehicleType == MotorbikeType {
// 		return pl.MotorbikeSpotCount >= pl.MaxMotorbikeCount
// 	}

// 	// Cars can be parked at compact or large spots
// 	if vehicleType == CarType {
// 		return (pl.CompactSpotCount + pl.LargeSpotCount) >= (pl.MaxCompactCount + pl.MaxLargeCount)
// 	}

// 	// Electric cars can be parked at compact, large, or electric spots
// 	return (pl.CompactSpotCount + pl.LargeSpotCount + pl.ElectricSpotCount) >= (pl.MaxCompactCount + pl.MaxLargeCount + pl.MaxElectricCount)
// }

// incrementSpotCount increments the parking spot count based on the vehicle type
// func (pl *ParkingLot) incrementSpotCount(vehicleType VehicleType) {
// 	switch vehicleType {
// 	case TruckType, VanType:
// 		pl.LargeSpotCount++
// 	case MotorbikeType:
// 		pl.MotorbikeSpotCount++
// 	case CarType:
// 		if pl.CompactSpotCount < pl.MaxCompactCount {
// 			pl.CompactSpotCount++
// 		} else {
// 			pl.LargeSpotCount++
// 		}
// 	case ElectricType:
// 		if pl.ElectricSpotCount < pl.MaxElectricCount {
// 			pl.ElectricSpotCount++
// 		} else if pl.CompactSpotCount < pl.MaxCompactCount {
// 			pl.CompactSpotCount++
// 		} else {
// 			pl.LargeSpotCount++
// 		}
// 	}
// }

// isFull checks if the entire parking lot is full
// func (pl *ParkingLot) isFull() bool {
// 	pl.mu.Lock()
// 	defer pl.mu.Unlock()

// 	for key := range pl.ParkingFloors {
// 		if !pl.ParkingFloors[key].IsFull() {
// 			return false
// 		}
// 	}
// 	return true
// }

// addParkingFloor adds a parking floor to the parking lot
func (pl *ParkingLot) AddParkingFloor(floor *parkingfloor.ParkingFloor) error {
	// pl.mu.Lock()
	// defer pl.mu.Unlock()

	pl.ParkingFloors[floor.Name] = floor
	return nil
}

// addEntrancePanel adds an entrance panel to the parking lot
func (pl *ParkingLot) AddEntrancePanel(id, name string) {
	// pl.mu.Lock()
	// defer pl.mu.Unlock()
	entrancePanel := EntrancePanel{
		ID:   id,
		Name: name,
	}
	// Store in the database or perform any other required actions
	// ...
	pl.EntrancePanels[entrancePanel.ID] = &entrancePanel
}

// addExitPanel adds an exit panel to the parking lot
func (pl *ParkingLot) AddExitPanel(id, name string) {
	// pl.mu.Lock()
	// defer pl.mu.Unlock()
	exitPanel := ExitPanel{
		ID:   id,
		Name: name,
	}

	pl.ExitPanels[exitPanel.ID] = &exitPanel
}

func (xp *ExitPanel) ProcessPayment(ticket *parkingticket.ParkingTicket, payment *Payment) {
	// Process payment for a parking ticket.
}

func (ep *EntrancePanel) PrintTicket( vehicle *Vehicle.Vehicle) *parkingticket.ParkingTicket {
	ticket := generateTicket(vehicle)
	return nil
}

func DisplayParkingLotInfo(prklot *ParkingLot) {
	fmt.Println("************** Welcome to Parking Lot ", prklot.Name, " ***************")
	fmt.Println("----------------------------------------------------------------------------")
	fmt.Println(" No of Entrance Panels ", len(prklot.EntrancePanels))
	fmt.Println(" No of Exit Panels ", len(prklot.ExitPanels))
	fmt.Println(" No of Parking Floors ", len(prklot.ParkingFloors))
}


func generateTicket(v *Vehicle.Vehicle) *parkingticket.ParkingTicket{
	parkingticket.ParkingTicket.
}