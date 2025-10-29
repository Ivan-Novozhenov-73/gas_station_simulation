package services

import (
	"gas_station_simulation/internal/models"
)

func StartSimulation(config *models.Station) {
	createPumps(config)
	go createCars(config)

	config.WaitGroupPumps.Wait()
}
