package distance

import (
	"math"
)

// Euclidean distance between two vectors.
func Euclidean(p []float64, q []float64) (d float64, err error) {
	if err = validateInputs(p, q); err != nil {
		return
	}
	for i, pi := range p {
		d += (pi - q[i]) * (pi - q[i])
	}
	return math.Sqrt(d), nil
}
