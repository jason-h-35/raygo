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

func Test_Equals(t *testing.T) {
	same1 := NewTuple(1, 2, 3, 0)
	same2 := Vector(1, 2, 3)
	diff := Point(1, 2, 3)
	zero := NewTuple(0, 0, 0, 0)
	if !same1.Equals(same2) || !same2.Equals(same1) {
		t.Errorf("expected %v Equals %v", same1, same2)
	}
	if same2.Equals(diff) {
		t.Errorf("expected %v Not Equals %v", same2, diff)
	}
	if zero.Equals(same1) || zero.Equals(same2) || zero.Equals(diff) {
		t.Errorf("zero tuple shouldn't equal %v, %v, or %v", same1, same2, diff)
	}
}

func Test_Plus(t *testing.T) {
	point := Point(1, 2, 3)
	vector := Vector(1, 2, 3)
	test1 := point.Plus(vector)
	test2 := vector.Plus(point)
	test3 := vector.Plus(vector)
	if !test1.Equals(test2) {
		t.Errorf("expected Plus to be commutative: %v would Equal %v", test1, test2)
	}
	if !test1.IsPoint() || !test2.IsPoint() {
		t.Errorf("expected a Vector Plus a Point would be a Point: %v", test1)
	}
	if !test3.IsVector() {
		t.Errorf("expected a Vector Plus a Vector would be a Vector: %v", test3)
	}
}
