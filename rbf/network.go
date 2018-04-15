package rbf

import (
	"errors"
	"fmt"

	"github.com/a-h/ml/distance"
)

// Neuron defines an RBF neuron in an RBF network.
type Neuron struct {
	InputWeights  []float64
	Centroid      []float64
	Distance      distance.Function
	RBF           Function
	OutputWeights []float64
}

// Calculate the output of the neuron.
func (n Neuron) Calculate(input []float64) (op []float64, err error) {
	if len(n.InputWeights) != len(input) {
		err = fmt.Errorf("rbf: the input vector has a length of %d values and should have the same number of input weights, but we have %d neuron input weights",
			len(input), len(n.InputWeights))
		return
	}

	// Scale the input against the neuron's weights
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
	output := n.RBF(d)
	op = make([]float64, len(n.OutputWeights))
	for i, outputWeight := range n.OutputWeights {
		op[i] = output * outputWeight
	}
	return
}

// OutputCount returns the number of output nodes.
func (n Neuron) OutputCount() int {
	return len(n.OutputWeights)
}

// NewBias creates a bias node with the specified number of outputs.
func NewBias(count int) (b Bias) {
	b.Outputs = make([]float64, count)
	for i := 0; i < len(b.Outputs); i++ {
		b.Outputs[i] = 1.0
	}
	return
}

// Bias is a node with constant output.
type Bias struct {
	Outputs []float64
}

// Calculate returns the Output values regardless of input.
func (b Bias) Calculate(input []float64) (op []float64, err error) {
	return b.Outputs, nil
}

// OutputCount returns the number of output nodes.
func (b Bias) OutputCount() int {
	return len(b.Outputs)
}

// Node defines the behaviour required for an RBF node to function.
type Node interface {
	Calculate(input []float64) (output []float64, err error)
	OutputCount() int
}

// NewNetwork creates a new network from the input nodes, checking that the configuration
// matches.
func NewNetwork(neurons ...Node) (n Network, err error) {
	if len(neurons) == 0 {
		err = errors.New("rbf: Unable to create network, since there are no neurons")
		return
	}
	count := neurons[0].OutputCount()
	for i, n := range neurons[1:] {
		if n.OutputCount() != count {
			err = fmt.Errorf("rbf: Mismatched neuron output length, neuron %d has %d output nodes, but expected %d outputs",
				i, n.OutputCount(), count)
		}
	}
	n = Network(neurons)
	return
}

// Network defines the nodes in an RBF network.
type Network []Node

// Calculate the output of the network.
func (neurons Network) Calculate(input []float64) (op []float64, err error) {
	if len(neurons) == 0 {
		err = errors.New("rbf: Unable to calculate result for RBF network, since there are no neurons")
		return
	}

	var configured bool

	for i, n := range neurons {
		// Calculate the neuron's output.
		var nv []float64
		nv, err = n.Calculate(input)
		if err != nil {
			return
		}
		// Initialise the network output if required.
		if !configured {
			op = make([]float64, len(nv))
			configured = true
		}
		// Check that the output node count matches the expected count.
		if len(nv) != len(op) {
			err = fmt.Errorf("rbf: The RBF has been configured with %d output nodes, but neuron %d has %d output nodes",
				len(op), i, len(nv))
			return
		}
		// Update the output weights.
		for i, nnv := range nv {
			op[i] += nnv
		}
	}
	return
}
