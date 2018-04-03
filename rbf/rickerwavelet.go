package rbf

import (
	"math"
)

// RickerWavelet radial basis function (RBF).
func RickerWavelet(r float64) float64 {
	return (1.0 - (r * r)) * math.Pow(math.E, (-(r*r)/2))
}
