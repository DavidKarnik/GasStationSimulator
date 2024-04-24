package Services

import (
	"gasStation/Struct"
	"math/rand"
	"time"
)

func runRegister(registerFree chan struct{}, queue chan Struct.Car, register Struct.CashRegister) {
	for {
		<-registerFree
		car := <-queue
		register.mutex.Lock()
		register.totalCars++
		register.mutex.Unlock()
		car.registerTime = time.Duration(rand.Intn(int(register.handleTimeMax-register.handleTimeMin))) + register.handleTimeMin
		time.Sleep(car.registerTime)
		register.totalTime += car.registerTime
		register.mutex.Unlock()
		registerFree <- struct{}{}
	}
}
