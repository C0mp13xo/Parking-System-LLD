package parkinglot

import (
	"fmt"
	"sync"
)

type ParkingLot struct {
	Name               string
	Address            Location
	ParkingRate        ParkingRate
	CompactSpotCount   int
	LargeSpotCount     int
	MotorbikeSpotCount int
	ElectricSpotCount  int
	MaxCompactCount    int
	MaxLargeCount      int
	MaxMotorbikeCount  int
	MaxElectricCount   int
	EntrancePanels     map[string]*EntrancePanel
	ExitPanels         map[string]*ExitPanel
	ParkingFloors      map[string]*ParkingFloor
	ActiveTickets      map[string]*ParkingTicket
	mu                 sync.Mutex // Mutex to synchronize access to shared resources
}

var parkingLot *ParkingLot
var once sync.Once

// getInstance returns the singleton instance of ParkingLot
func getInstance() *ParkingLot {
	once.Do(func() {
		parkingLot = &ParkingLot{}
		// Initialize variables, parking floors, entrance and exit panels, and active tickets from the database
		// ...
	})
	return parkingLot
}

// getNewParkingTicket issues a new parking ticket for the given vehicle
func (pl *ParkingLot) getNewParkingTicket(vehicle *Vehicle) (*ParkingTicket, error) {
	pl.mu.Lock()
	defer pl.mu.Unlock()

	if pl.isFull(vehicle.Type) {
		return nil, fmt.Errorf("Parking full for vehicle type %v", vehicle.Type)
	}

	ticket := &ParkingTicket{}
	vehicle.AssignTicket(ticket)
	ticket.SaveInDB()

	// If the ticket is successfully saved in the database, increment the parking spot count
	pl.incrementSpotCount(vehicle.Type)

	pl.ActiveTickets[ticket.TicketNumber] = ticket
	return ticket, nil
}

// isFull checks if the parking lot is full for the given vehicle type
func (pl *ParkingLot) isFull(vehicleType VehicleType) bool {
	pl.mu.Lock()
	defer pl.mu.Unlock()

	// Trucks and vans can only be parked in LargeSpot
	if vehicleType == TruckType || vehicleType == VanType {
		return pl.LargeSpotCount >= pl.MaxLargeCount
	}

	// Motorbikes can only be parked at motorbike spots
	if vehicleType == MotorbikeType {
		return pl.MotorbikeSpotCount >= pl.MaxMotorbikeCount
	}

	// Cars can be parked at compact or large spots
	if vehicleType == CarType {
		return (pl.CompactSpotCount + pl.LargeSpotCount) >= (pl.MaxCompactCount + pl.MaxLargeCount)
	}

	// Electric cars can be parked at compact, large, or electric spots
	return (pl.CompactSpotCount + pl.LargeSpotCount + pl.ElectricSpotCount) >= (pl.MaxCompactCount + pl.MaxLargeCount + pl.MaxElectricCount)
}

// incrementSpotCount increments the parking spot count based on the vehicle type
func (pl *ParkingLot) incrementSpotCount(vehicleType VehicleType) {
	switch vehicleType {
	case TruckType, VanType:
		pl.LargeSpotCount++
	case MotorbikeType:
		pl.MotorbikeSpotCount++
	case CarType:
		if pl.CompactSpotCount < pl.MaxCompactCount {
			pl.CompactSpotCount++
		} else {
			pl.LargeSpotCount++
		}
	case ElectricType:
		if pl.ElectricSpotCount < pl.MaxElectricCount {
			pl.ElectricSpotCount++
		} else if pl.CompactSpotCount < pl.MaxCompactCount {
			pl.CompactSpotCount++
		} else {
			pl.LargeSpotCount++
		}
	}
}

// isFull checks if the entire parking lot is full
func (pl *ParkingLot) isFull() bool {
	pl.mu.Lock()
	defer pl.mu.Unlock()

	for key := range pl.ParkingFloors {
		if !pl.ParkingFloors[key].IsFull() {
			return false
		}
	}
	return true
}

// addParkingFloor adds a parking floor to the parking lot
func (pl *ParkingLot) addParkingFloor(floor *ParkingFloor) {
	pl.mu.Lock()
	defer pl.mu.Unlock()

	// Store in the database or perform any other required actions
	// ...
	pl.ParkingFloors[floor.Name] = floor
}

// addEntrancePanel adds an entrance panel to the parking lot
func (pl *ParkingLot) addEntrancePanel(entrancePanel *EntrancePanel) {
	pl.mu.Lock()
	defer pl.mu.Unlock()

	// Store in the database or perform any other required actions
	// ...
	pl.EntrancePanels[entrancePanel.ID] = entrancePanel
}

// addExitPanel adds an exit panel to the parking lot
func (pl *ParkingLot) addExitPanel(exitPanel *ExitPanel) {
	pl.mu.Lock()
	defer pl.mu.Unlock()

	// Store in the database or perform any other required actions
	// ...
	pl.ExitPanels[exitPanel.ID] = exitPanel
}
