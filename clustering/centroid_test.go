package clustering

import (
	"errors"
	"testing"
)

func TestCentroid(t *testing.T) {
	tests := []struct {
		name          string
		vectorLength  int
		data          []Vector
		expected      Vector
		expectedError error
	}{
		{
			name:          "Empty",
			vectorLength:  3,
			data:          []Vector{},
			expected:      Vector{0, 0, 0},
			expectedError: errors.New("centroid: no data provided"),
		},
		{
			name:         "All ones",
			vectorLength: 3,
			data: []Vector{
				Vector{1, 1, 1},
				Vector{1, 1, 1},
				Vector{1, 1, 1},
				Vector{1, 1, 1},
			},
			expected: Vector{1, 1, 1},
		},
		{
			name:         "Average of 1 and 2 is 1.5",
			vectorLength: 3,
			data: []Vector{
				Vector{1, 2, 1},
				Vector{2, 1, 2},
			},
			expected: Vector{1.5, 1.5, 1.5},
		},
	}

	for _, test := range tests {
		actual, err := Centroid(test.data)
		if test.expectedError != nil {
			if err.Error() != test.expectedError.Error() {
				t.Fatalf("%s: expected error '%v', but got '%v'", test.name, test.expectedError, err)
			}
			continue
		}

		if err != nil {
			t.Fatalf("%s: got error %v", test.name, err)
		}
		if !actual.Eq(test.expected) {
			t.Errorf("%s: expected %v, got %v", test.name, test.expected, actual)
		}
	}
}
