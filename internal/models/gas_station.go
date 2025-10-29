package models

type Pump struct {
	id          int
	income      float32
	fuel_volume map[Fuel]int
}

func NewPump(new_id int) *Pump {
	return &Pump{
		id:     new_id,
		income: 0,
		fuel_volume: map[Fuel]int{
			A_95:   0,
			A_92:   0,
			Diesel: 0,
		},
	}
}
