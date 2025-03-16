package crdt

import "golang.org/x/exp/constraints"

// LamportID represents a Lamport ID
type LamportID[T constraints.Ordered] struct {
	replicaId T
	count     int
}

// IsLessThan returns true if the ID is less than the other
func (id *LamportID[T]) IsLessThan(other *LamportID[T]) bool {
	if id.count != other.count {
		return id.count < other.count
	}
	return id.replicaId < other.replicaId
}
