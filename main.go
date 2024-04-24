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

	// Creating fuel stands
	var stands []*Struct.FuelStand
	standCount := 0
	// Adding gas stands
	for i := 0; i < Struct.NumGas; i++ {
		stands = append(stands, Services.NewFuelStand(standCount, Struct.Gas, Struct.StandBuffer))
		standCount++
	}
	// Adding diesel stands
	for i := 0; i < Struct.NumDiesel; i++ {
		stands = append(stands, Services.NewFuelStand(standCount, Struct.Diesel, Struct.StandBuffer))
		standCount++
	}
	// Adding lpg stands
	for i := 0; i < Struct.NumLPG; i++ {
		stands = append(stands, Services.NewFuelStand(standCount, Struct.LPG, Struct.StandBuffer))
		standCount++
	}
	// Adding electric stands
	for i := 0; i < Struct.NumElectric; i++ {
		stands = append(stands, Services.NewFuelStand(standCount, Struct.Electric, Struct.StandBuffer))
		standCount++
	}
	// Creating registers
	var registers []*Struct.CashRegister
	for i := 0; i < Struct.NumRegisters; i++ {
		registers = append(registers, Services.NewCashRegister(i, Struct.RegisterBuffer))
	}
	everythingEnds.Add(1)
	// Car creation routine
	go Services.CreateCarsRoutine()
	// Stand routines
	Struct.StandCreationWaiter.Add(standCount)
	for _, stand := range stands {
		go Services.FuelStandRoutine(stand)
	}
	Struct.StandCreationWaiter.Wait()
	// CashRegister routines
	for _, register := range registers {
		go Services.RegisterRoutine(register)
	}
	// Car shuffling routine
	go Services.FindStandRoutine(stands)
	// Register shuffling routine
	go Services.FindRegister(registers)

	// Create output.yaml and wait for print (sync.WaitGroup)
	go Services.EvaluationRoutine(&everythingEnds)

	// End synchronizations
	Struct.StandFinishWaiter.Wait()
	close(Struct.BuildingQueue)
	Struct.RegisterWaiter.Wait()
	close(Struct.Exit)

	everythingEnds.Wait()
}

// loadConfigFile loads configuration from yaml into set variables
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
