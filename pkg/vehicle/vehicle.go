package Vehicle

type VehicleType int

const (
	CarType VehicleType = iota
	VanType
	TruckType
	MotorcycleType
	ElectricType
)

type Vehicle struct {
	LicenseNumber string
	Type          VehicleType
	Ticket        *ParkingTicket
}

func NewVehicle(vt VehicleType) *Vehicle {
	return &Vehicle{
		Type: vt,
	}
}

func (v *Vehicle) AssignTicket(ticket *ParkingTicket) {
	v.Ticket = ticket
}

type Car struct {
	*Vehicle
}

func NewCar() *Car {
	return &Car{
		Vehicle: NewVehicle(CarType),
	}
}

type Van struct {
	*Vehicle
}

func NewVan() *Van {
	return &Van{
		Vehicle: NewVehicle(VanType),
	}
}

type Truck struct {
	*Vehicle
}

func NewTruck() *Truck {
	return &Truck{
		Vehicle: NewVehicle(TruckType),
	}
}

type Motorcycle struct {
	*Vehicle
}

func NewMotorcycle() *Motorcycle {
	return &Motorcycle{
		Vehicle: NewVehicle(MotorcycleType), // Assuming MotorcycleType is defined
	}
}

type ElectricVehicle struct {
	*Vehicle
}

func NewElectricVehicle() *ElectricVehicle {
	return &ElectricVehicle{
		Vehicle: NewVehicle(ElectricType), // Assuming ElectricType is defined
	}
}

type ParkingTicket struct {
	// Define ParkingTicket attributes here
}
