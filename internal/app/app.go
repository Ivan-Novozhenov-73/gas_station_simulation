package app

import (
	"fmt"
	"gas_station_simulation/internal/config"
	"path/filepath"
)

func RunApp() {
	path := filepath.Join("..", "..", "configs", "config.json")

	config, err := config.LoadConfig(path)
	if err == nil {
		fmt.Println(config)
	}
}
