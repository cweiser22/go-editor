package crdt

import "golang.org/x/exp/constraints"

type RGA[T constraints.Ordered] struct {
	clock      *LamportClock[T]
	operations []LamportID[T]
}
