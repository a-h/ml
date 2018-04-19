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
	deviation := []float64{1.0, 1.0}
	f, err := NewGaussianVector(6.0, center, deviation)
	if err != nil {
		t.Fatal("unexpected error executing the NewGaussianVector function:", err)
	}
	actual, err := f([]float64{3.0, 5.0})
	if err != nil {
		t.Fatal("unexpected error executing VectorFunction:", err)
	}
	if actual != 6.0 {
		t.Errorf("expected the peak of the curve to be at { 3.0, 5.0 } and the value to be 6.0, but got %v", actual)
	}
}

func TestGaussianVectorErrors(t *testing.T) {
	// Creating the function.
	_, err := NewGaussianVector(6.0, []float64{3.0, 5.0}, nil)
	if err == nil {
		t.Fatal("expected error executing the NewGaussianVector function, but didn't get one")
	}
	if err.Error() != "gaussian: cannot create function with mismatched dimensions (2 and 0)" {
		t.Errorf("unexpected error message: %v", err)
	}

	// Executing with invalid values.
	f, err := NewGaussianVector(6.0, []float64{3.0, 5.0}, []float64{1.0, 2.0})
	if err != nil {
		t.Fatal("unexpected error executing the NewGaussianVector function:", err)
	}
	_, err = f([]float64{})
	if err == nil {
		t.Fatal("expected error executing the function with invalid parameters, but didn't get one")
	}
	_, err = f(nil)
	if err == nil {
		t.Fatal("expected error executing the function with a nil parameter, but didn't get one")
	}
}
