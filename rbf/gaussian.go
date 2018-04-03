package rbf

import (
	"math"
)

// Gaussian radial basis function (RBF).
func Gaussian(r float64) float64 {
	return math.Pow(math.E, -(r * r))
}
