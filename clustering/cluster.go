package clustering

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

// Cluster is a cluster of data.
type Cluster struct {
	vectorLength int
	Data         map[string]Vector
	Centroid     Vector
}

func (c *Cluster) String() string {
	b := []byte("{ ")
	buf := bytes.NewBuffer(b)
	for _, v := range c.Data {
		buf.WriteString(fmt.Sprintf("%v, ", v))
	}
	buf.WriteString(fmt.Sprintf("centroid: %v }", c.Centroid))
	return buf.String()
}

// ClusterMember defines the member of a cluster.
type ClusterMember struct {
	ID     string
	Vector Vector
}

// NewClusterMember creates a ClusterMember from a vector.
func NewClusterMember(v Vector) (cm ClusterMember, err error) {
	bytes := make([]byte, 16)
	_, err = rand.Read(bytes)
	cm = ClusterMember{
		ID:     base64.StdEncoding.EncodeToString(bytes),
		Vector: v,
	}
	return
}

// NewCluster creates a cluster.
func NewCluster(vectorLength int) *Cluster {
	return &Cluster{
		Data:         map[string]Vector{},
		vectorLength: vectorLength,
		Centroid:     make(Vector, vectorLength),
	}
}

// VectorLength sets the expected length of any vectors which make up the cluster.
func (c *Cluster) checkVectorLength(v Vector) (err error) {
	if len(v) != c.vectorLength {
		err = fmt.Errorf("expected vector length %d, but got %d", c.vectorLength, len(v))
	}
	return
}

// Add adds a vector to the cluster.
func (c *Cluster) Add(v Vector) (err error) {
	m, err := NewClusterMember(v)
	if err != nil {
		return
	}
	return c.AddMembers(m)
}

// AddRange adds multiple vector to the cluster.
func (c *Cluster) AddRange(v ...Vector) (err error) {
	for _, vv := range v {
		m, err := NewClusterMember(vv)
		if err != nil {
			return err
		}
		err = c.AddMember(m)
		if err != nil {
			return err
		}
	}
	return
}

// AddMember adds a member to the cluster.
func (c *Cluster) AddMember(m ClusterMember) (err error) {
	return c.AddMembers(m)
}

// AddMembers adds more than one member to the cluster.
func (c *Cluster) AddMembers(m ...ClusterMember) (err error) {
	for _, mm := range m {
		err = c.checkVectorLength(mm.Vector)
		if err != nil {
			return
		}
		c.Data[mm.ID] = mm.Vector
	}
	c.Centroid = calculateCentroid(c.vectorLength, c.Data)
	return
}

// Remove a vector from the cluster.
func (c *Cluster) Remove(v ClusterMember) {
	delete(c.Data, v.ID)
	c.Centroid = calculateCentroid(c.vectorLength, c.Data)
}

// Centroid calculates the average position of the members in the cluster.
func calculateCentroid(length int, data map[string]Vector) Vector {
	var op = make([]float64, length)

	// Avoid divide by zero exceptions.
	if len(data) == 0 {
		return op
	}

	// Sum the data.
	for _, v := range data {
		for j, vj := range v {
			op[j] += vj
		}
	}

	// Divide by the length of data to get the average.
	for i, f := range op {
		op[i] = f / float64(len(data))
	}
	return op
}
