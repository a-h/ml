package clustering

import (
	"testing"
)

func TestClusterCentroid(t *testing.T) {
	tests := []struct {
		name         string
		vectorLength int
		data         []Vector
		expected     Vector
	}{
		{
			name:         "Empty",
			vectorLength: 3,
			data:         []Vector{},
			expected:     Vector{0, 0, 0},
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
		cluster := NewCluster(test.vectorLength)
		cluster.AddRange(test.data...)
		actual := cluster.Centroid
		if !actual.Eq(test.expected) {
			t.Fatalf("%s: expected %v, got %v", test.name, test.expected, actual)
		}
	}
}

func TestClusterCentroidChangesOnDataModification(t *testing.T) {
	mutations := []struct {
		operation string
		data      []ClusterMember
		expected  Vector
	}{
		{
			operation: "add",
			data: []ClusterMember{
				ClusterMember{
					ID:     "a",
					Vector: Vector{1, 1, 1},
				},
			},
			expected: Vector{1, 1, 1},
		},
		{
			operation: "addrange",
			data: []ClusterMember{
				ClusterMember{
					ID:     "b",
					Vector: Vector{3, 3, 3},
				},
				ClusterMember{
					ID:     "c",
					Vector: Vector{3, 3, 3},
				},
				ClusterMember{
					ID:     "d",
					Vector: Vector{3, 3, 3},
				},
			},
			expected: Vector{2.5, 2.5, 2.5},
		},
		{
			operation: "delete",
			data: []ClusterMember{
				ClusterMember{
					ID:     "a",
					Vector: Vector{1, 1, 1},
				},
			},
			expected: Vector{3, 3, 3},
		},
	}

	cluster := NewCluster(3)
	for i, test := range mutations {
		switch test.operation {
		case "add":
			err := cluster.AddMember(test.data[0])
			if err != nil {
				t.Errorf("[%d] %s: failed to add to cluster: %s", i, test.operation, err)
			}
		case "addrange":
			err := cluster.AddMembers(test.data...)
			if err != nil {
				t.Errorf("[%d] %s: failed to addrange to cluster: %s", i, test.operation, err)
			}
		case "delete":
			cluster.Remove(test.data[0])
		}
		actual := cluster.Centroid
		if !actual.Eq(test.expected) {
			t.Fatalf("[%d] %s: expected %v, got %v", i, test.operation, test.expected, actual)
		}
	}
}
