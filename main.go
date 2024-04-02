package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"math/rand"
	"os"
	"sync"
	"time"
)

// Car represents a car in the simulation
type Car struct {
	id           int
	arriveTime   time.Time
	queueTime    time.Duration
	stationTime  time.Duration
	registerTime time.Duration
}

// Station represents a gas station pump
type Station struct {
	stationType  string
	serveTimeMin time.Duration
	serveTimeMax time.Duration
	mutex        sync.Mutex
	totalCars    int
	totalTime    time.Duration
}

// CashRegister represents a cash register at the station
type CashRegister struct {
	handleTimeMin time.Duration
	handleTimeMax time.Duration
	mutex         sync.Mutex
	totalCars     int
	totalTime     time.Duration
}

// Config stores the configuration parameters for the simulation
type Config struct {
	Cars struct {
		Count          int           `yaml:"count"`
		ArrivalTimeMin time.Duration `yaml:"arrival_time_min"`
		ArrivalTimeMax time.Duration `yaml:"arrival_time_max"`
	}
	Stations struct {
		Gas struct {
			Count        int           `yaml:"count"`
			ServeTimeMin time.Duration `yaml:"serve_time_min"`
			ServeTimeMax time.Duration `yaml:"serve_time_max"`
		}
	}
	Registers struct {
		Count         int           `yaml:"count"`
		HandleTimeMin time.Duration `yaml:"handle_time_min"`
		HandleTimeMax time.Duration `yaml:"handle_time_max"`
	}
}

func main() {
	// Načtení konfigurace z config.yaml
	var config Config
	// ... (načtení konfigurace z YAML)
	path := "./config.yaml"
	cfg, err := loadConfig(path)
	if err != nil {
		fmt.Println("Chyba při načítání konfigurace:", err)
		return
	}
	// Inicializace kanálů
	carChannel := make(chan Car, cfg.Cars.Count)
	queue := make(chan Car, cfg.Cars.Count) // Buffer channel to limit queue size (optional)
	stationFree := make(chan struct{}, cfg.Stations.Gas.Count)
	registerFree := make(chan struct{}, cfg.Registers.Count)

	// Spuštění goroutines
	go func() {
		for i := 0; i < config.Cars.Count; i++ {
			car := Car{id: i, arriveTime: time.Now()}
			time.Sleep(time.Duration(rand.Intn(int(config.Cars.ArrivalTimeMax-config.Cars.ArrivalTimeMin))) + config.Cars.ArrivalTimeMin)
			carChannel <- car
		}
		close(carChannel)
	}()

	go runStation(stationFree, queue, Station{stationType: "gas", serveTimeMin: config.Stations.Gas.ServeTimeMin, serveTimeMax: config.Stations.Gas.ServeTimeMax})

	for i := 0; i < config.Registers.Count; i++ {
		go runRegister(registerFree, queue, CashRegister{handleTimeMin: config.Registers.HandleTimeMin, handleTimeMax: config.Registers.HandleTimeMax})
	}

	// Simulace
	var totalQueueTime time.Duration
	var totalStationTime time.Duration
	var totalRegisterTime time.Duration
	var startTime = time.Now()
	for i := 0; i < config.Cars.Count; i++ {
		car := <-carChannel
		car.queueTime = time.Since(car.arriveTime)
		queue <- car
		stationFree <- struct{}{}
	}

	for i := 0; i < config.Cars.Count; i++ {
		car := <-queue
		totalQueueTime += car.queueTime
		totalStationTime += car.stationTime
		totalRegisterTime += car.registerTime
	}

	// Výpis statistik
	fmt.Println("Statistiky:")
	fmt.Println("Stanice:")
	fmt.Println("  Gas:")
	fmt.Printf("    Celkem aut: %d\n", getStationStats("gas", config.Stations).totalCars)
	fmt.Printf("    Průměrná doba čekání: %s\n", calculateAvgRegisterTime(totalRegisterTime, config.Cars.Count))
	fmt.Println("Celkový čas simulace:", time.Since(startTime))
}

func loadConfig(path string) (config Config, err error) {
	// Otevření souboru s konfigurací
	file, err := os.Open(path)
	if err != nil {
		return config, err
	}
	defer file.Close()

	// Dekódování YAML do struktury Config
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func runStation(stationFree chan struct{}, queue chan Car, station Station) {
	for {
		<-stationFree
		car := <-queue
		station.mutex.Lock()
		station.totalCars++
		station.mutex.Unlock()
		car.stationTime = time.Duration(rand.Intn(int(station.serveTimeMax-station.serveTimeMin))) + station.serveTimeMin
		time.Sleep(car.stationTime)
		station.totalTime += car.stationTime
		station.mutex.Unlock()
		stationFree <- struct{}{}
	}
}

func runRegister(registerFree chan struct{}, queue chan Car, register CashRegister) {
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

func getStationStats(stationType string, stations Config.Station) (stats Station) {
	if stationType == "gas" {
		stats = stations.Gas
	} else {
		// Handle error or other station types if needed
	}
	stats.avgQueueTime = calculateAvgWaitTime(stats.totalCars, 0, stats.totalTime)
	return stats
}

func calculateAvgWaitTime(totalCars int, totalWaitTime time.Duration, totalTime time.Duration) time.Duration {
	if totalCars == 0 {
		return 0
	}
	return totalWaitTime / time.Duration(totalCars)
}

func calculateAvgRegisterTime(totalTime time.Duration, totalCars int) time.Duration {
	if totalCars == 0 {
		return 0
	}
	return totalTime / time.Duration(totalCars)
}
