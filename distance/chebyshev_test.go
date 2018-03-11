package distance

import "testing"

func TestChebyshev(t *testing.T) {
	tests := []struct {
		name     string
		p        []float64
		q        []float64
		expected float64
	}{
		{
			name:     "Pythagoran triangle from origin",
			p:        []float64{0, 0},
			q:        []float64{4, 3},
			expected: 4,
		},
		{
			name:     "Larger X",
			p:        []float64{32, 0},
			q:        []float64{16, 3},
			expected: 16,
		},
		{
			name:     "Larger Y",
			p:        []float64{0, -32},
			q:        []float64{0, 0},
			expected: 32,
		},
	}

	for _, test := range tests {
		actual, err := Chebyshev(test.p, test.q)
		if err != nil {
			t.Fatalf("%s: %v", test.name, err)
		}
		if actual != test.expected {
			t.Errorf("%s: for input %v and %v, expected %v, but got %v",
				test.name, test.p, test.q, test.expected, actual)
		}
	}
}
