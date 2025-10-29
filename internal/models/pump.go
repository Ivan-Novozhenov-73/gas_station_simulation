package models

// Модель данных для колонки
type Pump struct {
	ID         int
	Income     float32
	FuelVolume map[Fuel]int
}

func NewPump(new_id int) *Pump {
	return &Pump{
		ID:     new_id,
		Income: 0,
		FuelVolume: map[Fuel]int{
			A_95:   0,
			A_92:   0,
			Diesel: 0,
		},
	}
}

func (pump *Pump) SetPump(_income float32, _fuelVolume int, _fuel Fuel) {
	pump.Income += _income
	pump.FuelVolume[_fuel] += _fuelVolume
}
