package Struct

import (
	"time"
)

type Car struct {
	ID                 int
	Fuel               string
	StandQueueEnter    time.Time
	StandQueueTime     time.Duration
	RegisterQueueEnter time.Time
	RegisterQueueTime  time.Duration
	FuelTime           time.Duration
	PayTime            time.Duration
	TotalTime          time.Duration
}
