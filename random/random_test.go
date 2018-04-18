package random

import "testing"

func TestFloat64(t *testing.T) {
	tests := []struct {
		min, max float64
	}{
		{
			min: -1.0,
			max: 1.0,
		},
		{
			min: 5.0,
			max: 10.0,
		},
		{
			min: 0.0,
			max: 1.0,
		},
	}

	for _, test := range tests {
		for _, output := range Float64Vector(test.min, test.max, 10000) {
			if output > test.max || output < test.min {
				t.Errorf("for min %f and max %f got %f which was outside expected range", test.min, test.max, output)
			}
		}
	}
}
