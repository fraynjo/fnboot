package internal

import (
	"math/rand"
	"time"
)

func InRange[T FnNumber, T1 FnNumber, T2 FnNumber](value T, start T1, end T2) bool {
	return float64(value) >= float64(start) && float64(value) <= float64(end)
}

func Random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
