package rbf

import (
	"fmt"
	"math"
)

// NewGaussian creates a 1D Gaussian function with the following parameters.
// a = height of the peak
// b = center position
// c = standard deviation
func NewGaussian(a, b, c float64) Function {
	return func(x float64) float64 {
		numerator := (b - x) * (b - x)
		denominator := 2.0 * (c * c)
		return a * math.Pow(math.E, -(numerator/denominator))
	}
}

// NewGaussianVector creates a Gaussian function with the following parameters.
// a = height of the peak
// b = center positions for each dimension
// c = standard deviations for each dimension
// See https://math.stackexchange.com/questions/112406/gaussian-formula-for-n-dimensions
func NewGaussianVector(a float64, b []float64, c []float64) (f VectorFunction, err error) {
	if len(b) != len(c) {
		err = fmt.Errorf("gaussian: cannot create function with mismatched dimensions (%d and %d)", len(b), len(c))
	}
	f = func(v []float64) float64 {
		var sum float64
		for i, x := range v {
			numerator := (b[i] - x) * (b[i] - x)
			denominator := 2.0 * (c[i] * c[i])
			sum += numerator / denominator
		}
		return a * math.Pow(math.E, -sum)
	}
	return
}
