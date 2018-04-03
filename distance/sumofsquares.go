package distance

// SumOfSquares calculates the sum of squares between two input vectors.
func SumOfSquares(p []float64, q []float64) (r float64, err error) {
	if err = validateInputs(p, q); err != nil {
		return
	}
	for i, pp := range p {
		qq := q[i]
		r += (qq - pp) * (qq - pp)
	}
	return
}
