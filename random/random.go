package random

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func Float64(min, max float64) float64 {
	return (max-min)*rand.Float64() + min
}

func Float64Vector(min, max float64, size int) []float64 {
	op := make([]float64, size)
	for i := 0; i < size; i++ {
		op[i] = Float64(min, max)
	}
	return op
}
