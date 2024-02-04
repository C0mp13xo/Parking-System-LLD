package payment

import "parking-system-lld/pkg/parkingticket"

type PaymentProcessor interface {
	ProcessPayment(ticket *parkingticket.ParkingTicket) bool
}

type CashPayment struct {
}

func (c *CashPayment) ProcessPayment(ticket parkingticket.ParkingTicket) bool {

	return true
}

type CreditCardPayment struct {
}

func (c *CreditCardPayment) ProcessPayment(ticket parkingticket.ParkingTicket) bool {

	return true
}
