package tracer

import "testing"

func Test_NewTuple(t *testing.T) {
	expect := Tuple{X: 4.3, Y: -4.2, Z: 3.1, W: 1}
	got := NewTuple(4.3, -4.2, 3.1, 1)
	if got != expect {
		t.Errorf("expected: %v. got: %v", expect, got)
	}
}

func Test_Point(t *testing.T) {
	expect := Tuple{X: 4.3, Y: -4.2, Z: 3.1, W: 1}
	got := Point(4.3, -4.2, 3.1)
	if got != expect {
		t.Errorf("expected: %v. got: %v", expect, got)
	}
	if !got.IsPoint() || !expect.IsPoint() {
		t.Errorf("expected IsPoint() == true for both bare struct and New method.")
	}
}

func Test_Vector(t *testing.T) {
	expect := Tuple{X: 4.3, Y: -4.2, Z: 3.1, W: 0}
	got := Vector(4.3, -4.2, 3.1)
	if got != expect {
		t.Errorf("expected: %v. got: %v", expect, got)
	}
	if !got.IsVector() || !expect.IsVector() {
		t.Errorf("expected IsVector() == true for both bare struct and New method.")
	}
}
