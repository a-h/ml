package distance

import (
	"math"
)

// RootMeanSquare calculates the root mean square (RMS) distance between two input vectors.
func RootMeanSquare(p []float64, q []float64) (r float64, err error) {
	r, err = MeanSquare(p, q)
	r = math.Sqrt(r)
	return
}
