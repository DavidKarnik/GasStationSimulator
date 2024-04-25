package Services

import (
	"fmt"
	"gasStation/Struct"
	"time"
)

// NewFuelStand - Creates a FuelStand
func NewFuelStand(id int, fuel Struct.FuelType, bufferSize int) *Struct.FuelStand {
	return &Struct.FuelStand{
		Id:    id,
		Type:  fuel,
		Queue: make(chan *Struct.Car, bufferSize),
	}
}

// FindStandRoutine - Finds the best stand according to fuel type
func FindStandRoutine(stands []*Struct.FuelStand) {
	// Station entrance queue
	for car := range Struct.Arrivals {
		// Initialization
		var bestStand *Struct.FuelStand
		bestQueueLength := -1
		// Finding best stand
		for _, stand := range stands {
			if stand.Type == car.Fuel {
				queueLength := len(stand.Queue)
				if bestQueueLength == -1 || queueLength < bestQueueLength {
					bestStand = stand
					bestQueueLength = queueLength
				}
			}
		}
		bestStand.Queue <- car
	}
	// Closing all stands
	for _, stand := range stands {
		close(stand.Queue)
	}
}

// FuelStandRoutine - Go routine for FuelStand
func FuelStandRoutine(fs *Struct.FuelStand) {
	defer Struct.StandFinishWaiter.Done()
	Struct.StandFinishWaiter.Add(1)
	fmt.Printf("Fuel stand (%d)    -> open\n", fs.Id+1)
	Struct.StandCreationWaiter.Done()
	// Stand queue
	for car := range fs.Queue {
		car.StandQueueTime = time.Duration(time.Since(car.StandQueueEnter).Milliseconds())
		doFuelingSleeping(car)
		car.CarSync.Add(1)
		// Sending car to registers
		Struct.BuildingQueue <- car
		// Wait for payment to complete
		car.CarSync.Wait()
	}
	fmt.Printf("Fuel stand (%d)    -> closed\n", fs.Id+1)
}

// doFuelingSleeping does fueling
func doFuelingSleeping(car *Struct.Car) {
	switch car.Fuel {
	case Struct.Gas:
		car.FuelTime = randomTime(Struct.GasMinT, Struct.GasMaxT)
	case Struct.Diesel:
		car.FuelTime = randomTime(Struct.DieselMinT, Struct.DieselMaxT)
	case Struct.LPG:
		car.FuelTime = randomTime(Struct.LpgMinT, Struct.LpgMaxT)
	case Struct.Electric:
		car.FuelTime = randomTime(Struct.ElectricMinT, Struct.ElectricMaxT)
	}
	doSleeping(car.FuelTime)
}
