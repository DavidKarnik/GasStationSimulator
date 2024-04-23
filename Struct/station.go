package Struct

import (
	"sync"
)

// Vars - numbers just for initialization

var NumGas = 1
var NumDiesel = 1
var NumLPG = 1
var NumElectric = 1

var GasMinT = 1
var GasMaxT = 2
var DieselMinT = 1
var DieselMaxT = 2
var LpgMinT = 1
var LpgMaxT = 2
var ElectricMinT = 1
var ElectricMaxT = 2

var StandFinishWaiter sync.WaitGroup
var StandCreationWaiter sync.WaitGroup
var StandBuffer = 2

type FuelStand struct {
	Id    int
	Type  FuelType
	Queue chan *Car
}
