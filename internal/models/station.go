package models

import "sync"

type Station struct {
	NumberOfPumps  int            // количество колонок
	QueueSize      int            // размер очереди на заправку
	Queue          chan Car       // очередь на заправку
	NumberOfCars   int            // количество машин, которые буду заправляться
	WaitGroupPumps sync.WaitGroup // waitgroup для синхронизации работы колонок
	MutexPumps     sync.Mutex     // Мьютекс для записи отчетности колонки
	CarsServed     int32          // количество обслуженных машин
	CarsRefused    int            // количество необслуженных машин
	Income         float32        // общий доход АЗС
	FuelVolume     map[Fuel]int   // количество заправленного топлива
}

// Создает новую заправочную станцию с заданными параметрами
func NewStation(_numberOfPumps, _queueSize, _numberOfCars int) *Station {
	if _numberOfPumps == 0 && _queueSize == 0 && _numberOfCars == 0 {
		return &Station{
			NumberOfPumps:  4,
			QueueSize:      8,
			Queue:          make(chan Car, 8),
			NumberOfCars:   100,
			WaitGroupPumps: sync.WaitGroup{},
			MutexPumps:     sync.Mutex{},
			CarsServed:     0,
			CarsRefused:    0,
			Income:         0,
			FuelVolume: map[Fuel]int{
				A_95:   0,
				A_92:   0,
				Diesel: 0,
			},
		}
	}
	return &Station{
		NumberOfPumps:  _numberOfPumps,
		QueueSize:      _queueSize,
		Queue:          make(chan Car, _queueSize),
		NumberOfCars:   _numberOfCars,
		WaitGroupPumps: sync.WaitGroup{},
		MutexPumps:     sync.Mutex{},
		CarsServed:     0,
		CarsRefused:    0,
		Income:         0,
		FuelVolume: map[Fuel]int{
			A_95:   0,
			A_92:   0,
			Diesel: 0,
		},
	}
}
