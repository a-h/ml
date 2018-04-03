package distance

// MeanSquare calculates the mean square (RMS) distance between two input vectors.
func MeanSquare(p []float64, q []float64) (r float64, err error) {
	if err = validateInputs(p, q); err != nil {
		return
	}
	for i, pp := range p {
		qq := q[i]
		r += (qq - pp) * (qq - pp)
	}
	n := float64(len(p))
	r = (1.0 / n) * r
	return
}
