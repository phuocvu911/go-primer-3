package mylist

import "testing"

func valuesOf(l *List) []any {
	vals := make([]any, 0, l.Len())
	for e := l.Front(); e != nil; e = e.Next() {
		vals = append(vals, e.Value)
	}
	return vals
}

func TestList_NewAndInit(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() *List
		wantLen int
		wantVal []any
	}{
		{
			name: "new empty list",
			setup: func() *List {
				return New()
			},
			wantLen: 0,
			wantVal: nil,
		},
		{
			name: "init clears populated list",
			setup: func() *List {
				l := New()
				l.PushBack("a")
				l.PushBack("b")
				l.PushBack("c")
				return l
			},
			wantLen: 0,
			wantVal: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := tt.setup()
			if got := l.Init(); got != l {
				t.Fatalf("Init() = %p, want %p", got, l)
			}
			if got := l.Len(); got != tt.wantLen {
				t.Fatalf("Len() = %d, want %d", got, tt.wantLen)
			}
			if got := valuesOf(l); !equalAnySlice(got, tt.wantVal) {
				t.Fatalf("valuesOf(l) = %v, want %v", got, tt.wantVal)
			}
			if l.Front() != nil {
				t.Fatalf("Front() = %v, want nil", l.Front())
			}
		})
	}
}

func TestList_PushBack(t *testing.T) {
	tests := []struct {
		name   string
		input  []any
		after  []any
		wantFn func(*List)
	}{
		{
			name:  "empty list",
			input: []any{1},
			after: []any{1},
		},
		{
			name:  "append to existing list",
			input: []any{"a", "b", "c"},
			after: []any{"a", "b", "c"},
		},
		{
			name:  "nil values are allowed",
			input: []any{nil, "x", nil},
			after: []any{nil, "x", nil},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New()
			for _, v := range tt.input {
				e := l.PushBack(v)
				if e == nil {
					t.Fatalf("PushBack(%v) returned nil element", v)
				}
				if e.Value != v {
					t.Fatalf("element.Value = %v, want %v", e.Value, v)
				}
			}
			if got, want := l.Len(), len(tt.after); got != want {
				t.Fatalf("Len() = %d, want %d", got, want)
			}
			if got := valuesOf(l); !equalAnySlice(got, tt.after) {
				t.Fatalf("valuesOf(l) = %v, want %v", got, tt.after)
			}
		})
	}
}

func TestList_InsertAfter(t *testing.T) {
	tests := []struct {
		name    string
		before  []any
		markIdx int
		value   any
		want    []any
		wantNil bool
	}{
		{
			name:    "insert in middle",
			before:  []any{"a", "b", "c"},
			markIdx: 1,
			value:   "x",
			want:    []any{"a", "b", "x", "c"},
		},
		{
			name:    "insert after tail",
			before:  []any{"a", "b"},
			markIdx: 1,
			value:   "c",
			want:    []any{"a", "b", "c"},
		},
		{
			name:    "nil mark returns nil",
			before:  []any{"a", "b"},
			markIdx: -1,
			value:   "x",
			want:    []any{"a", "b"},
			wantNil: true,
		},
		{
			name:    "mark from other list returns nil",
			before:  []any{"a", "b"},
			markIdx: 0,
			value:   "x",
			want:    []any{"a", "b"},
			wantNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New()
			for _, v := range tt.before {
				l.PushBack(v)
			}

			var mark *Element
			if tt.markIdx >= 0 {
				mark = l.Front()
				for i := 0; i < tt.markIdx && mark != nil; i++ {
					mark = mark.Next()
				}
			}

			if tt.name == "mark from other list returns nil" {
				other := New()
				mark = other.PushBack("z")
			}

			got := l.InsertAfter(tt.value, mark)
			if tt.wantNil {
				if got != nil {
					t.Fatalf("InsertAfter(%v, mark) = %v, want nil", tt.value, got)
				}
				if gotVals := valuesOf(l); !equalAnySlice(gotVals, tt.want) {
					t.Fatalf("valuesOf(l) = %v, want %v", gotVals, tt.want)
				}
				return
			}

			if got == nil {
				t.Fatalf("InsertAfter(%v, mark) returned nil, want new element", tt.value)
			}
			if got.Value != tt.value {
				t.Fatalf("InsertAfter(%v, mark).Value = %v, want %v", tt.value, got.Value, tt.value)
			}
			if gotVals := valuesOf(l); !equalAnySlice(gotVals, tt.want) {
				t.Fatalf("valuesOf(l) = %v, want %v", gotVals, tt.want)
			}
		})
	}
}

func TestList_Remove(t *testing.T) {
	tests := []struct {
		name      string
		before    []any
		removeIdx int
		wantVal   any
		wantAfter []any
		wantNil   bool
		otherList bool
	}{
		{
			name:      "remove head from multi-item list",
			before:    []any{"a", "b", "c"},
			removeIdx: 0,
			wantVal:   "a",
			wantAfter: []any{"b", "c"},
		},
		{
			name:      "remove middle element",
			before:    []any{"a", "b", "c"},
			removeIdx: 1,
			wantVal:   "b",
			wantAfter: []any{"a", "c"},
		},
		{
			name:      "remove tail",
			before:    []any{"a", "b", "c"},
			removeIdx: 2,
			wantVal:   "c",
			wantAfter: []any{"a", "b"},
		},
		{
			name:      "remove only element",
			before:    []any{"a"},
			removeIdx: 0,
			wantVal:   "a",
			wantAfter: nil,
		},
		{
			name:      "remove nil element",
			before:    []any{"a"},
			removeIdx: -1,
			wantNil:   true,
			wantAfter: []any{"a"},
		},
		{
			name:      "remove from other list",
			before:    []any{"a"},
			removeIdx: 0,
			otherList: true,
			wantVal:   "a",
			wantAfter: []any{"a"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New()
			for _, v := range tt.before {
				l.PushBack(v)
			}

			var e *Element
			if tt.removeIdx >= 0 {
				e = l.Front()
				for i := 0; i < tt.removeIdx && e != nil; i++ {
					e = e.Next()
				}
			}

			if tt.otherList {
				other := New()
				e = other.PushBack(tt.before[0])
			}

			if tt.wantNil {
				if got := l.Remove(nil); got != nil {
					t.Fatalf("Remove(nil) = %v, want nil", got)
				}
				if gotVals := valuesOf(l); !equalAnySlice(gotVals, tt.wantAfter) {
					t.Fatalf("valuesOf(l) = %v, want %v", gotVals, tt.wantAfter)
				}
				return
			}

			got := l.Remove(e)
			if got != tt.wantVal {
				t.Fatalf("Remove(e) = %v, want %v", got, tt.wantVal)
			}
			if gotVals := valuesOf(l); !equalAnySlice(gotVals, tt.wantAfter) {
				t.Fatalf("valuesOf(l) = %v, want %v", gotVals, tt.wantAfter)
			}
		})
	}
}

func equalAnySlice(a, b []any) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
