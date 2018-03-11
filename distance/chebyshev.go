package distance

import (
	"math"
)

// Chebyshev calculates the chessboard difference between two vectors.
func Chebyshev(p []float64, q []float64) (d float64, err error) {
	if err = validateInputs(p, q); err != nil {
		return
	}
	for i, pi := range p {
		if n := math.Abs(pi - q[i]); n > d {
			d = n
		}
	}
	return
}
