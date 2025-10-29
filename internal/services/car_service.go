package services

import (
	"gas_station_simulation/internal/models"
	"math/rand/v2"
	"time"
)

func createCars(config *models.Station) {
	for i := 1; i <= config.NumberOfCars; i++ {
		tankVolume := (rand.IntN(4) + 4) * 10
		var fuel models.Fuel
		fuelType := rand.IntN(3)
		switch fuelType {
		case 0:
			fuel = models.A_92
		case 1:
			fuel = models.A_95
		default:
			fuel = models.Diesel
		}
		car := models.NewCar(i, fuel, tankVolume)
		config.Queue <- *car
		time.Sleep(time.Duration(rand.IntN(5)+1) * time.Millisecond)
	}

	close(config.Queue)
}
