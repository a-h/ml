package rbf

import (
	"reflect"
	"testing"
)

func TestBiasNode(t *testing.T) {
	tests := []struct {
		name     string
		count    int
		input    []float64
		expected []float64
	}{
		{
			name:     "get the same output regardless of input (1)",
			count:    1,
			input:    []float64{6.0},
			expected: []float64{1.0},
		},
		{
			name:     "get the same output regardless of input (2)",
			count:    1,
			input:    []float64{12.0},
			expected: []float64{1.0},
		},
		{
			name:     "changing the count increases the number of output values",
			count:    3,
			input:    []float64{12.0},
			expected: []float64{1.0, 1.0, 1.0},
		},
		{
			name:     "the input is ignored",
			count:    3,
			input:    nil,
			expected: []float64{1.0, 1.0, 1.0},
		},
	}

	for _, test := range tests {
		b := NewBias(test.count)
		actual, err := b.Calculate(test.input)
		if err != nil {
			t.Errorf("%s: unexpected error: %v", test.name, err)
			continue
		}
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("%s: for input %v, expected output %v, but got %v", test.name,
				test.input, test.expected, actual)
		}
		if b.OutputCount() != test.count {
			t.Errorf("%s: expected output count of %d, got %d", test.name, test.count, b.OutputCount())
		}
	}
}
