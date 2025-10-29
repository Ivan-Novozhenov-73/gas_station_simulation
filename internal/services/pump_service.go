package services

import (
	"fmt"
	"gas_station_simulation/internal/models"
	"sync/atomic"
	"time"
)

// Создает заправочные колонки
func createPumps(config *models.Station) {
	for i := 1; i <= config.NumberOfPumps; i++ {
		pump := models.NewPump(i)
		config.WaitGroupPumps.Add(1)
		go func() {
			defer config.WaitGroupPumps.Done()
			pumpWork(pump, config)
		}()
	}
}

// Иммитирует работу одной заправочной колонки
func pumpWork(pump *models.Pump, config *models.Station) {
	fmt.Printf("Колонка #%d начала работу\n", pump.ID)
	for car := range config.Queue {
		fmt.Printf("Колонка #%d заправляет машину #%d\n", pump.ID, car.ID)
		income := float32(car.TankVolume) * models.FuelCost[car.Fuel]
		pump.SetPump(income, car.TankVolume, car.Fuel)
		time.Sleep(time.Duration(car.TankVolume*10) * time.Millisecond)
		fmt.Printf("Колонка #%d заправила машину #%d на %d л (%f₽)\n",
			pump.ID, car.ID, car.TankVolume, income)
		atomic.AddInt32(&config.CarsServed, 1)
	}
	config.MutexPumps.Lock()
	config.Income += pump.Income
	config.FuelVolume[models.A_92] += pump.FuelVolume[models.A_92]
	config.FuelVolume[models.A_95] += pump.FuelVolume[models.A_95]
	config.FuelVolume[models.Diesel] += pump.FuelVolume[models.Diesel]
	config.MutexPumps.Unlock()

	fmt.Printf("Колонка #%d завершила работу\n", pump.ID)
}
