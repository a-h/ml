package distance

// Function is the interface for all distance measurements.
type Function func(p []float64, q []float64) (d float64, err error)
