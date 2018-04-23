package rbf

import (
	"encoding/json"
	"fmt"

	"github.com/a-h/ml/random"
)

// NewNode returns a new Node.
func NewNode(inputCount int, outputCount int) *Node {
	return &Node{
		InputWeights:  random.Float64Vector(-10, 10, inputCount),
		Centroid:      random.Float64Vector(-10, 10, inputCount),
		Width:         random.Float64(-10, 10),
		OutputWeights: random.Float64Vector(-10, 10, outputCount),
	}
}

// Node defines a node in an RBF network which uses a distance function to calculate distance.
type Node struct {
	InputWeights []float64
	Centroid     []float64
	// RBF function parameters.
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

	// Scale the distance using the RBF function then multiply by the scalar output weights.
	output, err := NewGaussianVector(1.0, n.Centroid, n.Width)(scaledInput)
	if err != nil {
		err = fmt.Errorf("rbf: could not calculate gaussian RBF: %v", err)
	}
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
	return len(n.InputWeights) + 1 + len(n.OutputWeights) // 1 = a value for n.Width
}

// GetMemory returns the node's internal state as an array.
func (n *Node) GetMemory() (op []float64) {
	//TODO: Benchmark this approach and check that it's OK. I think it is given Go's slice internals.
	op = append(op, n.InputWeights...)
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
	n.Width = m[index]
	index++
	n.OutputWeights = m[index:]
}
