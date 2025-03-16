package crdt

import "testing"

func TestIsLessThan(t *testing.T) {
	// first test two in the same counter
	replicaId := 1
	lc := NewLamportClock(replicaId)

	id1, id2 := lc.Tick(), lc.Tick()

	if !id1.IsLessThan(id2) {
		t.Errorf("id1 should have been less than id2 but wasn't")
	}

	lc1 := NewLamportClock(1)
	lc2 := NewLamportClock(2)

	// creates two pairs in alternating node order
	id1, id2 = lc1.Tick(), lc2.Tick()
	id3, id4 := lc2.Tick(), lc1.Tick()

	if !id1.IsLessThan(id2) {
		t.Errorf("id1 should be less than id2")
	}
	if !id4.IsLessThan(id3) {
		t.Errorf("id4 should be less than id3")
	}

}
