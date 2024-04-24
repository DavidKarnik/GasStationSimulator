package Services

import "time"

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
