package tracer

import (
	"math"
	"testing"
)

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
		t.Errorf("expected IsPoint() to be true for both bare struct and New method.")
	}
	notPoint := Vector(1, 1, 1)
	if notPoint.IsPoint() {
		t.Errorf("expected IsPoint() to be false for %v", notPoint)
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
	notVector := Point(1, 1, 1)
	if notVector.IsVector() {
		t.Errorf("expected IsVector() to be false for %v", notVector)
	}
}

func Test_Equals(t *testing.T) {
	same1 := NewTuple(1, 2, 3, 0)
	same2 := Vector(1, 2, 3)
	diff := Point(1, 2, 3)
	zero := NewTuple(0, 0, 0, 0)
	if Point(1, 2, 3).Equals(Point(2, 2, 3)) {
		t.Errorf("Tuples should be unequal because X")
	}
	if Point(1, 2, 3).Equals(Point(1, 1, 3)) {
		t.Errorf("Tuples should be unequal because Y")
	}
	if Point(1, 2, 3).Equals(Point(1, 2, 2)) {
		t.Errorf("Tuples should be unequal because Z")
	}
	if Point(1, 2, 3).Equals(Vector(1, 2, 3)) {
		t.Errorf("Tuples should be unequal because  W")
	}
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

func Test_Minus_Method(t *testing.T) {
	// Subtracting two points
	p1 := Point(3, 2, 1)
	p2 := Point(5, 6, 7)
	result2p := p1.Minus(p2)
	expect2p := Vector(-2, -4, -6)
	if !result2p.Equals(expect2p) {
		t.Errorf("%v Minus %v should be %v, but was %v", p1, p2, expect2p, result2p)
	}
	// Subtracting a vector from a point
	v2 := Vector(5, 6, 7)
	resultvp := p1.Minus(v2)
	expectvp := Point(-2, -4, -6)
	if !resultvp.Equals(expectvp) {
		t.Errorf("%v Minus %v should be %v, but was %v", p1, v2, expectvp, resultvp)
	}
	// Subtracting a vector from a vector
	v1 := Vector(3, 2, 1)
	result2v := v1.Minus(v2)
	expect2v := Vector(-2, -4, -6)
	if !result2v.Equals(expect2v) {
		t.Errorf("%v Minus %v should be %v, but was %v", v1, v2, expect2v, result2v)
	}
}

func Test_Minus_Func(t *testing.T) {
	t1 := NewTuple(1, -2, 3, -4)
	result := Minus(t1)
	expect := NewTuple(-1, 2, -3, 4)
	if !result.Equals(expect) {
		t.Errorf("Minus %v should be %v, but was %v", t1, expect, result)
	}
}

func Test_Times(t *testing.T) {
	t1 := NewTuple(1, -2, 3, -4)
	f1 := 3.5
	result1 := t1.Times(f1)
	expect1 := NewTuple(3.5, -7, 10.5, -14)
	if !result1.Equals(expect1) {
		t.Errorf("%v Times %v should be %v, but was %v", t1, f1, expect1, result1)
	}
	f2 := 0.5
	result2 := t1.Times(f2)
	expect2 := NewTuple(0.5, -1, 1.5, -2)
	if !result2.Equals(expect2) {
		t.Errorf("%v Times %v should be %v, but was %v", t1, f2, expect2, result2)
	}
}

func Test_Divide(t *testing.T) {
	t1 := NewTuple(1, -2, 3, -4)
	div := 2.0
	result := t1.Divide(div)
	expect := NewTuple(0.5, -1, 1.5, -2)
	if !result.Equals(expect) {
		t.Errorf("%v Divide %v should be %v, but was %v", t1, div, expect, result)
	}
}

func Test_Divide_ByZero(t *testing.T) {
	defer func() { _ = recover() }()
	t1 := NewTuple(1, -2, 3, -4)
	_ = t1.Divide(0.0)
	t.Errorf("Dividing %v by zero scalar did not panic", t1)
}

func Test_Length(t *testing.T) {
	tuples := []Tuple{
		Vector(1, 0, 0),
		Vector(0, 1, 0),
		Vector(0, 0, 1),
		Vector(1, 2, 3),
		Vector(-1, -2, -3),
	}

	lengths := []float64{
		1.0, 1.0, 1.0, math.Sqrt(14.0), math.Sqrt(14.0),
	}

	for ix, tup := range tuples {
		if tup.Length() != lengths[ix] {
			t.Errorf("Length of %v should be %v, but was %v", tup, lengths[ix], tup.Length())
		}
	}

}
