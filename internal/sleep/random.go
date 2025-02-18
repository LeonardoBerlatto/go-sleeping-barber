package sleep

import (
	"math/rand"
	"time"
)

func RandomSleep(max time.Duration) {
	time.Sleep(time.Duration(rand.Intn(int(max))))
}

func Random(min, max time.Duration) time.Duration {
	return time.Duration(rand.Intn(int(max-min)) + int(min))
}
