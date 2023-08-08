package tracer

import "testing"

func Test_NewTuple(t *testing.T) {
	tuple := NewTuple(4.3, -4.2, 3.1, 1.0)
	if tuple.X != 4.3 || tuple.Y != -4.2 || tuple.Z != 3.1 || tuple.W != 1.0 {
		t.Error("tuple fields incorrect")
	}
}
