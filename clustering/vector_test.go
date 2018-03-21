package clustering

import "testing"

func TestVectorEqual(t *testing.T) {
	tests := []struct {
		name     string
		a        Vector
		b        Vector
		expected bool
	}{
		{
			name:     "All ones",
			a:        Vector{1, 1, 1},
			b:        Vector{1, 1, 1},
			expected: true,
		},
		{
			name:     "Completely different",
			a:        Vector{1, 1, 1},
			b:        Vector{2, 3, 4},
			expected: false,
		},
		{
			name:     "Different lengths",
			a:        Vector{1},
			b:        Vector{2, 3, 4},
			expected: false,
		},
	}

	for _, test := range tests {
		actual := test.a.Eq(test.b)
		if actual != test.expected {
			t.Errorf("%s: expected %v, got %v", test.name, test.expected, actual)
		}
	}
}
