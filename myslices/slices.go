package myslices

// Equal reports if s1 and s2 are equal in terms of length and element.
func IntsEqual(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

// IntsClone make a copy of s, share no storage.
func IntsClone(s []int) []int {
	if s == nil {
		return nil
	}
	res := make([]int, len(s))
	copy(res, s)
	return res
}

// IntsInsert inserts the values v into s at index i.
func IntsInsert(s []int, i int, v []int) ([]int, bool) {
	if i < 0 || i > len(s) {
		return nil, false
	}

	res := make([]int, len(s)+len(v))
	copy(res, s[:i])            // copy the first part of s into res
	copy(res[i:], v)            //now the v
	copy(res[i+len(v):], s[i:]) //rest of s
	return res, true
}

// IntsMax finds the max value in s.
func IntsMax(s []int) (int, bool) {
	if len(s) == 0 {
		return 0, false
	}

	res := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] > res {
			res = s[i]
		}
	}
	return res, true
}

// IntsDeleteFunc deletes elements from s inplace for which del returns true.
func IntsDeleteFunc(s []int, del func(int) bool) []int {
	i := 0
	//find elems to keep, and increase the new tail accordinglys
	for _, v := range s {
		if !del(v) {
			s[i] = v
			i++
		}
	}
	clear(s[i:]) // zero the leftover tail so nothing stale remains
	return s[:i]
}

// InsFilter keeps elements in s for which keep returns true.
func IntsFilter(s []int, keep func(int) bool) []int {
	if len(s) == 0 {
		return []int{}
	}
	res := make([]int, 0, len(s))
	for _, v := range s {
		if keep(v) {
			res = append(res, v)
		}
	}
	return res
}

// IntsMap applies f to each element of s and returns a new slice with the results.
func IntsMap(s []int, f func(int) int) []int {
	if len(s) == 0 {
		return []int{}
	}
	res := make([]int, len(s))
	for i, v := range s {
		res[i] = f(v)
	}
	return res
}
