package mysort

import (
	"math"
	"reflect"
	"slices"
	"testing"
)

type StringSlice []string

func (s StringSlice) Len() int           { return len(s) }
func (s StringSlice) Less(i, j int) bool { return len(s[i]) < len(s[j]) }
func (s StringSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func TestSort(t *testing.T) {
	tests := []struct {
		name string
		data Interface
		want Interface
	}{
		{
			name: "integer slice",
			data: IntSlice{3, 1, 2},
			want: IntSlice{1, 2, 3},
		},
		{
			name: "float slice",
			data: Float64Slice{3.5, -1, 2},
			want: Float64Slice{-1, 2, 3.5},
		},
		{
			name: "empty slice",
			data: IntSlice{},
			want: IntSlice{},
		},
		{
			name: "nil slice",
			data: IntSlice(nil),
			want: IntSlice(nil),
		},
		{
			name: "string slice",
			data: StringSlice{"Hive", "Helsinki", "Hiver9"},
			want: StringSlice{"Hive", "Hiver9", "Helsinki"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Sort(tt.data)
			if !reflect.DeepEqual(tt.data, tt.want) {
				t.Errorf("Sort() got = %v, want %v", tt.data, tt.want)
			}
		})
	}
}

func TestInts(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "unsorted values",
			input: []int{3, 1, 2},
			want:  []int{1, 2, 3},
		},
		{
			name:  "duplicates and negative values",
			input: []int{2, -1, 2, 0},
			want:  []int{-1, 0, 2, 2},
		},
		{
			name:  "already sorted",
			input: []int{-2, 0, 5},
			want:  []int{-2, 0, 5},
		},
		{
			name:  "nil slice",
			input: nil,
			want:  nil,
		},
		{
			name:  "empty non-nil slice",
			input: []int{},
			want:  []int{},
		},
		{
			name:  "single value",
			input: []int{1},
			want:  []int{1},
		},
		{
			name:  "integer boundaries",
			input: []int{math.MaxInt, 0, math.MinInt},
			want:  []int{math.MinInt, 0, math.MaxInt},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := slices.Clone(tt.input)
			Ints(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ints() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64s(t *testing.T) {
	tests := []struct {
		name  string
		input []float64
		want  []float64
		todo  string
	}{
		{
			name:  "unsorted values",
			input: []float64{3.5, -1, 2},
			want:  []float64{-1, 2, 3.5},
		},
		{
			name:  "duplicates and negative values",
			input: []float64{2, -1, 2, 0},
			want:  []float64{-1, 0, 2, 2},
		},
		{
			name:  "already sorted",
			input: []float64{-2, 0, 5},
			want:  []float64{-2, 0, 5},
		},
		{
			name:  "nil slice",
			input: nil,
			want:  nil,
		},
		{
			name:  "empty non-nil slice",
			input: []float64{},
			want:  []float64{},
		},
		{
			name:  "single value",
			input: []float64{1},
			want:  []float64{1},
		},
		{
			name:  "infinities, signed zero, and finite boundaries",
			input: []float64{math.Inf(1), math.SmallestNonzeroFloat64, math.Copysign(0, -1), math.Inf(-1), math.MaxFloat64, 0},
			want:  []float64{math.Inf(-1), math.Copysign(0, -1), 0, math.SmallestNonzeroFloat64, math.MaxFloat64, math.Inf(1)},
		},
		{
			name:  "NaN ordering (TODO)",
			input: []float64{1, math.NaN(), -1},
			want:  []float64{math.NaN(), -1, 1},
			todo:  "TODO: decide the intended ordering of NaN relative to finite values",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.todo != "" {
				t.Skip(tt.todo)
			}

			got := slices.Clone(tt.input)
			Float64s(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float64s() got = %v, want %v", got, tt.want)
			}
		})
	}
}
