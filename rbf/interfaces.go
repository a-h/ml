package rbf

// ExecutableNode defines the behaviour required for an RBF node to function.
type ExecutableNode interface {
	Calculate(input []float64) (output []float64, err error)
	OutputCount() int
}

// Trainable defines the behaviour of an item (e.g. node, network) which can be trained.
type Trainable interface {
	GetMemorySize() int
	GetMemory() []float64
	SetMemory(m []float64)
}
