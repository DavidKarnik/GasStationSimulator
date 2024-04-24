package Services

import (
	"gasStation/Struct"
	"math/rand"
	"sync"
	"time"
)

// CreateCarsRoutine - Cars arrival routine
func CreateCarsRoutine() {
	for i := 0; i < Struct.NumCars; i++ {
		// Adds a new car to station queue
		Struct.Arrivals <- &Struct.Car{ID: i, Fuel: generateRandomFuelType(), CarSync: &sync.WaitGroup{}, StandQueueEnter: time.Now()}
		// Staggers car creation
		stagger := time.Duration(rand.Intn(Struct.ArrivalTimeMax-Struct.ArrivalTimeMin) + Struct.ArrivalTimeMin)
		time.Sleep(stagger * time.Millisecond)
	}
	close(Struct.Arrivals)
}

func generateRandomFuelType() Struct.FuelType {
	fuelTypes := []Struct.FuelType{Struct.Gas, Struct.Diesel, Struct.Electric, Struct.LPG}
	randomIndex := rand.Intn(len(fuelTypes))
	return fuelTypes[randomIndex]
}
