package distance

import "testing"

func TestErrors(t *testing.T) {
	tests := []struct {
		name     string
		p        []float64
		q        []float64
		expected error
	}{
		{
			name:     "Zero length vectors",
			p:        []float64{},
			q:        []float64{},
			expected: ErrZeroLengthVector,
		},
		{
			name:     "Mismatched lengths",
			p:        []float64{-1},
			q:        []float64{3, 3},
			expected: ErrMismatchedVectorLengths,
		},
	}

	for _, test := range tests {
		testDistance("Chebyshev", func() (d float64, err error) {
			return Chebyshev(test.p, test.q)
		}, test.expected, t)
		testDistance("Euclidean", func() (d float64, err error) {
			return Euclidean(test.p, test.q)
		}, test.expected, t)
		testDistance("Manhattan", func() (d float64, err error) {
			return Manhattan(test.p, test.q)
		}, test.expected, t)
	}
}

func testDistance(name string, function func() (d float64, err error), expected error, t *testing.T) {
	_, actual := function()
	if actual == nil {
		t.Errorf("%s: expected %v, but got nil", name, expected)
	}
	if actual != expected {
		t.Errorf("%s: expected error %v, but got %v", name, expected, actual)
	}
}
