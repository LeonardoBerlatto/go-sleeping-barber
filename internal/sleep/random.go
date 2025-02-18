package sleep

import (
	"math/rand"
	"time"
)

func RandomSleep(max time.Duration) {
	time.Sleep(time.Duration(rand.Intn(int(max))))
}
