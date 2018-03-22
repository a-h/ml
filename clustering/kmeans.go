package clustering

import (
	"errors"
	"math/rand"
	"time"

	"github.com/a-h/ml/distance"
)

// KMeans cluster the input vectors into n clusters using the distance function d.
func KMeans(data []Vector, n int, d distance.Function) (assignment []int, err error) {
	if n <= 0 {
		return nil, errors.New("KMeans: n cannot be less than or equal to zero")
	}
	if data == nil {
		return nil, errors.New("KMeans: data cannot be nil")
	}
	if len(data) == 0 {
		return nil, errors.New("KMeans: data cannot be empty")
	}

	// Assign data to random clusters, but make sure every cluster has something in it.
	assignment = make([]int, len(data))

	r := rand.New(rand.NewSource(time.Now().Unix()))
	assigned := map[int]interface{}{}
	for i := 0; i < n; i++ {
		for {
			to := r.Intn(len(data))
			if _, ok := assigned[to]; !ok {
				assignment[to] = i
				assigned[to] = true
				break
			}
		}
	}
	for i := 0; i < len(data); i++ {
		if _, ok := assigned[i]; !ok {
			assignment[i] = r.Intn(n)
		}
	}

	// Create the centroids array once.
	centroids := make([]Vector, n)

	var done bool
	for {
		// Exit when done processing.
		if done {
			break
		}
		// Calculate / recalculate centroids.
		err = Centroids(data, n, assignment, &centroids)
		if err != nil {
			return assignment, err
		}
		done = true
		for i, v := range data {
			currentAssignmentIndex := assignment[i]
			newAssignmentIndex, err := findNearest(&v, &centroids, d)
			if err != nil {
				return assignment, err
			}
			if currentAssignmentIndex != newAssignmentIndex {
				assignment[i] = newAssignmentIndex
				done = false
			}
		}
	}
	return assignment, err
}

func findNearest(v *Vector, centroids *[]Vector, d distance.Function) (n int, err error) {
	if centroids == nil || len(*centroids) == 0 {
		err = errors.New("KMeans: no centroids provided")
		return
	}
	cs := *centroids

	// Calculate the first distance as a starting point.
	nd, err := d(*v, cs[0])
	if err != nil {
		return
	}
	// Do the rest.
	start := 1
	for i, centroid := range cs[start:] {
		var cd float64
		cd, err = d(*v, centroid)
		if err != nil {
			return
		}
		if cd < nd {
			nd = cd
			n = i + start // We started from [1:], so add one back in.
		}
	}
	return
}
