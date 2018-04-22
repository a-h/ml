package rbf

import "testing"

func TestGaussian(t *testing.T) {
	actual := NewGaussian(3.0, 2.0, 6.0)(2.0)
	if actual != 3.0 {
		t.Errorf("expected the peak of the curve to be at 2.0 and the value to be 3.0, but got %v", actual)
	}
}

func TestGaussianVectorSuccess(t *testing.T) {
	center := []float64{3.0, 5.0}
	f := NewGaussianVector(6.0, center, 1.0)
	actual, err := f([]float64{3.0, 5.0})
	if err != nil {
		t.Fatal("unexpected error executing VectorFunction:", err)
	}
	if actual != 6.0 {
		t.Errorf("expected the peak of the curve to be at { 3.0, 5.0 } and the value to be 6.0, but got %v", actual)
	}
}

func TestGaussianVectorErrors(t *testing.T) {
	// Executing with invalid values.
	f := NewGaussianVector(6.0, []float64{3.0, 5.0}, 1.0)
	_, err := f([]float64{})
	if err == nil {
		t.Fatal("expected error executing the function with invalid parameters, but didn't get one")
	}
	_, err = f(nil)
	if err == nil {
		t.Fatal("expected error executing the function with a nil parameter, but didn't get one")
	}
}
