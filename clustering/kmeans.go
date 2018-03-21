package clustering

import (
	"errors"
	"math/rand"

	"github.com/a-h/ml/distance"
)

// KMeans cluster the input vectors into n clusters using the distance function d.
func KMeans(data []Vector, n int, d distance.Function) (clusters []*Cluster, err error) {
	if n <= 0 {
		return nil, errors.New("KMeans: n cannot be less than or equal to zero")
	}
	if data == nil {
		return nil, errors.New("KMeans: data cannot be nil")
	}
	if len(data) == 0 {
		return nil, errors.New("KMeans: data cannot be empty")
	}

	clusters = make([]*Cluster, n)
	for i := 0; i < n; i++ {
		clusters[i] = NewCluster(len(data[0]))
	}

	// Initially assign all of the vectors to a random cluster.
	randomlyAssign(clusters, data)

	for {
		done := true
		for _, source := range clusters {
			for id, v := range source.Data {
				cm := ClusterMember{
					ID:     id,
					Vector: v,
				}
				target, err := findNearest(cm.Vector, clusters, d)
				if err != nil {
					return clusters, err
				}
				if target != source {
					source.Remove(cm)
					target.AddMember(cm)
					done = false
				}
			}
		}
		if done {
			break
		}
	}
	return clusters, err
}

func randomlyAssign(cs []*Cluster, vectors []Vector) {
	for _, v := range vectors {
		r := rand.Intn(len(cs))
		cs[r].Add(v)
	}
}

func findNearest(v Vector, cs []*Cluster, d distance.Function) (nearest *Cluster, err error) {
	if len(cs) == 0 {
		err = errors.New("KMeans: no clusters provided")
		return
	}
	var nd float64
	nearestIndex := -1
	for i := 0; i < len(cs); i++ {
		var cd float64
		cd, err = d(v, cs[i].Centroid)
		if err != nil {
			return
		}
		if cd < nd || nearestIndex < 0 {
			nd = cd
			nearestIndex = i
		}
	}
	nearest = cs[nearestIndex]
	return
}
