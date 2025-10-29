package models

import "sync"

type Station struct {
	NumberOfPumps  int
	QueueSize      int
	Queue          chan Car
	NumberOfCars   int
	WaitGroupPumps sync.WaitGroup
	//stopqueue chan int
}

func NewStation(_numberOfPumps, _queueSize, _numberOfCars int) *Station {
	if _numberOfPumps == 0 && _queueSize == 0 && _numberOfCars == 0 {
		return &Station{
			NumberOfPumps:  4,
			QueueSize:      8,
			Queue:          make(chan Car, 8),
			NumberOfCars:   100,
			WaitGroupPumps: sync.WaitGroup{},
		}
	}
	return &Station{
		NumberOfPumps:  _numberOfPumps,
		QueueSize:      _queueSize,
		Queue:          make(chan Car, _queueSize),
		NumberOfCars:   _numberOfCars,
		WaitGroupPumps: sync.WaitGroup{},
	}
}
