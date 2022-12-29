package number

import (
	"github.com/fraynjo/fnboot"
	"math/rand"
	"time"
)

func InRange[T fnboot.FnNumber, T1 fnboot.FnNumber, T2 fnboot.FnNumber](value T, start T1, end T2) bool {
	return float64(value) >= float64(start) && float64(value) <= float64(end)
}

func Random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
