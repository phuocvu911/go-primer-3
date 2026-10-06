package myslices

import "testing"

func TestIntsEqual(t *testing.T) {
	tests := []struct {
		name string
		s1   []int
		s2   []int
		want bool
	}{
		{"empty and nil are equal", nil, []int{}, true},
		{"same values compare equal", []int{1, 2, 3, 4}, []int{1, 2, 3, 4}, true},
		{"different lengths compare unequal", []int{1, 2, 3}, []int{1, 2, 3, 4}, false},
		{"different elements compare unequal", []int{1, 2, 3}, []int{1, 9, 3}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IntsEqual(tt.s1, tt.s2)
			if got != tt.want {
				t.Fatalf("IntsEqual(%#v, %#v) = %v, expected %v", tt.s1, tt.s2, got, tt.want)
			}
		})
	}
}

func TestIntsClone(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{"clone copies values", []int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		{"nil input clone to nil", nil, nil},
		{"empty input clone to a non-nil one", []int{}, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clone := IntsClone(tt.input)
			if !IntsEqual(clone, tt.want) {
				t.Fatalf("IntsClone(%#v) = %#v, want %#v", tt.input, clone, tt.want)
			}
		})
	}

	t.Run("clone is independent of original 1", func(t *testing.T) {
		original := []int{10, 20, 30}
		clone := IntsClone(original)
		clone[0] = 99
		if original[0] != 10 {
			t.Fatalf("clone should not reference original: original=%#v clone=%#v", original, clone)
		}
	})

	t.Run("clone is independent of original 2", func(t *testing.T) {
		original := []int{10, 20, 30}
		clone := IntsClone(original)
		original[1] = 67
		if clone[1] != 20 {
			t.Fatalf("clone should not reference original: original=%#v clone=%#v", original, clone)
		}
	})
}

func TestIntsInsert(t *testing.T) {
	tests := []struct {
		name  string
		s     []int
		index int
		v     []int
		want  []int
		ok    bool
	}{
		{"insert in middle", []int{1, 2, 3}, 1, []int{9, 10}, []int{1, 9, 10, 2, 3}, true},
		{"insert at front", []int{1, 2, 3}, 0, []int{9, 10}, []int{9, 10, 1, 2, 3}, true},
		{"insert at end", []int{1, 2, 3}, 3, []int{9, 10}, []int{1, 2, 3, 9, 10}, true},
		{"insert into nil is fine", nil, 0, []int{9, 10}, []int{9, 10}, true},
		{"empty inserted values are fine", []int{1, 2}, 1, nil, []int{1, 2}, true},
		{"negative index", []int{1, 2, 3}, -1, []int{9}, nil, false},
		{"out-of-range index", []int{1, 2, 3}, 10, []int{9}, nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := IntsInsert(tt.s, tt.index, tt.v)
			if ok != tt.ok || !IntsEqual(got, tt.want) {
				t.Fatalf("IntsInsert(%#v, %d, %#v) = %#v, %v; want %#v, %v", tt.s, tt.index, tt.v, got, ok, tt.want, tt.ok)
			}
		})
	}

	t.Run("s is never modified", func(t *testing.T) {
		original := []int{1, 2, 3}
		_, _ = IntsInsert(original, 1, []int{9, 8})
		if !IntsEqual(original, []int{1, 2, 3}) {
			t.Fatalf("original slice should not be modified: got %#v", original)
		}
	})

	t.Run("s(nil) is never modified", func(t *testing.T) {
		s := IntsClone(nil)
		_, _ = IntsInsert(s, 0, []int{9, 10})
		if s != nil {
			t.Fatalf("insert into nil should not modify original slice: got %#v", s)
		}
	})

	t.Run("s is never modified, including the part of its array beyond len(s)", func(t *testing.T) {
		// Create a slice with capacity > length to test the area beyond len
		original := make([]int, 3, 5)
		original[0], original[1], original[2] = 1, 2, 3

		// Set values in the capacity area beyond length
		beyond := original[:cap(original)]
		beyond[3], beyond[4] = 67, 69

		_, _ = IntsInsert(original, 1, []int{9, 8})

		// Check that the capacity area beyond len is also unchanged
		if beyond[3] != 67 || beyond[4] != 69 {
			t.Fatalf("capacity area beyond len(s) was modified: got [%d, %d] at indices [3,4], want [67, 69]", beyond[3], beyond[4])
		}
	})
}

func TestIntsMax(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
		ok    bool
	}{
		{"basic maximum", []int{3, 9, 1, 7, 5}, 9, true},
		{"single element", []int{-4}, -4, true},
		{"all negatives", []int{-12, -3, -9}, -3, true},
		{"zero is max", []int{-67, -3, 0, -11}, 0, true},
		{"nil input", nil, 0, false},
		{"empty input", []int{}, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			max, ok := IntsMax(tt.input)
			if ok != tt.ok || max != tt.want {
				t.Fatalf("IntsMax(%#v) = %d, %v; want %d, %v", tt.input, max, ok, tt.want, tt.ok)
			}
		})
	}
}

func TestIntsDeleteFunc(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		del   func(int) bool
		want  []int
	}{
		{"delete matching elements while preserving order", []int{1, 2, 3, 4, 5, 6}, func(v int) bool { return v%2 == 0 }, []int{1, 3, 5}},
		{"delete nothing ", []int{10, 20, 30}, func(v int) bool { return v > 100 }, []int{10, 20, 30}},
		{"nil input", nil, func(v int) bool { return true }, nil},
		{"delete all results in empty slice", []int{1, 2, 3}, func(v int) bool { return true }, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IntsDeleteFunc(tt.input, tt.del)
			if !IntsEqual(got, tt.want) {
				t.Fatalf("IntsDeleteFunc failed: got %#v want %#v", got, tt.want)
			}
		})
	}

	t.Run("work in place and zeroes the positions beyond the new length", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5}
		del := func(v int) bool { return v%2 == 0 }
		result := IntsDeleteFunc(s, del)

		expected := []int{1, 3, 5}
		if !IntsEqual(result, expected) {
			t.Fatalf("delete func failed: got %#v want %#v", result, expected)
		}

		// Check that the original slice has been modified in place
		if !IntsEqual(s[:len(result)], expected) {
			t.Fatalf("original slice not modified in place: got %#v want %#v", s[:len(result)], expected)
		}

		// Check that the positions beyond the new length are zeroed
		for i := len(result); i < len(s); i++ {
			if s[i] != 0 {
				t.Fatalf("position %d beyond new length not zeroed: got %d want 0", i, s[i])
			}
		}
	})
}

func TestIntsFilter(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		keep  func(int) bool
		want  []int
	}{
		{"keep matching values", []int{-2, -1, 0, 1, 2, 3}, func(v int) bool { return v > 0 }, []int{1, 2, 3}},
		{"keep nothing", []int{1, 2, 3}, func(v int) bool { return false }, []int{}},
		{"keep everything", []int{1, 2, 3}, func(v int) bool { return true }, []int{1, 2, 3}},
		{"nil input", nil, func(v int) bool { return true }, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IntsFilter(tt.input, tt.keep)
			if !IntsEqual(got, tt.want) {
				t.Fatalf("filter failed: got %#v want %#v", got, tt.want)
			}
		})
	}

	t.Run("s is not modified", func(t *testing.T) {
		original := []int{1, 2, 3, 4, 5}
		keep := func(v int) bool { return v%2 == 0 }
		result := IntsFilter(original, keep)

		want := []int{2, 4}
		if !IntsEqual(result, want) {
			t.Fatalf("filter failed: got %#v want %#v", result, want)
		}

		// Check that the original slice has not been modified
		if !IntsEqual(original, []int{1, 2, 3, 4, 5}) {
			t.Fatalf("original slice should not be modified: got %#v want %#v", original, []int{1, 2, 3, 4, 5})
		}
	})
}

func TestIntsMap(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		f     func(int) int
		want  []int
	}{
		{"map transforms each element", []int{1, 2, 3, 4}, func(v int) int { return v * v }, []int{1, 4, 9, 16}},
		{"identity map preserves values", []int{10, 20, 30}, func(v int) int { return v }, []int{10, 20, 30}},
		{"nil input", nil, func(v int) int { return v + 1 }, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IntsMap(tt.input, tt.f)
			if !IntsEqual(got, tt.want) {
				t.Fatalf("map failed: got %#v want %#v", got, tt.want)
			}
		})
	}

	t.Run("s is not modified", func(t *testing.T) {
		original := []int{1, 2, 3, 4, 5}
		f := func(v int) int { return v * 2 }
		result := IntsMap(original, f)

		want := []int{2, 4, 6, 8, 10}
		if !IntsEqual(result, want) {
			t.Fatalf("map failed: got %#v want %#v", result, want)
		}

		// Check that the original slice has not been modified
		if !IntsEqual(original, []int{1, 2, 3, 4, 5}) {
			t.Fatalf("original slice should not be modified: got %#v want %#v", original, []int{1, 2, 3, 4, 5})
		}
	})
}
