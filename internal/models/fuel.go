package models

// Модель данных для топлива
type Fuel string

// Виды топлива, поддерживаемые в системе
const (
	A_95   Fuel = "А-95"
	A_92   Fuel = "А-92"
	Diesel Fuel = "Дизель"
)

// Цены на каждый вид топлива
var FuelCost = map[Fuel]float32{
	A_95:   65.41,
	A_92:   61.02,
	Diesel: 72.61,
}
