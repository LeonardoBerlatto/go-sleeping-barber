package sleep

import (
	"testing"
	"time"
)

func TestRandom(t *testing.T) {
	minTime := 1 * time.Second
	maxTime := 5 * time.Second

	for i := 0; i < 100; i++ {
		result := Random(minTime, maxTime)
		if result < minTime || result > maxTime {
			t.Errorf("Random(%v, %v) = %v, want between %v and %v", minTime, maxTime, result, minTime, maxTime)
		}
	}
}
