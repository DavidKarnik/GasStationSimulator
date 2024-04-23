package Struct

import (
	"sync"
	"time"
)

// Vars - numbers just for initialization

var Arrivals = make(chan *Car, 10)
var ArrivalTimeMin = 1
var ArrivalTimeMax = 2
var NumCars = 10

type FuelType string

const (
	Gas      = "gas"
	Diesel   = "diesel"
	LPG      = "LPG"
	Electric = "electric"
)

type Car struct {
	ID                 int
	Fuel               FuelType
	StandQueueEnter    time.Time
	StandQueueTime     time.Duration
	RegisterQueueEnter time.Time
	RegisterQueueTime  time.Duration
	FuelTime           time.Duration
	PayTime            time.Duration
	TotalTime          time.Duration
	CarSync            *sync.WaitGroup
}
