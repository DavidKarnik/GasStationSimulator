package main

import (
	"gasStation/Services"
	"gasStation/Struct"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"sync"
)

// my main synchronization
var everythingEnds sync.WaitGroup

// main function
func main() {
	loadConfigFile()

	// initialize stands and registers -----------------------------------------
	var stands []*Struct.FuelStand
	var registers []*Struct.CashRegister
	standCount := 0
	// map with fuel types -> gas stands
	fuelTypes := map[Struct.FuelType]int{
		Struct.Gas:      Struct.NumGas,
		Struct.Diesel:   Struct.NumDiesel,
		Struct.LPG:      Struct.NumLPG,
		Struct.Electric: Struct.NumElectric,
	}
	// .. append all gas stands
	for fuelType, numStands := range fuelTypes {
		for i := 0; i < numStands; i++ {
			stands = append(stands, Services.NewFuelStand(standCount, fuelType, Struct.StandBuffer))
			standCount++
		}
	}
	// add cash registers to array
	for i := 0; i < Struct.NumRegisters; i++ {
		registers = append(registers, Services.NewCashRegister(i, Struct.RegisterBuffer))
	}
	everythingEnds.Add(1)

	// Creating routines ------------------------------------------------------
	go Services.CreateCarsRoutine() // car
	// for Stands
	Struct.StandCreationWaiter.Add(standCount)
	for _, stand := range stands {
		go Services.FuelStandRoutine(stand)
	}
	Struct.StandCreationWaiter.Wait()
	// for Registers
	for _, register := range registers {
		go Services.RegisterRoutine(register)
	}

	// go routines for finding best scenarios
	go Services.FindStandRoutine(stands)
	go Services.FindCashRegister(registers)

	// create output.yaml and wait for print (sync.WaitGroup)
	go Services.EvaluationRoutine(&everythingEnds)

	// end synchronizations -> close
	Struct.StandFinishWaiter.Wait()
	close(Struct.BuildingQueue)
	Struct.RegisterWaiter.Wait()
	close(Struct.Exit)

	everythingEnds.Wait()
}

// loadConfigFile - loads configuration from yaml into set variables
func loadConfigFile() {
	file, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Error reading config.yaml file: %v", err)
	}

	var config Struct.Config
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatalf("Error unmarshalling config.yaml file: %v", err)
	}

	fillStructVariablesForServices(config)
}

func fillStructVariablesForServices(config Struct.Config) {
	Struct.ArrivalTimeMin = config.Cars.ArrivalTimeMin
	Struct.ArrivalTimeMax = config.Cars.ArrivalTimeMax
	Struct.NumCars = config.Cars.Count

	Struct.NumGas = config.Stations.Gas.Count
	Struct.GasMinT = config.Stations.Gas.ServeTimeMin
	Struct.GasMaxT = config.Stations.Gas.ServeTimeMax

	Struct.NumDiesel = config.Stations.Diesel.Count
	Struct.DieselMinT = config.Stations.Diesel.ServeTimeMin
	Struct.DieselMaxT = config.Stations.Diesel.ServeTimeMax

	Struct.NumLPG = config.Stations.Lpg.Count
	Struct.LpgMinT = config.Stations.Lpg.ServeTimeMin
	Struct.LpgMaxT = config.Stations.Lpg.ServeTimeMax

	Struct.NumElectric = config.Stations.Electric.Count
	Struct.ElectricMinT = config.Stations.Electric.ServeTimeMin
	Struct.ElectricMaxT = config.Stations.Electric.ServeTimeMax

	Struct.NumRegisters = config.Registers.Count
	Struct.MinPaymentT = config.Registers.HandleTimeMin
	Struct.MaxPaymentT = config.Registers.HandleTimeMax
}
