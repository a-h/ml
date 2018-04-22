package training

// Trainee defines the behaviour something needs to provide if it can be trained.
// Examples include rbf.Network and rbf.Node.
type Trainee interface {
	Calculate(input []float64) (output []float64, err error)
	GetMemorySize() int
	GetMemory() []float64
	SetMemory(m []float64)
}

// Algorithm defines a training algorith, e.g. RandomGreedy.
type Algorithm interface {
	Next(e float64) []float64
	BestMemory() (memory []float64)
	BestError() (e float64)
}

// Data is data used to train something.
type Data struct {
	Input    []float64
	Expected []float64
}
