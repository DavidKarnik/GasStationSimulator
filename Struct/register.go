package Struct

import (
	"sync"
)

var RegisterWaiter sync.WaitGroup
var RegisterBuffer = 3
var NumRegisters = 1
var MinPaymentT = 1
var MaxPaymentT = 2

var BuildingQueue = make(chan *Car, 10)
var Exit = make(chan *Car)

type CashRegister struct {
	Id    int
	Queue chan *Car
}
