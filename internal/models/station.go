package models

type Station struct {
	numberOfPumps int
	queueSize     int
	queue         chan Car
	//stopqueue chan int
}

func NewStation(_numberOfPumps, _queueSize int) *Station {
	if _numberOfPumps == 0 || _queueSize == 0 {
		return &Station{
			numberOfPumps: 4,
			queueSize:     8,
			queue:         make(chan Car, 8),
		}
	}
	return &Station{
		numberOfPumps: _numberOfPumps,
		queueSize:     _queueSize,
		queue:         make(chan Car, _queueSize),
	}
}
