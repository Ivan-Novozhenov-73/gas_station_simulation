package services

import (
	"fmt"
	"gas_station_simulation/internal/models"
)

func StartSimulation(config *models.Station) {
	createPumps(config)
	go createCars(config)

	config.WaitGroupPumps.Wait()

	fmt.Println("\nИтоговый отчет:")
	fmt.Printf("---------------------------\n")
	fmt.Printf("Обслужено автомобилей: %d\n", config.CarsServed)
	fmt.Printf("Отказов по очереди: %d\n", config.CarsRefused)
	fmt.Printf("Выручка: %f₽\n", config.Income)
	fmt.Printf("Продано топлива:\n")
	fmt.Printf("\tА-92: %d\n", config.FuelVolume[models.A_92])
	fmt.Printf("\tА-95: %d\n", config.FuelVolume[models.A_92])
	fmt.Printf("\tДизель: %d\n", config.FuelVolume[models.A_92])
	fmt.Printf("---------------------------")
}
