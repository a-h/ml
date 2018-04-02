package clustering

import (
	"errors"
)

// Assign vectors to their clusters.
func Assign(data []Vector, assignment []int) ([]Cluster, error) {
	if len(assignment) != len(data) {
		return nil, errors.New("assignment must equal the amount of input")
	}
	clusters := 0
	for _, v := range assignment {
		if v > clusters {
			clusters = v
		}
	}
	op := make([]Cluster, clusters+1)
	for i, a := range assignment {
		existing := op[a]
		op[a] = append(existing, data[i])
	}
	return op, nil
}

// Cluster is a slice of vectors.
type Cluster []Vector

// Eq compares two clusters to determine whether they are equal.
func (c Cluster) Eq(o Cluster) bool {
	if len(c) != len(o) {
		return false
	}
	for i, v := range c {
		if !o[i].Eq(v) {
			return false
		}
	}
	return true
}

// Clusters is a slice of clusters.
type Clusters []Cluster

// Eq compares two slices to determine whether they are equal.
func (c Clusters) Eq(o Clusters) bool {
	if len(c) != len(o) {
		return false
	}
	for i, v := range c {
		if !o[i].Eq(v) {
			return false
		}
	}
	return true
}
