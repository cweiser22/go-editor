package crdt

import (
	"cmp"
)

// sentinel value to represent a blank operation
var NilID = LamportId[]{}

type RGA[T cmp.Ordered] struct {
	clock      *LamportClock[T]
	operations map[LamportId[T]]RGAOperation[T]
}

type RGAOperation[T cmp.Ordered] struct {
	id        LamportId[T]
	deleted   bool
	character rune
	prevId    LamportId[T]
	nextId    LamportId[T]
}

func NewRGA[T cmp.Ordered](replicaId T) *RGA[T] {
	return &RGA[T]{
		clock:      NewLamportClock(replicaId),
		operations: make(map[LamportId[T]]RGAOperation[T]),
	}
}

func (r *RGA[T]) Insert(prevId LamportId[T], c rune) {
	opId := r.clock.Tick()

	var nextId LamportId[T] = NilID
	if prevId != nil {
		prev := r.operations[prevId]
		nextId := prev.nextId
	}

	newOp := RGAOperation[T]{
		id:        opId,
		deleted:   false,
		character: c,
		prevId:    prevId,
		nextId:    nextId,
	}
	prev.nextId = newOp.id
}

func (r *RGA[T]) Delete(id LamportId[T]) {

}

func (r *RGA[T]) Merge(rga RGA[T]) {}
