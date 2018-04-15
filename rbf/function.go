package rbf

// Function is a radial basis function, for use in an RBF network.
// Examples are rbf.Gaussian and rbf.RickerWavelet.
//TODO: Update the function to accept a width parameter.
type Function func(v float64) (r float64)
