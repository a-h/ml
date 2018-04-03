package distance

import (
	"math"
)

// RootMeanSquare calculates the root mean square (RMS) distance between two input vectors.
func RootMeanSquare(p []float64, q []float64) (r float64, err error) {
	if err = validateInputs(p, q); err != nil {
		return
	}
	for i, pp := range p {
		qq := q[i]
		r += (qq - pp) * (qq - pp)
	}
	n := float64(len(p))
	r = math.Sqrt((1.0 / n) * r)
	return
}
