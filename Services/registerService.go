package Services

import (
	"fmt"
	"gasStation/Struct"
	"time"
)

// NewCashRegister - Creates a CashRegister
func NewCashRegister(id, bufferSize int) *Struct.CashRegister {
	return &Struct.CashRegister{
		Id:    id,
		Queue: make(chan *Struct.Car, bufferSize),
	}
}

// FindCashRegister - Finds the best cash register for a car
func FindCashRegister(registers []*Struct.CashRegister) {
	// Station building queue
	for car := range Struct.BuildingQueue {
		var bestRegister *Struct.CashRegister
		bestQueueLength := -1
		// Finding best register
		for _, register := range registers {
			queueLength := len(register.Queue)
			if bestQueueLength == -1 || queueLength < bestQueueLength {
				bestRegister = register
				bestQueueLength = queueLength
			}
		}
		car.RegisterQueueEnter = time.Now()
		bestRegister.Queue <- car
	}
	// Closing all registers
	for _, register := range registers {
		close(register.Queue)
	}
}

// RegisterRoutine - Go routine for registers
func RegisterRoutine(cs *Struct.CashRegister) {
	defer Struct.RegisterWaiter.Done()
	Struct.RegisterWaiter.Add(1)
	fmt.Printf("Cash register (%d) -> open\n", cs.Id+1)
	// Station shop queue
	for car := range cs.Queue {
		car.RegisterQueueTime = time.Duration(time.Since(car.RegisterQueueEnter).Milliseconds())
		doPaymentSleeping(car)
		// Signaling finished payment to stand
		car.CarSync.Done()
		// Sending car to exit queue
		Struct.Exit <- car
	}
	fmt.Printf("Cash register (%d) -> closed\n", cs.Id+1)
}

// doPaymentSleeping does payment
func doPaymentSleeping(car *Struct.Car) {
	car.PayTime = randomTime(Struct.MinPaymentT, Struct.MaxPaymentT)
	doSleeping(car.PayTime)
}
