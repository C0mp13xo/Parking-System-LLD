package account

type Account struct {
	UserName string
	Password string
	Status   AccountStatus
	Person   Person
}

type AccountStatus int

const (
// Define status constants here
)

type Person struct {
	// Define Person attributes here
}

func (a *Account) ResetPassword() bool {
	// Implement resetPassword logic here
	return true
}

type Admin struct {
	Account
}

func (admin *Admin) AddParkingFloor(floor ParkingFloor) bool {
	// Implement logic to add parking floor
	return true
}

func (admin *Admin) AddParkingSpot(floorName string, spot ParkingSpot) bool {
	// Implement logic to add parking spot
	return true
}

func (admin *Admin) AddParkingDisplayBoard(floorName string, displayBoard ParkingDisplayBoard) bool {
	// Implement logic to add parking display board
	return true
}

func (admin *Admin) AddCustomerInfoPanel(floorName string, infoPanel CustomerInfoPanel) bool {
	// Implement logic to add customer info panel
	return true
}

func (admin *Admin) AddEntrancePanel(entrancePanel EntrancePanel) bool {
	// Implement logic to add entrance panel
	return true
}

func (admin *Admin) AddExitPanel(exitPanel ExitPanel) bool {
	// Implement logic to add exit panel
	return true
}

type ParkingAttendant struct {
	Account
}

func (attendant *ParkingAttendant) ProcessTicket(ticketNumber string) bool {
	// Implement logic to process ticket
	return true
}

type ParkingFloor struct {
	// Define ParkingFloor attributes here
}

type ParkingSpot struct {
	// Define ParkingSpot attributes here
}

type ParkingDisplayBoard struct {
	// Define ParkingDisplayBoard attributes here
}

type CustomerInfoPanel struct {
	// Define CustomerInfoPanel attributes here
}

type EntrancePanel struct {
	// Define EntrancePanel attributes here
}

type ExitPanel struct {
	// Define ExitPanel attributes here
}
