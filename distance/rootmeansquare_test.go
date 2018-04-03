package distance

import (
	"math"
	"testing"
)

func TestRootMeanSquare(t *testing.T) {
	tests := []struct {
		name     string
		p        []float64
		q        []float64
		expected float64
	}{
		{
			name:     "Equal",
			p:        []float64{1, 2, 3, 4, 5},
			q:        []float64{1, 2, 3, 4, 5},
			expected: 0,
		},
		{
			name:     "One above",
			p:        []float64{1, 2, 3, 4, 5},
			q:        []float64{2, 3, 4, 5, 6},
			expected: 1, // Square root of average error is 1.
		},
		{
			name:     "Two below",
			p:        []float64{2, 3, 4, 5, 6},
			q:        []float64{0, 1, 2, 3, 4},
			expected: 2, // Square root of average error is 1.
		},
		{
			name:     "Same number twice",
			p:        []float64{0, -1},
			q:        []float64{9, 8},
			expected: 9, // Square root of average squared error is 9.
		},
		{
			name:     "More complex",
			p:        []float64{1, 2, 3},
			q:        []float64{2, 12, 14},
			expected: math.Sqrt((1 + 100 + 121) / 3), // Square root of average squared error is 3.
		},
	}

	for _, test := range tests {
		actual, err := RootMeanSquare(test.p, test.q)
		if err != nil {
			t.Fatalf("%s: %v", test.name, err)
		}
		if actual != test.expected {
			t.Errorf("%s: for input %v and %v, expected %v, but got %v",
				test.name, test.p, test.q, test.expected, actual)
		}
	}
}
