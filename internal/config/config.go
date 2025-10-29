package config

import (
	"encoding/json"
	"errors"
	"gas_station_simulation/internal/models"
	"os"
)

// Модель для десериализации json файла
type configParams struct {
	NumberOfPumps int `json:"number_of_pumps"`
	QueueSize     int `json:"queue_size"`
	NumberOfCars  int `json:"number_of_cars"`
}

// Чтение данных из файла конфига
func LoadConfig(path string) (*models.Station, error) {
	configFile, err := os.Open(path)
	if err != nil {
		defaultStation := models.NewStation(0, 0, 0)
		return defaultStation, errors.New("ошибка при открытии файла конфигурации, использованы стандартные параметры заправки")
	}
	defer configFile.Close()

	params := new(configParams)
	if err = json.NewDecoder(configFile).Decode(&params); err != nil {
		defaultStation := models.NewStation(0, 0, 0)
		return defaultStation, errors.New("ошибка при чтении данных из файла конфигурации, использованы стандартные параметры заправки")
	}

	if params.NumberOfPumps < 4 || params.QueueSize < 8 || params.NumberOfCars < 100 {
		defaultStation := models.NewStation(0, 0, 0)
		return defaultStation, errors.New("слишком маленькие значения у параметров конфигурации:" +
			"\nколичество колонок – минимум 4" +
			"\nмаксимальный размер очереди – минимум 8" +
			"\nколичество машин – минимум 100")
	}

	config := models.NewStation(params.NumberOfPumps, params.QueueSize, params.NumberOfCars)

	return config, nil
}
