package account

import (
	"fmt"
	"parking-system-lld/pkg/parkingfloor"
	ParkingLot "parking-system-lld/pkg/parkinglot"
	Parkingspot "parking-system-lld/pkg/parkingspot"
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

func (admin *Admin) AddParkingFloor(parkinglot *ParkingLot.ParkingLot, pfName string, cap int) (string, error) {
	// Implement logic to add parking floor
	if len(pfName) == 0 {
		return "", fmt.Errorf("error : parking floor name provided is empty")
	}
	pf := parkingfloor.NewParkingFloor(pfName, cap)
	err := parkinglot.AddParkingFloor(pf)
	if err != nil {
		return "", err
	}
	return pfName, nil
}

func (admin *Admin) AddParkingSpot(parkinglot *ParkingLot.ParkingLot, pfName string, spottype Parkingspot.ParkingSpotType) bool {
	// Implement logic to add parking spot
	parkinglot.ParkingFloors[pfName].AddParkingSpot(spottype)
	return true
}

func (admin *Admin) AddEntrancePanel(parkinglot *ParkingLot.ParkingLot, id, name string) {
	parkinglot.AddEntrancePanel(id, name)
}

func (admin *Admin) AddExitPanel(parkinglot *ParkingLot.ParkingLot, id, name string) {
	parkinglot.AddExitPanel(id, name)
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
