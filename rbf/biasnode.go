package rbf

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
