package rbf

// Function is a radial basis function, for use in an RBF network.
// Examples are rbf.NewGaussian and rbf.RickerWavelet.
type Function func(v float64) (r float64)

// VectorFunction is a radial basis function which operates on a vector.
// Examples are rbf.NewGaussianVector.
type VectorFunction func(v []float64) (r float64, err error)
