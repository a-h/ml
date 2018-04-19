package rbf

import "testing"

func TestGaussian(t *testing.T) {
	actual := NewGaussian(3.0, 2.0, 6.0)(2.0)
	if actual != 3.0 {
		t.Errorf("expected the peak of the curve to be at 2.0 and the value to be 3.0, but got %v", actual)
	}
}

func TestGaussianVector(t *testing.T) {
	center := []float64{3.0, 5.0}
	deviation := []float64{0.0, 0.0}
	f, err := NewGaussianVector(6.0, center, deviation)
	if err != nil {
		t.Fatal("unexpected error executing the NewGaussianVector function:", err)
	}
	actual, err := f([]float64{3.0, 5.0})
	if err != nil {
		t.Fatal("unexpected error executing VectorFunction:", err)
	}
	if actual != 3.0 {
		t.Errorf("expected the peak of the curve to be at { 3.0, 5.0 } and the value to be 6.0, but got %v", actual)
	}
}
