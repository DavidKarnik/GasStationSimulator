package Services

import (
	"gasStation/Struct"
	"math/rand"
	"time"
)

func runStation(stationFree chan struct{}, queue chan Struct.Car, station Struct.Station) {
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

func getStationStats(stationType string, stations Struct.Config) (stats Struct.Station) {
	if stationType == "gas" {
		stats = stations.Stations.Gas
	} else {
		// Handle error or other station types if needed
	}
	stats.avgQueueTime = calculateAvgWaitTime(stats.totalCars, 0, stats.totalTime)
	return stats
}
