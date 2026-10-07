package mymaps

import (
	"reflect"
	"testing"
)

func TestEqual(t *testing.T) {
	tests := []struct {
		name string
		m1   map[string]int
		m2   map[string]int
		want bool
	}{
		{"both nil", nil, nil, true},
		{"nil and empty are equal", nil, map[string]int{}, true},
		{"empty and empty are equal", map[string]int{}, map[string]int{}, true},
		{"same entries compare equal regardless of key order", map[string]int{"a": 1, "b": 2, "c": 3}, map[string]int{"c": 3, "a": 1, "b": 2}, true},
		{"zero values are preserved", map[string]int{"a": 0, "b": -1}, map[string]int{"b": -1, "a": 0}, true},
		{"different values compare unequal", map[string]int{"a": 1}, map[string]int{"a": 2}, false},
		{"different keys compare unequal", map[string]int{"a": 1}, map[string]int{"b": 1}, false},
		{"different lengths compare unequal", map[string]int{"a": 1, "b": 2}, map[string]int{"a": 1}, false},
		{"key with 0 different from absent key", map[string]int{"a": 1, "b": 0}, map[string]int{"a": 1}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Equal(tt.m1, tt.m2)
			if got != tt.want {
				t.Fatalf("Equal(%#v, %#v) = %v, want %v", tt.m1, tt.m2, got, tt.want)
			}
		})
	}
}

func TestClone(t *testing.T) {
	tests := []struct {
		name  string
		input map[string]int
		want  map[string]int
	}{
		{"nil clones to nil", nil, nil},
		{"empty map clones to empty map", map[string]int{}, map[string]int{}},
		{"non-empty input clone keeps values", map[string]int{"a": 1, "b": 2, "c": 3}, map[string]int{"a": 1, "b": 2, "c": 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clone := Clone(tt.input)
			if !reflect.DeepEqual(clone, tt.want) {
				t.Fatalf("Clone(%#v) = %#v, want %#v", tt.input, clone, tt.want)
			}
		})
	}

	t.Run("clone is independent of original", func(t *testing.T) {
		original := map[string]int{"a": 1, "b": 2}
		clone := Clone(original)
		clone["a"] = 99
		if original["a"] != 1 {
			t.Fatalf("clone should not reference original: original=%#v clone=%#v", original, clone)
		}
	})

	t.Run("mutating original after clone does not mutate clone", func(t *testing.T) {
		original := map[string]int{"a": 1, "b": 2}
		clone := Clone(original)
		original["b"] = 42
		if clone["b"] != 2 {
			t.Fatalf("clone should be independent: original=%#v clone=%#v", original, clone)
		}
	})
}

func TestClear(t *testing.T) {
	tests := []struct {
		name  string
		input map[string]int
		want  map[string]int
	}{
		{"non-empty map is emptied", map[string]int{"a": 1, "b": 2, "c": 3}, map[string]int{}},
		{"name: empty map remains empty", map[string]int{}, map[string]int{}},
		{"name: nil map remains nil", nil, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Clear(tt.input)
			if !reflect.DeepEqual(tt.input, tt.want) {
				t.Fatalf("Clear(%#v) left %#v, want %#v", tt.input, tt.input, tt.want)
			}
		})
	}

	t.Run("clear on nil map does not panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Fatalf("Clear(nil) panicked: %v", r)
			}
		}()
		Clear(nil)
	})
}
