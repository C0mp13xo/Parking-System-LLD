package account

import (
	"fmt"
	ParkingLot "parking-system-lld/pkg/parkinglot"
	Parkingspot "parking-system-lld/pkg/parkingspot"
	"parking-system-lld/pkg/parkingticket"
	Vehicle "parking-system-lld/pkg/vehicle"
)

type AccountStatus int

const (
	ACTIVE AccountStatus = iota
	CANCELLED
	CLOSED
	BLACKLISTED
	NONE
)

type AccountType int

const (
	ADMIN AccountType = iota
	ATTENDANT
	CUSTOMER
)

type Person struct {
	Name    string
	Address string
	Email   string
	Phone   string
}

type Account struct {
	UserName string
	Password string
	Status   AccountStatus
	Person   Person
	AccountType
}

func (a *Account) ResetPassword() bool {
	// Implement resetPassword logic here
	return true
}

type Admin struct {
	Account
}

func AddAdmin(person Person, username, password string) *Admin {
	return &Admin{
		Account: Account{
			Person:      person,
			UserName:    username,
			Password:    password,
			AccountType: ADMIN,
		},
	}
}

func (admin *Admin) DisplayAdminInfo() {
	fmt.Printf(" ---- Admin %s was added to the system \n", admin.Person.Name)

}

func (admin *Admin) AddParkingFloor(parkinglot *ParkingLot.ParkingLot, pfName string, cap int) error {
	// Implement logic to add parking floor
	if len(pfName) == 0 {
		return fmt.Errorf("error : parking floor name provided is empty ")
	}

	err := parkinglot.AddParkingFloor(pfName, cap)
	if err != nil {
		return err
	}
	fmt.Printf("admin %s has added Parking Floor %s \n", admin.Person.Name, pfName)
	return nil
}

func (admin *Admin) AddEntrancePanel(parkinglot *ParkingLot.ParkingLot, id, name string) {
	parkinglot.AddEntrancePanel(id, name)
	fmt.Printf("admin %s has added Entrance Panel %s \n", admin.Person.Name, name)
}

func (admin *Admin) AddExitPanel(parkinglot *ParkingLot.ParkingLot, id, name string) {
	parkinglot.AddExitPanel(id, name)
	fmt.Printf("admin %s has added Exit Panel %s \n", admin.Person.Name, name)
}

func (admin *Admin) CreateHandicappedSpot(p *ParkingLot.ParkingLot, floorIdx int, count int) error {
	for i := 1; i <= count; i++ {
		err := admin.createParkingSpot(p, floorIdx, Parkingspot.Handicapped)
		if err != nil {
			return err
		}
	}
	return nil
}

func (admin *Admin) CreateLargeSpot(p *ParkingLot.ParkingLot, floorIdx int, count int) error {
	for i := 1; i <= count; i++ {
		err := admin.createParkingSpot(p, floorIdx, Parkingspot.Large)
		if err != nil {
			return err
		}
	}
	return nil
}

func (admin *Admin) CreateCompactSpot(p *ParkingLot.ParkingLot, floorIdx int, count int) error {
	for i := 1; i <= count; i++ {
		err := admin.createParkingSpot(p, floorIdx, Parkingspot.Compact)
		if err != nil {
			return err
		}
	}
	return nil
}

func (admin *Admin) CreateElectricSpot(p *ParkingLot.ParkingLot, floorIdx int, count int) error {
	for i := 1; i <= count; i++ {
		err := admin.createParkingSpot(p, floorIdx, Parkingspot.Electric)
		if err != nil {
			return err
		}
	}
	return nil
}

func (admin *Admin) CreateMotorbikeSpot(p *ParkingLot.ParkingLot, floorIdx int, count int) error {
	for i := 1; i <= count; i++ {
		err := admin.createParkingSpot(p, floorIdx, Parkingspot.Motorbike)
		if err != nil {
			return err
		}
	}
	return nil
}

func (admin *Admin) createParkingSpot(parkinglot *ParkingLot.ParkingLot, floorIdx int, parkingspottype Parkingspot.ParkingSpotType) error {
	err := parkinglot.ParkingFloors[floorIdx].AddParkingSpot(parkingspottype)
	return err
}

type ParkingAttendant struct {
	Account
}

func (admin *Admin) AddParkingAttendant(person Person, username, password string) *ParkingAttendant {
	return &ParkingAttendant{
		Account: Account{
			Person:      person,
			UserName:    username,
			Password:    password,
			AccountType: ATTENDANT,
		},
	}
}

func (attendant *ParkingAttendant) ProcessTicket(ticketNumber string) bool {
	// Implement logic to process ticket
	return true
}

type Customer struct {
	Ticket  *parkingticket.ParkingTicket
	Vehicle *Vehicle.Vehicle
}

func GenerateCustomer(v *Vehicle.Vehicle) *Customer {
	return &Customer{
		Vehicle: v,
	}
}

func (cust *Customer) AssignTicket(tkt *parkingticket.ParkingTicket) {
	cust.Ticket = tkt
}
