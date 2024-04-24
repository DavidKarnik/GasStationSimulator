package Services

import (
	"math/rand"
	"time"
)

// commonService - used in registerService and stationService

// randomTime - Generates a random time from min to max
func randomTime(min, max int) time.Duration {
	return time.Duration(rand.Intn(max-min) + min)
}

// doSleeping - Blocking Sleep
func doSleeping(delay time.Duration) {
	time.Sleep(delay * time.Millisecond)
}
