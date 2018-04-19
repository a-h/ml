package rbf

// ExecutableNode defines the behaviour required for an RBF node to function.
type ExecutableNode interface {
	Calculate(input []float64) (output []float64, err error)
	OutputCount() int
}

// TrainableNode defines the behaviour of a node which can be trained.
type TrainableNode interface {
	GetMemorySize() int
	GetMemory() []float64
	SetMemory(m []float64)
}
