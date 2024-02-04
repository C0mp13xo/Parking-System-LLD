package parkingticket

import (
	"crypto/rand"
	"fmt"
	Vehicle "parking-system-lld/pkg/vehicle"
	"time"
)

type ParkingTicketStatus int

const (
	ACTIVE ParkingTicketStatus = iota
	PAID
	LOST
)

type ParkingTicket struct {
	TicketNumber string
	IssuedAt     time.Time
	PayedAt      time.Time
	Status       ParkingTicketStatus
	AmountPaid   float64
	Vehicle      Vehicle.Vehicle
}

func (status ParkingTicketStatus) String() string {
	switch status {
	case ACTIVE:
		return "Active"
	case PAID:
		return "Paid"
	case LOST:
		return "Lost"
	default:
		return "Unknown"
	}
}

func (t *ParkingTicket) CalculateFee() float64 {

	return 0.0 // Example implementation
}

func (t *ParkingTicket) MarkPaid() {
	t.Status = PAID
}

func (t *ParkingTicket) PrintTicketInformation() {
	fmt.Printf("Ticket ID: %s\nVehicle: %s\nEntry Time: %s\nExit Time: %s\nAmount Paid: %.2f\n",
		t.TicketNumber, t.Vehicle.Type.String(), t.IssuedAt.String(), t.PayedAt.String(), t.AmountPaid)
}

func GenerateTicket(*Vehicle.Vehicle) *ParkingTicket {
	return &ParkingTicket{
		TicketNumber: rand.Prime(),
		
	}
}