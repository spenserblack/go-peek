// Package peek provides utilities to peek one item ahead in a pull iterator.
package peek

// PullFunc is a function that returns the next value from a pull-style iterator. See the
// documentation for [iter.Pull] for more information.
type PullFunc[V any] = func() (V, bool)

// Peek takes next function (typically from [iter.Pull]) and returns a new next function
// and a peek function. You can call peek as many times as you want without consuming the iterator.
// Once next is called, the peeked item (if peek was called) is consumed.
func Peek[V any](next PullFunc[V]) (pull PullFunc[V], peek PullFunc[V]) {
	var (
		peekItem V
		peekOk   bool
	)

	pull = func() (V, bool) {
		if peekOk {
			peekOk = false
			return peekItem, true
		}
		return next()
	}
	peek = func() (V, bool) {
		if peekOk {
			return peekItem, true
		}
		peekItem, peekOk = next()
		return peekItem, peekOk
	}

	return pull, peek
}

// PullFunc2 is a function that returns the next value from a pull-style iterator. See the
// documentation for [iter.Pull2] for more information.
type PullFunc2[K any, V any] = func() (K, V, bool)

// Peek2 takes next function (typically from [iter.Pull]) and returns a new next function
// and a peek function. You can call peek as many times as you want without consuming the iterator.
// Once next is called, the peeked item (if peek was called) is consumed.
func Peek2[K any, V any](next PullFunc2[K, V]) (pull PullFunc2[K, V], peek PullFunc2[K, V]) {
	var (
		peekKey   K
		peekValue V
		peekOk    bool
	)

	pull = func() (K, V, bool) {
		if peekOk {
			peekOk = false
			return peekKey, peekValue, true
		}
		return next()
	}
	peek = func() (K, V, bool) {
		if peekOk {
			return peekKey, peekValue, true
		}
		peekKey, peekValue, peekOk = next()
		return peekKey, peekValue, peekOk
	}

	return pull, peek
}
