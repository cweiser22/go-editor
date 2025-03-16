package crdt

import (
	"cmp"
)

// LamportId represents a Lamport ID
type LamportId[T cmp.Ordered] struct {
	replicaId T
	count     int
}

// IsLessThan returns true if the ID is less than the other
func (id *LamportId[T]) IsLessThan(other *LamportId[T]) bool {
	if id.count != other.count {
		return id.count < other.count
	}
	return id.replicaId < other.replicaId
}
