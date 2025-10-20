package models

type Car struct {
	id          int
	fuel        Fuel
	tank_volume int
}

func newCar(new_id int, new_fuel Fuel, new_tank_volume int) *Car {
	return &Car{
		id:          new_id,
		fuel:        new_fuel,
		tank_volume: new_tank_volume,
	}
}
