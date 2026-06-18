# go-peek

[![Go Reference](https://pkg.go.dev/badge/github.com/spenserblack/go-peek.svg)](https://pkg.go.dev/github.com/spenserblack/go-peek)
[![CI](https://github.com/spenserblack/go-peek/actions/workflows/ci.yml/badge.svg)](https://github.com/spenserblack/go-peek/actions/workflows/ci.yml)

## Description

This package helps convert [pull-style iterators][pull-iter] into a form that allows you to peek
one value ahead.

## Example

```go
tokens := []string{"Go", "Gopher", "Golang"}
seq := slices.Values(tokens)
next, stop := iter.Pull(seq)
defer stop()
pull, peek := peek.Peek(next)
// You usually should not call next after this point, and instead call either pull or peek.

// Get the next item, consuming it.
item, ok := pull()

// Get the next item, but do not consume it, so that it can be returned from a subsequent call to
// pull or peek.
peeked, ok := peek()
```

[pull-iter]: https://pkg.go.dev/iter#hdr-Pulling_Values
