package main

import (
	"log"
	Account "parking-system-lld/pkg/account"
	ParkingLot "parking-system-lld/pkg/parkinglot"
	Parkingspot "parking-system-lld/pkg/parkingspot"
)

var attendants []*Account.ParkingAttendant

func main() {
	location := ParkingLot.Location{
		State:        "Delhi",
		City:         "Badarpur",
		ZipCode:      "110044",
		StreetAdress: "E-22, Molarbad",
		Country:      "India",
	}
	parkinglot := ParkingLot.CreateInstance(location, "ABC Parking")
	//creating admin
	p := Account.Person{
		Name:    "admin",
		Address: "North California, Dallas Street, H. No - 3332",
		Email:   "admin@parking.com",
		Phone:   "+22 45678 2134",
	}
	admin1 := Account.AddAdmin(p, "Admin", "Admin@123")

	parkinglot.AddEntrancePanel("1", "Ground Floor")
	parkinglot.AddExitPanel("1", "Ground Floor")
	parkingfloors := []string{}
	pfname, err := admin1.AddParkingFloor(parkinglot, "Ground Floor", 100)
	if err != nil {
		log.Fatalf("error creating parking floor")
	}
	parkingfloors = append(parkingfloors, pfname)
	admin1.AddParkingFloor(parkinglot, "Level 1", 50)
	ParkingLot.DisplayParkingLotInfo(parkinglot)

	// on ground floor, we are planning to open 10 spots for Large, 20 for compact, 30 for
	// motorcycles, 5 for handicapped, 2 for electric cars
	for i := 1; i <= 10; i++ {
		admin1.AddParkingSpot(parkinglot, parkinglot.ParkingFloors[parkingfloors[len(parkingfloors)-1]].Name, Parkingspot.Compact)
	}

	attendantPerson := Account.Person{
		Name:    "attendant 1",
		Address: "Badarpur border, molarband , n-35",
		Email:   "attendant1@parking.com",
		Phone:   "+22 234 2134",
	}
	att1 := admin1.AddParkingAttendant(attendantPerson, "attendant1", "attendant1@123")
	attendants = append(attendants, att1)

}
