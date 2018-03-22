package clustering

import "errors"

// Centroid calculates the average position of the members in the cluster.
func Centroid(data []Vector) (c Vector, err error) {
	if data == nil || len(data) == 0 {
		err = errors.New("centroid: no data provided")
		return
	}

	// Initialize array.
	c = make([]float64, len(data[0]))

	// Sum the data.
	for _, v := range data {
		for j, vj := range v {
			c[j] += vj
		}
	}

	// Divide by the length of data to get the average.
	for i, f := range c {
		c[i] = f / float64(len(data))
	}

	return
}

// Centroids calculates multiple centroids in a single operation, and reduces memory allocations
// by accepting a pointer to an existing centroids vector.
func Centroids(data []Vector, n int, assignments []int, centroids *[]Vector) (err error) {
	if data == nil || len(data) == 0 {
		return errors.New("centroids: no data provided")
	}
	if centroids == nil {
		return errors.New("centroids: centroids cannot be nil")
	}

	cs := *centroids

	// Setup the centroids array.
	if len(cs) != n {
		// We should return n centroids.
		cs = make([]Vector, n)
	}
	if cs[0] == nil {
		// Each centroid should be the length of the data vector.
		for i := 0; i < n; i++ {
			cs[i] = make(Vector, len(data[0]))
		}
	}

	// Reset the centroid values.
	for _, v := range cs {
		for j := 0; j < len(v); j++ {
			v[j] = 0
		}
	}

	// Make each centroid index be the sum of data in that cluster.
	for i, v := range data {
		assignment := assignments[i]
		for j, vj := range v {
			cs[assignment][j] += vj
		}
	}

	// Divide by the length of data to get the average.
	for _, c := range cs {
		for i, f := range c {
			c[i] = f / float64(len(data))
		}
	}

	return
}
