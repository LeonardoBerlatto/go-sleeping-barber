package sleep

import (
	"math/rand"
	"time"
)

func Random(min, max time.Duration) time.Duration {
	return time.Duration(rand.Intn(int(max-min)) + int(min))
}
