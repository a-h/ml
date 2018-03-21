package clustering

import (
	"math/rand"
	"testing"

	"github.com/a-h/ml/distance"
)

func TestKMeans(t *testing.T) {
	tests := []struct {
		name     string
		input    []Vector
		n        int
		expected []*Cluster
	}{
		{
			name: "Single output",
			input: []Vector{
				{0, 0, 0},
				{1, 1, 1},
			},
			n: 1,
			expected: []*Cluster{
				createCluster(3, Vector{0, 0, 0}, Vector{1, 1, 1}),
			},
		},
		{
			name: "Two inputs, two outputs",
			input: []Vector{
				{0, 0, 0},
				{1, 1, 1},
			},
			n: 2,
			expected: []*Cluster{
				createCluster(3, Vector{0, 0, 0}),
				createCluster(3, Vector{1, 1, 1}),
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
			expected: []*Cluster{
				createCluster(2, Vector{-12, 0}, Vector{-16, 0}),
				createCluster(2, Vector{12, 0}, Vector{16, 0}),
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
			expected: []*Cluster{
				createCluster(2, Vector{3, -100}, Vector{40, -200}),
				createCluster(2, Vector{16, 100}, Vector{10, 200}),
			},
		},
	}

	for _, test := range tests {
		actual, err := KMeans(test.input, test.n, distance.Euclidean)
		if err != nil {
			t.Fatalf("%s: unexpected error: %v", test.name, err)
		}
		if len(actual) != len(test.expected) {
			t.Errorf("%s: expected length of %d, got %d", test.name, len(test.expected), len(actual))
		}
		// Since groups are assigned randomly.
		for _, expectedCluster := range test.expected {
			// Find the first cluster in actual that contains the first item in the expected cluster.
			var firstExpectedVector Vector
			for _, dex := range expectedCluster.Data {
				firstExpectedVector = dex
				break
			}

			actualMatchingCluster := findVectorInClusters(firstExpectedVector, actual)

			if actualMatchingCluster == nil {
				t.Fatalf("%s: could not find cluster which contained vector %v in actual %v", test.name,
					firstExpectedVector, actual)
			}

			// Check that everything in the expected cluster is also present in the actual cluster.
			for _, v := range expectedCluster.Data {
				var contained bool
				for _, vv := range actualMatchingCluster.Data {
					if vv.Eq(v) {
						contained = true
					}
				}
				if !contained {
					t.Errorf("%s: expected cluster to contain %v, but it contained %v", test.name,
						expectedCluster, actualMatchingCluster)
				}
			}
		}
	}
}

func findVectorInClusters(v Vector, clusters []*Cluster) *Cluster {
	for _, c := range clusters {
		for _, cv := range c.Data {
			if cv.Eq(v) {
				return c
			}
		}
	}
	return nil
}

func TestFindNearest(t *testing.T) {
	tests := []struct {
		name                    string
		input                   Vector
		clusters                []*Cluster
		expectedClusterCentroid Vector
	}{
		{
			name:  "A",
			input: Vector{1, 1, 1},
			clusters: []*Cluster{
				createCluster(3, Vector{1, 1, 1}),
			},
			expectedClusterCentroid: Vector{1, 1, 1},
		},
		{
			name:  "B",
			input: Vector{1, 1, 1},
			clusters: []*Cluster{
				createCluster(3, Vector{1, 1, 1}),
				createCluster(3, Vector{0, 0, 0}),
			},
			expectedClusterCentroid: Vector{1, 1, 1},
		},
		{
			name:  "closest to 1, 1, 1",
			input: Vector{-1, -12, -100},
			clusters: []*Cluster{
				createCluster(3, Vector{1, 1, 1}),
				createCluster(3, Vector{3, 7, 9}),
			},
			expectedClusterCentroid: Vector{1, 1, 1},
		},
		{
			name:  "mostly close",
			input: Vector{3, 7, 3},
			clusters: []*Cluster{
				createCluster(3, Vector{12, 1, 1}),
				createCluster(3, Vector{3, 7, 9}),
				createCluster(3, Vector{12, 6, 3}),
			},
			expectedClusterCentroid: Vector{3, 7, 9},
		},
	}

	for _, test := range tests {
		c, err := findNearest(test.input, test.clusters, distance.Euclidean)
		if err != nil {
			t.Fatalf("%s: unexpected error: %v", test.name, err)
		}
		if !c.Centroid.Eq(test.expectedClusterCentroid) {
			t.Fatalf("%s: expected nearest cluster to have centroid %v, but got %v", test.name, test.expectedClusterCentroid, c.Centroid)
		}
	}
}

func createCluster(vectorLength int, vectors ...Vector) *Cluster {
	c := NewCluster(vectorLength)
	if err := c.AddRange(vectors...); err != nil {
		// panic to make writing table-driven tests easier.
		panic(err)
	}
	return c
}

func TestRandomlyAssign(t *testing.T) {
	tests := []struct {
		name     string
		clusters []*Cluster
		vectors  []Vector
	}{
		{
			name: "All in one",
			clusters: []*Cluster{
				NewCluster(3),
			},
			vectors: []Vector{
				{1, 1, 1},
				{2, 2, 2},
			},
		},
		{
			name: "50/50",
			clusters: []*Cluster{
				NewCluster(3),
				NewCluster(3),
			},
			vectors: []Vector{
				{1, 1, 1},
				{2, 2, 2},
			},
		},
	}

	for _, test := range tests {
		randomlyAssign(test.clusters, test.vectors)
		expectedCount := len(test.vectors)
		var actualCount int
		for _, c := range test.clusters {
			actualCount += len(c.Data)
		}
		if actualCount != expectedCount {
			t.Errorf("%s: expected the clusters to be assigned the %d vectors, but %d were present", test.name,
				expectedCount, actualCount)
		}
	}
}

func BenchmarkFib10(b *testing.B) {
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
