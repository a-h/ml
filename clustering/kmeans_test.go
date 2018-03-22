package clustering

import (
	"math/rand"
	"testing"

	"github.com/a-h/ml/distance"
)

func TestKMeans(t *testing.T) {
	tests := []struct {
		name  string
		input []Vector
		n     int
		// Each array holds the expected contents of a cluster.
		// For example, expected {0, 1} means that input values
		// input[0] and input[1] should be present together in a
		// cluster.
		expected [][]int
	}{
		{
			name: "Single output",
			input: []Vector{
				{0, 0, 0},
				{1, 1, 1},
			},
			n: 1,
			expected: [][]int{
				{0, 1},
			},
		},
		{
			name: "Two inputs, two outputs",
			input: []Vector{
				{0, 0, 0},
				{1, 1, 1},
			},
			n: 2,
			expected: [][]int{
				{0},
				{1},
			},
		},
		{
			name: "Left vs Right",
			input: []Vector{
				{-12, 0},
				{-16, 0},
				{12, 0},
				{16, 0},
			},
			n: 2,
			expected: [][]int{
				{0, 1},
				{2, 3},
			},
		},
		{
			name: "Up vs Down",
			input: []Vector{
				{3, -100},
				{40, -200},
				{16, 100},
				{10, 200},
			},
			n: 2,
			expected: [][]int{
				{0, 1},
				{2, 3},
			},
		},
	}

	for _, test := range tests {
		assignment, err := KMeans(test.input, test.n, distance.Euclidean)
		if err != nil {
			t.Fatalf("%s: unexpected error: %v", test.name, err)
		}
		if len(assignment) != len(test.input) {
			t.Errorf("%s: expected assignment length %d to equal the amount of data %d", test.name, len(assignment), len(test.input))
		}

		// Clusters are assigned randomly and may be in any order.
		allActualClusters := getClusters(test.input, assignment)

		for _, expectedCluster := range test.expected {
			// Find the first cluster that contains the expected item. The rest of the items should also be
			// there.
			firstExpected := test.input[expectedCluster[0]]

			var allOthersShouldBeFoundIn int
			for i, a := range allActualClusters {
			out:
				for _, vv := range a {
					if firstExpected.Eq(vv) {
						allOthersShouldBeFoundIn = i
						break out
					}
				}
			}

			// Get each expected vector in the cluster, and check that it's present in the
			// same cluster as the first match.
			actualCluster := allActualClusters[allOthersShouldBeFoundIn]
			for _, expectedToContainIndex := range expectedCluster {
				expected := test.input[expectedToContainIndex]
				var found bool
				for _, v := range actualCluster {
					if expected.Eq(v) {
						found = true
					}
				}
				if !found {
					t.Errorf("%s: expected cluster to contain indices %v, but it contained %v", test.name,
						expectedCluster, actualCluster)
				}
			}
		}
	}
}

func getClusters(data []Vector, assignment []int) map[int][]Vector {
	op := map[int][]Vector{}
	for i, a := range assignment {
		op[a] = append(op[a], data[i])
	}
	return op
}

func TestFindNearest(t *testing.T) {
	tests := []struct {
		name      string
		input     Vector
		centroids []Vector
		expected  Vector
	}{
		{
			name:  "A",
			input: Vector{1, 1, 1},
			centroids: []Vector{
				Vector{1, 1, 1},
			},
			expected: Vector{1, 1, 1},
		},
		{
			name:  "B",
			input: Vector{1, 1, 1},
			centroids: []Vector{
				Vector{1, 1, 1},
				Vector{0, 0, 0},
			},
			expected: Vector{1, 1, 1},
		},
		{
			name:  "closest to 1, 1, 1",
			input: Vector{-1, -12, -100},
			centroids: []Vector{
				Vector{3, 7, 9},
				Vector{1, 1, 1},
			},
			expected: Vector{1, 1, 1},
		},
		{
			name:  "mostly close",
			input: Vector{3, 7, 3},
			centroids: []Vector{
				Vector{12, 1, 1},
				Vector{3, 7, 9},
				Vector{12, 6, 3},
			},
			expected: Vector{3, 7, 9},
		},
	}

	for _, test := range tests {
		c, err := findNearest(&test.input, &test.centroids, distance.Euclidean)
		if err != nil {
			t.Fatalf("%s: unexpected error: %v", test.name, err)
		}
		actual := test.centroids[c]
		if !actual.Eq(test.expected) {
			t.Fatalf("%s: expected nearest cluster to have centroid %v, but got %v", test.name, test.expected, actual)
		}
	}
}

func BenchmarkKMeans(b *testing.B) {
	var data = generateData(100, 10000)
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		KMeans(data, 100, distance.Euclidean)
	}
}

func generateData(vectorLength, quantity int) (rv []Vector) {
	rv = make([]Vector, quantity)
	for i := 0; i < quantity; i++ {
		itm := make([]float64, vectorLength)
		for j := 0; j < vectorLength; j++ {
			itm[j] = rand.Float64()
		}
		rv[i] = Vector(itm)
	}
	return
}
