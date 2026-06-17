package peek

import (
	"iter"
	"slices"
	"testing"
)

func TestPeek(t *testing.T) {
	seq := slices.Values([]string{"one", "two", "three"})
	next, stop := iter.Pull(seq)
	defer stop()

	pull, peek := Peek(next)

	tests := []struct {
		name string
		f    PullFunc[string]
		want string
		ok   bool
	}{
		{
			name: "pull first",
			f:    pull,
			want: "one",
			ok:   true,
		},
		{
			name: "peek second value",
			f:    peek,
			want: "two",
			ok:   true,
		},
		{
			name: "peek second value again",
			f:    peek,
			want: "two",
			ok:   true,
		},
		{
			name: "peek second value yet again",
			f:    peek,
			want: "two",
			ok:   true,
		},
		{
			name: "pull second value",
			f:    pull,
			want: "two",
			ok:   true,
		},
		{
			name: "pull third value",
			f:    pull,
			want: "three",
			ok:   true,
		},
		{
			name: "peek beyond end",
			f:    peek,
			ok:   false,
		},
		{
			name: "pull beyond end",
			f:    pull,
			ok:   false,
		},
	}

	for i, tt := range tests {
		t.Logf(`test %d: %s`, i, tt.name)
		got, ok := tt.f()
		if ok != tt.ok {
			t.Fatalf(`ok = %v, want %v`, ok, tt.ok)
		}
		if got != tt.want {
			t.Fatalf(`got = %v, want %v`, got, tt.want)
		}
	}
}

func TestPullPeek(t *testing.T) {
	seq := slices.Values([]string{"one", "two", "three"})
	pull, peek, stop := PullPeek(seq)
	defer stop()

	got, ok := pull()
	if !ok || got != "one" {
		t.Fatalf(`pull() = (%v, %v), want (%v, %v)`, got, ok, "one", true)
	}

	got, ok = peek()
	if !ok || got != "two" {
		t.Fatalf(`peek() = (%v, %v), want (%v, %v)`, got, ok, "two", true)
	}

	got, ok = peek()
	if !ok || got != "two" {
		t.Fatalf(`peek() again = (%v, %v), want (%v, %v)`, got, ok, "two", true)
	}

	got, ok = pull()
	if !ok || got != "two" {
		t.Fatalf(`pull() after peek = (%v, %v), want (%v, %v)`, got, ok, "two", true)
	}
}

func TestPeek2(t *testing.T) {
	seq := func(yield func(int, string) bool) {
		for _, item := range []struct {
			key   int
			value string
		}{
			{key: 1, value: "one"},
			{key: 2, value: "two"},
			{key: 3, value: "three"},
		} {
			if !yield(item.key, item.value) {
				return
			}
		}
	}
	next, stop := iter.Pull2(seq)
	defer stop()

	pull, peek := Peek2(next)

	tests := []struct {
		name      string
		f         PullFunc2[int, string]
		wantKey   int
		wantValue string
		ok        bool
	}{
		{
			name:      "pull first",
			f:         pull,
			wantKey:   1,
			wantValue: "one",
			ok:        true,
		},
		{
			name:      "peek second value",
			f:         peek,
			wantKey:   2,
			wantValue: "two",
			ok:        true,
		},
		{
			name:      "peek second value again",
			f:         peek,
			wantKey:   2,
			wantValue: "two",
			ok:        true,
		},
		{
			name:      "peek second value yet again",
			f:         peek,
			wantKey:   2,
			wantValue: "two",
			ok:        true,
		},
		{
			name:      "pull second value",
			f:         pull,
			wantKey:   2,
			wantValue: "two",
			ok:        true,
		},
		{
			name:      "pull third value",
			f:         pull,
			wantKey:   3,
			wantValue: "three",
			ok:        true,
		},
		{
			name: "peek beyond end",
			f:    peek,
			ok:   false,
		},
		{
			name: "pull beyond end",
			f:    pull,
			ok:   false,
		},
	}

	for i, tt := range tests {
		t.Logf(`test %d: %s`, i, tt.name)
		gotKey, gotValue, ok := tt.f()
		if ok != tt.ok {
			t.Fatalf(`ok = %v, want %v`, ok, tt.ok)
		}
		if gotKey != tt.wantKey || gotValue != tt.wantValue {
			t.Fatalf(`key, value = (%v, %v), want (%v, %v)`, gotKey, gotValue, tt.wantKey, tt.wantValue)
		}
	}
}

func TestPullPeek2(t *testing.T) {
	seq := func(yield func(int, string) bool) {
		for _, item := range []struct {
			key   int
			value string
		}{
			{key: 1, value: "one"},
			{key: 2, value: "two"},
			{key: 3, value: "three"},
		} {
			if !yield(item.key, item.value) {
				return
			}
		}
	}
	pull, peek, stop := PullPeek2(seq)
	defer stop()

	gotKey, gotValue, ok := pull()
	if !ok || gotKey != 1 || gotValue != "one" {
		t.Fatalf(`pull() = (%v, %v, %v), want (%v, %v, %v)`, gotKey, gotValue, ok, 1, "one", true)
	}

	gotKey, gotValue, ok = peek()
	if !ok || gotKey != 2 || gotValue != "two" {
		t.Fatalf(`peek() = (%v, %v, %v), want (%v, %v, %v)`, gotKey, gotValue, ok, 2, "two", true)
	}

	gotKey, gotValue, ok = peek()
	if !ok || gotKey != 2 || gotValue != "two" {
		t.Fatalf(`peek() again = (%v, %v, %v), want (%v, %v, %v)`, gotKey, gotValue, ok, 2, "two", true)
	}

	gotKey, gotValue, ok = pull()
	if !ok || gotKey != 2 || gotValue != "two" {
		t.Fatalf(`pull() after peek = (%v, %v, %v), want (%v, %v, %v)`, gotKey, gotValue, ok, 2, "two", true)
	}
}
