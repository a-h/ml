package clustering

import (
	"testing"
)

func TestAssign(t *testing.T) {
	tests := []struct {
		name       string
		input      []Vector
		assignment []int
		expected   []Cluster
	}{
		{
			name: "All to one",
			input: []Vector{
				Vector{1, 2, 3},
				Vector{4, 5, 6},
			},
			assignment: []int{1, 1},
			expected: []Cluster{
				[]Vector{},
				[]Vector{
					Vector{1, 2, 3},
					Vector{4, 5, 6},
				},
			},
		},
		{
			name: "All to zero",
			input: []Vector{
				Vector{1, 2, 3},
				Vector{4, 5, 6},
			},
			assignment: []int{0, 0},
			expected: []Cluster{
				[]Vector{
					Vector{1, 2, 3},
					Vector{4, 5, 6},
				},
			},
		},
	}

	for _, test := range tests {
		actual, err := Assign(test.input, test.assignment)
		if err != nil {
			t.Fatalf("%s: unexpected error: %v", test.name, err)
		}
		if !Clusters(actual).Eq(Clusters(test.expected)) {
			t.Errorf("%s: Expected %v, got %v", test.name, test.expected, actual)
		}
	}
}

func TestClusterEqual(t *testing.T) {
	tests := []struct {
		a        Cluster
		b        Cluster
		expected bool
	}{
		{
			a: []Vector{
				Vector{1, 2, 3},
				Vector{4, 5, 6},
			},
			b: []Vector{
				Vector{1, 2, 3},
				Vector{4, 5, 6},
			},
			expected: true,
		},
		{
			a: []Vector{
				Vector{1, 2, 3},
				Vector{4, 5, 6},
			},
			b: []Vector{
				Vector{1, 2, 3},
			},
			expected: false,
		},
		{
			a: []Vector{
				Vector{1, 2, 3},
				Vector{4, 6},
			},
			b: []Vector{
				Vector{1, 2, 3},
			},
			expected: false,
		},
		{
			a: []Vector{
				Vector{1, 2, 3},
			},
			b: []Vector{
				Vector{3, 2, 1},
			},
			expected: false,
		},
	}

	for _, test := range tests {
		actual := test.a.Eq(test.b)
		if test.expected != actual {
			t.Fatalf("comparing %v and %v, expected %v, got %v", test.a, test.b, test.expected, actual)
		}
	}
}
