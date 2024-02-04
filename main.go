package main

import (
	"fmt"
	"log"
	Account "parking-system-lld/pkg/account"
	ParkingLot "parking-system-lld/pkg/parkinglot"
	Parkingspot "parking-system-lld/pkg/parkingspot"
	Vehicle "parking-system-lld/pkg/vehicle"
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
	admin1.DisplayAdminInfo()
	err := admin1.AddParkingFloor(parkinglot, "Ground Floor", 10)
	if err != nil {
		log.Fatalf("error creating parking floor: Ground Floor")
	}
	err = admin1.AddParkingFloor(parkinglot, "Level 1", 50)
	if err != nil {
		log.Fatalf("error creating parking floor: Level 1")
	}
	admin1.AddEntrancePanel(parkinglot, "1", "Ground Floor")
	admin1.AddExitPanel(parkinglot, "1", "Ground Floor")

	ParkingLot.DisplayParkingLotInfo(parkinglot)

	err = admin1.CreateHandicappedSpot(parkinglot, 0, 3)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = admin1.CreateCompactSpot(parkinglot, 0, 3)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = admin1.CreateLargeSpot(parkinglot, 0, 8)
	if err != nil {
		fmt.Println(err.Error())
	}
	// on ground floor, we are planning to open 10 spots for Large, 20 for compact, 30 for
	// motorcycles, 5 for handicapped, 2 for electric cars
	// for i := 1; i <= 10; i++ {
	// 	admin1.AddParkingSpot(parkinglot, parkinglot.ParkingFloors[parkingfloors[len(parkingfloors)-1]].Name, Parkingspot.Compact)
	// }

	attendantPerson := Account.Person{
		Name:    "attendant 1",
		Address: "Badarpur border, molarband , n-35",
		Email:   "attendant1@parking.com",
		Phone:   "+22 234 2134",
	}
	att1 := admin1.AddParkingAttendant(attendantPerson, "attendant1", "attendant1@123")
	fmt.Printf("Attendant -> %s is brought into the system \n", att1.Person.Name)
	attendants = append(attendants, att1)
	vehicle1 := Vehicle.NewVehicle(Vehicle.CAR, "DL 3CCK 2883")
	cust1 := Account.GenerateCustomer(vehicle1)

	// check if Entrance Exists
	if len(parkinglot.ParkingFloors) == 0 {
		log.Println("ParkingFloor is still in construction, no available spot for you")
	} else {
		if len(parkinglot.EntrancePanels) == 0 {
			log.Println("ParkingLot Entrance is in construction, please check in some days")
		} else {
			//check if spots exist
			var parkingspot *Parkingspot.ParkingSpot
			for idx, pf := range parkinglot.ParkingFloors {
				fmt.Println("printing displayboard of floor ", idx+1)
				pf.DisplayBoard()
				if len(pf.Spots) == 0 {
					fmt.Println("No parking spot present at the moment on floor ", pf.Name)
				} else {
					parkingspot, err = pf.FindParkingSpot(*cust1.Vehicle)
					if err != nil {
						fmt.Printf("error while finding parking spot on floor %s   error : %s \n", pf.Name, err.Error())
						continue
					}
					fmt.Println("Found parking spot ", parkingspot.Number)
					break
				}
			}
			if parkingspot != nil {
				parkinglot.EntrancePanels[0].PrintTicket(cust1.Vehicle)
			} else {
				fmt.Println("parking spot not found")
			}

		}
	}
}
