package models

type Gas_Station struct {
	id          int
	income      float32
	fuel_volume map[Fuel]int
}

func newCasStation(new_id int) *Gas_Station {
	return &Gas_Station{
		id:     new_id,
		income: 0,
		fuel_volume: map[Fuel]int{
			A_95:   0,
			A_92:   0,
			Diesel: 0,
		},
	}
}
