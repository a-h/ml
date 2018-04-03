package distance

import (
	"errors"
)

// ErrMismatchedVectorLengths is an error for when an operation cannot be carried out due to different sizes
// of vectors being used, e.g. 2 vectors, one with length 2 and the other of length 3 cannot be compared.
var ErrMismatchedVectorLengths = errors.New("distance: mismatched vector lengths")

// ErrZeroLengthVector is an error for when an input has a zero length vector - i.e. there is nothing to compare.
var ErrZeroLengthVector = errors.New("distance: zero length vector")

// ErrNilVector is an error for when an input has a zero length vector - i.e. there is nothing to compare.
var ErrNilVector = errors.New("distance: nil vector")

func validateInputs(p, q []float64) error {
	if p == nil || q == nil {
		return ErrNilVector
	}
	if len(p) != len(q) {
		return ErrMismatchedVectorLengths
	}
	if len(p) == 0 || len(q) == 0 {
		return ErrZeroLengthVector
	}
	return nil
}
