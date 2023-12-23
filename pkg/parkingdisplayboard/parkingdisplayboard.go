package parkingdisplayboard

import (
	"fmt"
	Parkingspot "parking-system-lld/pkg/parkingspot"
)

type ParkingDisplayBoard struct {
	ID                  string
	HandicappedFreeSpot *Parkingspot.HandicappedSpot
	CompactFreeSpot     *Parkingspot.CompactSpot
	LargeFreeSpot       *Parkingspot.LargeSpot
	MotorbikeFreeSpot   *Parkingspot.MotorbikeSpot
	ElectricFreeSpot    *Parkingspot.ElectricSpot
}

func (pdb *ParkingDisplayBoard) ShowEmptySpotNumber() {
	message := ""

	if pdb.HandicappedFreeSpot.IsFree() {
		message += fmt.Sprintf("Free Handicapped: %s", pdb.HandicappedFreeSpot.Number)
	} else {
		message += "Handicapped is full"
	}
	message += "\n"

	if pdb.CompactFreeSpot.IsFree() {
		message += fmt.Sprintf("Free Compact: %s", pdb.CompactFreeSpot.Number)
	} else {
		message += "Compact is full"
	}
	message += "\n"

	if pdb.LargeFreeSpot.IsFree() {
		message += fmt.Sprintf("Free Large: %s", pdb.LargeFreeSpot.Number)
	} else {
		message += "Large is full"
	}
	message += "\n"

	if pdb.MotorbikeFreeSpot.IsFree() {
		message += fmt.Sprintf("Free Motorbike: %s", pdb.MotorbikeFreeSpot.Number)
	} else {
		message += "Motorbike is full"
	}
	message += "\n"

	if pdb.ElectricFreeSpot.IsFree() {
		message += fmt.Sprintf("Free Electric: %s", pdb.ElectricFreeSpot.Number)
	} else {
		message += "Electric is full"
	}

	pdb.Show(message)
}

func (pdb *ParkingDisplayBoard) Show(message string) {
	fmt.Println(message)
}
