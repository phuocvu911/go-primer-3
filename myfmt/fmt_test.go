package myfmt

import (
	"errors"
	"testing"

	"goprimer/mylist"
)

type Hoang string
type myErr string

func (e myErr) Error() string {
	return string(e)
}

func TestSprint(t *testing.T) {
	errPing := errors.New("ping")

	tests := []struct {
		name  string
		input any
		want  string
	}{
		{"nil interface", nil, "<nil>"},
		{"nil pointer", (*int)(nil), "%!(unsupported type)"},
		{"empty string", "", ""},
		{"string", "hello", "hello"},
		{"error type", errPing, "ping"},
		{"custom error", myErr("custom error msg"), "custom error msg"},
		{"bool true", true, "true"},
		{"bool false", false, "false"},
		{"int zero", 0, "0"},
		{"int positive", 67, "67"},
		{"int negative", -69, "-69"},
		{"int64 zero", int64(0), "0"},
		{"int64 positive", int64(42), "42"},
		{"int64 negative", int64(-42), "-42"},
		{"uint64 zero", uint64(0), "0"},
		{"uint64 big", uint64(1) << 63, "9223372036854775808"},
		{"uint64 max", uint64(^uint64(0)), "18446744073709551615"},
		{"bytes empty", []byte{}, ""},
		{"bytes non-empty", []byte("hello"), "hello"},
		{"bytes nil", []byte(nil), ""},
		{"unsupported float", 3.14, "%!(unsupported type)"},
		{"unsupported struct", struct{}{}, "%!(unsupported type)"},
		{"unsupported custom string (should fail)", Hoang("hello"), "%!(unsupported type)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sprint(tt.input); got != tt.want {
				t.Fatalf("Sprint(%#v) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestSprintWithList(t *testing.T) {
	l := mylist.New()
	if l == nil {
		t.Fatal("New() returned nil list")
	}

	l.PushBack(123)
	l.PushBack("hello")
	l.PushBack(true)
	l.PushBack(errors.New("pong"))

	var got []string
	for e := l.Front(); e != nil; e = e.Next() {
		got = append(got, Sprint(e.Value))
	}

	want := []string{"123", "hello", "true", "pong"}
	if len(got) != len(want) {
		t.Fatalf("len(got) = %d, want %d; got = %#v", len(got), len(want), got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got[%d] = %q, want %q; got = %#v", i, got[i], want[i], got)
		}
	}
}
