package models

type Car struct {
	ID         int
	Fuel       Fuel
	TankVolume int
}

func NewCar(_id int, _fuel Fuel, _tankVolume int) *Car {
	return &Car{
		ID:         _id,
		Fuel:       _fuel,
		TankVolume: _tankVolume,
	}
}
