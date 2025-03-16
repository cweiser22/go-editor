package crdt

import (
	"cmp"
)

// LamportClock encapsulates the counter logic for a node that creates Lamport IDs
type LamportClock[T cmp.Ordered] struct {
	replicaID T
	counter   int
}

func NewLamportClock[T cmp.Ordered](replicaID T) *LamportClock[T] {
	return &LamportClock[T]{
		replicaID: replicaID,
		counter:   0,
	}
}

// Tick creates a new Lamport ID and increments the counter
func (lc *LamportClock[T]) Tick() LamportId[T] {
	newId := LamportId[T]{
		count:     lc.counter,
		replicaId: lc.replicaID,
	}
	lc.counter++
	return newId
}

// UpdateCounter sets the counter's to the largest of v + 1 and the current counter
func (lc *LamportClock[T]) UpdateCounter(v int) {
	lc.counter = max(lc.counter, v+1)
}
