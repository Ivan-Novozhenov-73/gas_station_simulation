package services

import (
	"fmt"
	"gas_station_simulation/internal/models"
	"math/rand/v2"
	"time"
)

// Создание потока машин
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

		// Проверка на заполненность очередеди
		fmt.Printf("Машина #%d приехала\n", i)
		if len(config.Queue) == cap(config.Queue) {
			fmt.Printf("Машина #%d не обслужена, очередь переполнена\n", i)
			config.CarsRefused++
		} else {
			config.Queue <- *car
		}
		time.Sleep(100 * time.Millisecond)
	}

	close(config.Queue)
}
