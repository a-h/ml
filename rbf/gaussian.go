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
		denominator := 2.0 * c * c
		return a * math.Exp(-(numerator / denominator))
	}
}

// NewGaussianVector creates a Gaussian function with the following parameters.
// a = height of the peak
// b = center positions for each dimension
// c = standard deviations for each dimension
// See https://math.stackexchange.com/questions/112406/gaussian-formula-for-n-dimensions
func NewGaussianVector(a float64, b []float64, c float64) (f VectorFunction) {
	f = func(v []float64) (float64, error) {
		if len(v) != len(b) {
			err := fmt.Errorf("gaussian: mismached count of comparison vector (%d) to input vector (%d)",
				len(b), len(v))
			return 0.0, err
		}
		var sum float64
		for i, x := range v {
			numerator := (b[i] - x) * (b[i] - x)
			denominator := 2.0 * c * c
			sum += numerator / denominator
		}
		return a * math.Exp(-sum), nil
	}
	return
}
