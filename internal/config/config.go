package config

import (
	"encoding/json"
	"errors"
	"gas_station_simulation/internal/models"
	"os"
)

type configParams struct {
	NumberOfPumps int `json:"number_of_pumps"`
	QueueSize     int `json:"queue_size"`
}

func LoadConfig(path string) (*models.Station, error) {
	configFile, err := os.Open(path)
	if err != nil {
		defaultStation := models.NewStation(0, 0)
		return defaultStation, errors.New("ошибка при открытии файла конфигурации, использованы стандартные параметры заправки")
	}
	defer configFile.Close()

	params := new(configParams)
	if err = json.NewDecoder(configFile).Decode(&params); err != nil {
		defaultStation := models.NewStation(0, 0)
		return defaultStation, errors.New("ошибка при чтении данных из файла конфигурации, использованы стандартные параметры заправки")
	}

	if params.NumberOfPumps < 4 || params.QueueSize < 8 {
		defaultStation := models.NewStation(0, 0)
		return defaultStation, errors.New("минимальное количество колонок – 4, минимальный размер очереди – 8")
	}

	config := models.NewStation(params.NumberOfPumps, params.QueueSize)

	return config, nil
}
