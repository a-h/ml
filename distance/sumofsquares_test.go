package distance

import "testing"

func TestSumOfSquares(t *testing.T) {
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
			expected: 5,
		},
		{
			name:     "Two below",
			p:        []float64{2, 3, 4, 5, 6},
			q:        []float64{0, 1, 2, 3, 4},
			expected: 4 + 4 + 4 + 4 + 4,
		},
	}

	for _, test := range tests {
		actual, err := SumOfSquares(test.p, test.q)
		if err != nil {
			t.Fatalf("%s: %v", test.name, err)
		}
		if actual != test.expected {
			t.Errorf("%s: for input %v and %v, expected %v, but got %v",
				test.name, test.p, test.q, test.expected, actual)
		}
	}
}
