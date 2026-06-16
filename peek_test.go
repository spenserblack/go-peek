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
