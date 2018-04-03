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
		{
			name:     "Nil input p",
			p:        nil,
			q:        []float64{3, 3},
			expected: ErrNilVector,
		},
		{
			name:     "Nil input q",
			p:        []float64{-1},
			q:        nil,
			expected: ErrNilVector,
		},
	}

	for _, test := range tests {
		testFunction("Chebyshev", func() (d float64, err error) {
			return Chebyshev(test.p, test.q)
		}, test.expected, t)
		testFunction("Euclidean", func() (d float64, err error) {
			return Euclidean(test.p, test.q)
		}, test.expected, t)
		testFunction("Manhattan", func() (d float64, err error) {
			return Manhattan(test.p, test.q)
		}, test.expected, t)
		testFunction("SumOfSquares", func() (d float64, err error) {
			return SumOfSquares(test.p, test.q)
		}, test.expected, t)
		testFunction("RootMeanSquare", func() (d float64, err error) {
			return RootMeanSquare(test.p, test.q)
		}, test.expected, t)
	}
}

func testFunction(name string, function func() (d float64, err error), expected error, t *testing.T) {
	_, actual := function()
	if actual == nil {
		t.Errorf("%s: expected %v, but got nil", name, expected)
	}
	if actual != expected {
		t.Errorf("%s: expected error %v, but got %v", name, expected, actual)
	}
}
