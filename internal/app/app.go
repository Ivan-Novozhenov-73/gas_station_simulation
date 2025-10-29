package app

import (
	"fmt"
	"gas_station_simulation/internal/config"
	"gas_station_simulation/internal/services"
	"path/filepath"
)

func RunApp() {
	path := filepath.Join("..", "..", "configs", "config.json")

	config, err := config.LoadConfig(path)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}

	services.StartSimulation(config)
}
