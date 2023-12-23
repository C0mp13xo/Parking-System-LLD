package common

type VehicleType string

const (
	CAR         VehicleType = "car"
	TRUCK                   = "truck"
	ELECTRICCAR             = "electriccar"
	VAN                     = "van"
	MOTORCYCLE              = "motorcycle"
)

type AccountStatus string

const (
	GOOD        AccountStatus = "good"
	BLOCKED                   = "blocked"
	BANNED                    = "banned"
	COMPROMISED               = "compromised"
	ARCHIVED                  = "archived"
	UNKNOWN                   = "unknown"
)

type ParkingTicketStatus string

const (
	ACTIVE ParkingTicketStatus = "active"
	PAID                       = "paid"
	LOST                       = "lost"
)

type Address struct {
	streetAddress string
	city          string
	state         string
	zipCode       string
	country       string
}

type Person struct {
	name    string
	address Address
	email   string
	phone   string
}
