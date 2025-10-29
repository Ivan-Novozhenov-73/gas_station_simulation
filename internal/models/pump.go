package models

type Pump struct {
	ID         int
	income     float32
	fuelVolume map[Fuel]int
}

func NewPump(new_id int) *Pump {
	return &Pump{
		ID:     new_id,
		income: 0,
		fuelVolume: map[Fuel]int{
			A_95:   0,
			A_92:   0,
			Diesel: 0,
		},
	}
}

func (pump *Pump) SetPump(_income float32, _fuelVolume int, _fuel Fuel) {
	pump.income = _income
	pump.fuelVolume[_fuel] = _fuelVolume
}
