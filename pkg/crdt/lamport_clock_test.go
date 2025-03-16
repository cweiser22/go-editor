package crdt

import "testing"

func TestNewLamportClock(t *testing.T) {
	replicaId := "1234"
	lc := NewLamportClock(replicaId)

	if _, ok := any(lc.replicaID).(string); !ok {
		t.Errorf("Replica ID not of string type.")
	}
}

func TestTick(t *testing.T) {
	replicaId := "1234"
	lc := NewLamportClock(replicaId)

	_ = lc.Tick()

	if lc.counter != 1 {
		t.Errorf("counter was %d when it should have been 1", lc.counter)
	}
}
