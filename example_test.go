package peek_test

import (
	"fmt"
	"iter"
	"slices"

	"github.com/spenserblack/go-peek"
)

// This package provides utilities for peeking ahead one item from a pull-style iterator.
func Example() {
	values := slices.Values([]string{"one", "two", "three"})
	next, stop := iter.Pull(values)
	pull, peek := peek.Peek(next)
	defer stop()

	// You usually should not call next after this point, and call either pull or peek instead.

	printResult := func(item string, ok bool) {
		if ok {
			fmt.Println(item)
		} else {
			fmt.Println("no value")
		}
	}

	// pull consumes the item
	printResult(pull())

	// peeked items are not consumed until they are pulled
	printResult(peek())
	printResult(peek())
	printResult(pull())

	printResult(peek())
	printResult(pull())

	// ok is false once the iterator is consumed
	printResult(peek())
	printResult(pull())
	// Output:
	// one
	// two
	// two
	// two
	// three
	// three
	// no value
	// no value
}
