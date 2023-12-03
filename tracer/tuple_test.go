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
	got := NewPointTuple(4.3, -4.2, 3.1)
	if got != expect {
		t.Errorf("expected: %v. got: %v", expect, got)
	}
	if !got.IsPoint() || !expect.IsPoint() {
		t.Errorf("expected IsPoint() to be true for both bare struct and New method.")
	}
	notPoint := NewVectorTuple(1, 1, 1)
	if notPoint.IsPoint() {
		t.Errorf("expected IsPoint() to be false for %v", notPoint)
	}
}

func Test_Vector(t *testing.T) {
	expect := Tuple{X: 4.3, Y: -4.2, Z: 3.1, W: 0}
	got := NewVectorTuple(4.3, -4.2, 3.1)
	if got != expect {
		t.Errorf("expected: %v. got: %v", expect, got)
	}
	if !got.IsVector() || !expect.IsVector() {
		t.Errorf("expected IsVector() == true for both bare struct and New method.")
	}
	notVector := NewPointTuple(1, 1, 1)
	if notVector.IsVector() {
		t.Errorf("expected IsVector() to be false for %v", notVector)
	}
}

func Test_Equals(t *testing.T) {
	same1 := NewTuple(1, 2, 3, 0)
	same2 := NewVectorTuple(1, 2, 3)
	diff := NewPointTuple(1, 2, 3)
	zero := NewTuple(0, 0, 0, 0)
	if NewPointTuple(1, 2, 3).Equals(NewPointTuple(2, 2, 3)) {
		t.Errorf("Tuples should be unequal because X")
	}
	if NewPointTuple(1, 2, 3).Equals(NewPointTuple(1, 1, 3)) {
		t.Errorf("Tuples should be unequal because Y")
	}
	if NewPointTuple(1, 2, 3).Equals(NewPointTuple(1, 2, 2)) {
		t.Errorf("Tuples should be unequal because Z")
	}
	if NewPointTuple(1, 2, 3).Equals(NewVectorTuple(1, 2, 3)) {
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
	point := NewPointTuple(1, 2, 3)
	vector := NewVectorTuple(1, 2, 3)
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
	p1 := NewPointTuple(3, 2, 1)
	p2 := NewPointTuple(5, 6, 7)
	result2p := p1.Minus(p2)
	expect2p := NewVectorTuple(-2, -4, -6)
	if !result2p.Equals(expect2p) {
		t.Errorf("%v Minus %v should be %v, but was %v", p1, p2, expect2p, result2p)
	}
	// Subtracting a vector from a point
	v2 := NewVectorTuple(5, 6, 7)
	resultvp := p1.Minus(v2)
	expectvp := NewPointTuple(-2, -4, -6)
	if !resultvp.Equals(expectvp) {
		t.Errorf("%v Minus %v should be %v, but was %v", p1, v2, expectvp, resultvp)
	}
	// Subtracting a vector from a vector
	v1 := NewVectorTuple(3, 2, 1)
	result2v := v1.Minus(v2)
	expect2v := NewVectorTuple(-2, -4, -6)
	if !result2v.Equals(expect2v) {
		t.Errorf("%v Minus %v should be %v, but was %v", v1, v2, expect2v, result2v)
	}
}

func Test_Minus_Func(t *testing.T) {
	t1 := NewTuple(1, -2, 3, -4)
	result := NewTuple(0, 0, 0, 0).Minus(t1)
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

func Test_Divide_Panic(t *testing.T) {
	defer func() { _ = recover() }()
	t1 := NewTuple(1, -2, 3, -4)
	_ = t1.Divide(0.0)
	t.Errorf("Dividing %v by zero scalar did not panic", t1)
}

func Test_Length(t *testing.T) {
	tuples := []Tuple{
		NewVectorTuple(1, 0, 0),
		NewVectorTuple(0, 1, 0),
		NewVectorTuple(0, 0, 1),
		NewVectorTuple(1, 2, 3),
		NewVectorTuple(-1, -2, -3),
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

func Test_Normalized(t *testing.T) {
	tuples := []Tuple{
		NewVectorTuple(4, 0, 0),
		NewVectorTuple(1, 2, 3),
	}
	normalized := []Tuple{
		NewVectorTuple(1, 0, 0),
		NewVectorTuple(1, 2, 3).Divide(math.Sqrt(14)),
	}
	for ix, tup := range tuples {
		if tup.Normalized() != normalized[ix] {
			t.Errorf("%v Normalized should be %v, but was %v", tup, normalized[ix], tup.Normalized())
		}
		if tup.Normalized().Length()-1 > eps {
			t.Errorf("%v Normalized should have Length of 1, but was %v", tup, tup.Normalized().Length())
		}
	}
}

func Test_Normalized_Panic(t *testing.T) {
	defer func() { _ = recover() }()
	t1 := NewTuple(0, 0, 0, 0)
	_ = t1.Normalized()
	t.Errorf("Normalizing %v did not panic", t1)
}

func Test_Dot(t *testing.T) {
	t1 := NewVectorTuple(1, 2, 3)
	t2 := NewVectorTuple(2, 3, 4)
	result := t1.Dot(t2)
	expect := 20.0
	if result != expect {
		t.Errorf("%v Dot %v should be %v, but was %v", t1, t2, expect, result)
	}
	if t1.Dot(t2) != t2.Dot(t1) {
		t.Errorf("%v Dot %v and %v Dot %v should be equal but are instead %v and %v",
			t1, t2, t2, t1, t1.Dot(t2), t2.Dot(t1))
	}
}

func Test_Dot_Panic(t *testing.T) {
	defer func() { _ = recover() }()
	p := NewPointTuple(1, 2, 3)
	v := NewVectorTuple(4, 5, 6)
	_ = p.Dot(v)
	t.Errorf("%v Dot %v did not panic", p, v)
}

func Test_Cross(t *testing.T) {
	v1 := NewVectorTuple(1, 2, 3)
	v2 := NewVectorTuple(2, 3, 4)
	expect1 := NewVectorTuple(-1, 2, -1)
	if v1.Cross(v2) != expect1 {
		t.Errorf("%v Cross %v should be %v, but was %v", v1, v2, expect1, v1.Cross(v2))
	}
	expect2 := NewVectorTuple(1, -2, 1)
	if v2.Cross(v1) != expect2 {
		t.Errorf("%v Cross %v should be %v, but was %v", v2, v1, expect2, v2.Cross(v1))
	}
}

func Test_Cross_Panic(t *testing.T) {
	defer func() { _ = recover() }()
	p := NewPointTuple(1, 2, 3)
	v := NewVectorTuple(4, 5, 6)
	_ = p.Cross(v)
	t.Errorf("%v Cross %v did not panic", p, v)
}
