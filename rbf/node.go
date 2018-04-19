package rbf

import (
	"encoding/json"
	"fmt"

	"github.com/a-h/ml/distance"
	"github.com/a-h/ml/random"
)

// NewNode returns a new Node.
func NewNode(inputCount int, outputCount int) *Node {
	return &Node{
		InputWeights:  random.Float64Vector(-100, 100, inputCount),
		Centroid:      random.Float64Vector(-100, 100, inputCount),
		Distance:      distance.Euclidean,
		Center:        random.Float64(-100, 100),
		Width:         random.Float64(-100, 100),
		OutputWeights: random.Float64Vector(-100, 100, outputCount),
	}
}

// Node defines a node in an RBF network which uses a distance function to calculate distance.
type Node struct {
	InputWeights []float64
	Centroid     []float64
	Distance     distance.Function `json:"-"`
	// RBF function parameters.
	Center        float64
	Width         float64
	OutputWeights []float64
}

func (n Node) String() string {
	b, err := json.Marshal(n)
	if err != nil {
		return fmt.Sprintf("rbf.Node: error marshalling to JSON: %v", err)
	}
	return string(b)
}

// Calculate the output of the node.
func (n *Node) Calculate(input []float64) (op []float64, err error) {
	if len(n.InputWeights) != len(input) {
		err = fmt.Errorf("rbf: the input vector has a length of %d values and should have the same number of input weights, but we have %d node input weights",
			len(input), len(n.InputWeights))
		return
	}

	// Scale the input against the node's weights
	scaledInput := make([]float64, len(input))
	for i, iv := range input {
		scaledInput[i] = iv * n.InputWeights[i]
	}

	// Calculate the distance against the node's center vector using the specified distance function.
	var d float64
	d, err = n.Distance(scaledInput, n.Centroid)
	if err != nil {
		return
	}

	// Scale the distance using the RBF function then multiply by the scalar output weights.
	output := NewGaussian(1.0, n.Width, n.Center)(d)
	op = make([]float64, len(n.OutputWeights))
	for i, outputWeight := range n.OutputWeights {
		op[i] = output * outputWeight
	}
	return
}

// OutputCount returns the number of output nodes.
func (n *Node) OutputCount() int {
	return len(n.OutputWeights)
}

// GetMemorySize returns the size of the node's internal state.
func (n *Node) GetMemorySize() int {
	return len(n.InputWeights) + 2 + len(n.OutputWeights) // 2 = a value for n.Center and n.Width

}

// GetMemory returns the node's internal state as an array.
func (n *Node) GetMemory() (op []float64) {
	//TODO: Benchmark this approach and check that it's OK. I think it is given Go's slice internals.
	op = append(op, n.InputWeights...)
	op = append(op, n.Center)
	op = append(op, n.Width)
	op = append(op, n.OutputWeights...)
	return
}

// SetMemory updates the node's internal state.
func (n *Node) SetMemory(m []float64) {
	//TODO: Add error handling here, check the lengths etc. Lots of opportunities to panic here.
	var index int
	n.InputWeights = m[index:len(n.InputWeights)]
	index = len(n.InputWeights)
	n.Center = m[index]
	index++
	n.Width = m[index]
	index++
	n.OutputWeights = m[index:]
}
