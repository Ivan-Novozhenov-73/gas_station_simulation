package services

import (
	"fmt"
	"gas_station_simulation/internal/models"
	"time"
)

func createPumps(config *models.Station) {
	for i := 1; i <= config.NumberOfPumps; i++ {
		pump := models.NewPump(i)
		config.WaitGroupPumps.Add(1)
		go func() {
			defer config.WaitGroupPumps.Done()
			pumpWork(pump, config.Queue)
		}()
	}
}

func pumpWork(pump *models.Pump, queue <-chan models.Car) {
	fmt.Printf("Колонока ID:%d начала работу\n", pump.ID)
	for car := range queue {
		fmt.Printf("Колонка ID:%d заправляет машину-%d\n", pump.ID, car.ID)
		income := float32(car.TankVolume) * models.FuelCost[car.Fuel]
		pump.SetPump(income, car.TankVolume, car.Fuel)
		time.Sleep(time.Duration(car.TankVolume*10) * time.Millisecond)
		fmt.Printf("Колонка ID:%d заправила машину-%d\n", pump.ID, car.ID)
	}

	fmt.Printf("Колонка ID:%d завершила работу\n", pump.ID)
}
