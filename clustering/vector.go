package clustering

// Vector is an array of numbers and its associated length.
type Vector []float64

// Eq compares two arrays.
func (v1 Vector) Eq(v2 Vector) bool {
	if len(v1) != len(v2) {
		return false
	}

	for i, v1i := range v1 {
		if v1i != v2[i] {
			return false
		}
	}

	return true
}
