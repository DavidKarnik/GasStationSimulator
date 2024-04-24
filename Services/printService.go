package Services

import (
	"fmt"
	"gasStation/Struct"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"sync"
	"time"
)

// EvaluationRoutine - Collects global data for output and prints them
func EvaluationRoutine(end *sync.WaitGroup) {
	// Cars
	var totalCars int
	var totalRegisterTime time.Duration
	var totalRegisterQueue time.Duration
	maxRegisterQueue := 0
	// Gas
	var totalGasTime time.Duration
	var totalGasQueue time.Duration
	maxGasQueue := 0
	gasCount := 0
	// Diesel
	var totalDieselTime time.Duration
	var totalDieselQueue time.Duration
	maxDieselQueue := 0
	dieselCount := 0
	// LPG
	var totalLPGTime time.Duration
	var totalLPGQueue time.Duration
	maxLPGQueue := 0
	lpgCount := 0
	// Electric
	var totalElectricTime time.Duration
	var totalElectricQueue time.Duration
	maxElectricQueue := 0
	electricCount := 0
	// Exit queue aggregates data
	for car := range Struct.Exit {
		totalCars++
		totalRegisterTime += car.PayTime
		totalRegisterQueue += car.RegisterQueueTime
		car.TotalTime = time.Duration(time.Since(car.StandQueueEnter).Milliseconds())
		if int(car.RegisterQueueTime) > maxRegisterQueue {
			maxRegisterQueue = int(car.RegisterQueueTime)
		}
		switch car.Fuel {
		case Struct.Gas:
			totalGasTime += car.FuelTime
			totalGasQueue += car.StandQueueTime
			gasCount++
			if int(car.StandQueueTime) > maxGasQueue {
				maxGasQueue = int(car.StandQueueTime)
			}
		case Struct.Diesel:
			totalDieselTime += car.FuelTime
			totalDieselQueue += car.StandQueueTime
			dieselCount++
			if int(car.StandQueueTime) > maxDieselQueue {
				maxDieselQueue = int(car.StandQueueTime)
			}
		case Struct.LPG:
			totalLPGTime += car.FuelTime
			totalLPGQueue += car.StandQueueTime
			lpgCount++
			if int(car.StandQueueTime) > maxLPGQueue {
				maxLPGQueue = int(car.StandQueueTime)
			}
		case Struct.Electric:
			//totalElectricTime += car.TotalTime
			totalElectricTime += car.FuelTime
			totalElectricQueue += car.StandQueueTime
			electricCount++
			if int(car.StandQueueTime) > maxElectricQueue {
				maxElectricQueue = int(car.StandQueueTime)
			}
		}
	}
	// Average values --------------------------------------------------------------------
	var averageGasQueue int
	if gasCount != 0 {
		averageGasQueue = int(totalGasQueue) / gasCount
	}
	var averageDieselQueue int
	if dieselCount != 0 {
		averageDieselQueue = int(totalDieselQueue) / dieselCount
	}
	var averageLPGQueue int
	if lpgCount != 0 {
		averageLPGQueue = int(totalLPGQueue) / lpgCount
	}
	var averageElectricQueue int
	if electricCount != 0 {
		averageElectricQueue = int(totalElectricQueue) / electricCount
	}
	var averageRegisterQueue int
	if totalCars != 0 {
		averageRegisterQueue = int(totalRegisterQueue) / totalCars
	}
	// Create yaml --------------------------------------------------------------------------
	stats := Struct.FinalStats{
		Gas: Struct.StationStats{
			TotalCars:    gasCount,
			TotalTime:    int(totalGasTime),
			AvgQueueTime: averageGasQueue,
			MaxQueueTime: maxGasQueue,
		},
		Diesel: Struct.StationStats{
			TotalCars:    dieselCount,
			TotalTime:    int(totalDieselTime),
			AvgQueueTime: averageDieselQueue,
			MaxQueueTime: maxDieselQueue,
		},
		LPG: Struct.StationStats{
			TotalCars:    lpgCount,
			TotalTime:    int(totalLPGTime),
			AvgQueueTime: averageLPGQueue,
			MaxQueueTime: maxLPGQueue,
		},
		Electric: Struct.StationStats{
			TotalCars:    electricCount,
			TotalTime:    int(totalElectricTime),
			AvgQueueTime: averageElectricQueue,
			MaxQueueTime: maxElectricQueue,
		},
		Registers: Struct.StationStats{
			TotalCars:    totalCars,
			TotalTime:    int(totalRegisterTime),
			AvgQueueTime: averageRegisterQueue,
			MaxQueueTime: maxRegisterQueue,
		},
	}
	yamlStats, err := yaml.Marshal(&stats)
	if err != nil {
		log.Fatalf("Error marshalling stats to YAML: %v", err)
	}

	err = os.WriteFile("output.yaml", yamlStats, 0777)
	if err != nil {
		log.Fatalf("Error writing YAML to file: %v", err)
	}

	fmt.Printf("Final statistics:\n%s\n", string(yamlStats))

	end.Done()
}
