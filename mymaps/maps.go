package mymaps

// Equal reports if 2 maps contains the same k-v pairs.
func Equal(m1, m2 map[string]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for key1, val1 := range m1 {
		val2, ok := m2[key1]
		if !ok || val1 != val2 {
			return false
		}
	}
	return true
}

// Clone returns an independent copy of m.
func Clone(m map[string]int) map[string]int {
	if m == nil {
		return nil
	}
	res := make(map[string]int)
	for key, val := range m {
		res[key] = val
	}
	return res
}

// Clears empty m, left with usable m.
func Clear(m map[string]int) {
	if m == nil {
		return
	}

	for key := range m {
		delete(m, key)
	}
}
