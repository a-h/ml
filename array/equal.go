package array

// EqFloat64 compares two arrays.
func EqFloat64(v1, v2 []float64) bool {
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
