package parkingticket

import (
	"fmt"
	Vehicle "parking-system-lld/pkg/vehicle"
	"time"

	"github.com/google/uuid"
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
	vehicle      *Vehicle.Vehicle
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
		t.TicketNumber, t.vehicle.Type.String(), t.IssuedAt.String(), t.PayedAt.String(), t.AmountPaid)
}

func GenerateTicket(v *Vehicle.Vehicle) *ParkingTicket {
	id := uuid.New()
	pt := &ParkingTicket{
		TicketNumber: fmt.Sprintf("%x", id[:]),
		IssuedAt:     time.Now(),
		vehicle:      v,
	}
	return pt
}
