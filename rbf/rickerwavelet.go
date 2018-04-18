package rbf

import (
	"math"
)

// NewRickerWavelet radial basis function (RBF).
func NewRickerWavelet(a, b, c float64) Function {
	return func(x float64) float64 {
		numerator := (b - x) * (b - x)
		denominator := 2.0 * (c * c)
		return (a - (x * x)) * math.Pow(math.E, -(numerator/denominator))
	}
}
