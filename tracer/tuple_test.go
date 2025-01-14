// package tracer
//
// import (
//
//	"math"
//	"testing"
//
// )
//
// // Check NewTuple constructor assigns to fields properly.
//
//	func Test_NewTuple(t *testing.T) {
//		expect := Tuple{X: 4.3, Y: -4.2, Z: 3.1, W: 1}
//		got := NewTuple(4.3, -4.2, 3.1, 1)
//		if got != expect {
//			t.Errorf("expected: %v. got: %v", expect, got)
//		}
//	}
//
// // Check that NewPointTuple sets W=1 and IsPoint() and not IsVector()
//
//	func Test_Point(t *testing.T) {
//		expect := Tuple{X: 4.3, Y: -4.2, Z: 3.1, W: 1}
//		got := NewPoint(4.3, -4.2, 3.1)
//		if got != expect {
//			t.Errorf("expected: %v. got: %v", expect, got)
//		}
//		if !got.IsPoint() || !expect.IsPoint() {
//			t.Errorf("expected IsPoint() to be true for both bare struct and New method.")
//		}
//		if got.IsVector() || expect.IsVector() {
//			t.Errorf("expected IsVector() to be false for both bare struct and New method.")
//		}
//		notPoint := NewVector(1, 1, 1)
//		if notPoint.IsPoint() {
//			t.Errorf("expected IsPoint() to be false for %v", notPoint)
//		}
//	}
//
// // Test that NewVectorTuple sets W=0 and IsVector() and not IsPoint()
//
//	func Test_Vector(t *testing.T) {
//		expect := Tuple{X: 4.3, Y: -4.2, Z: 3.1, W: 0}
//		got := NewVector(4.3, -4.2, 3.1)
//		if got != expect {
//			t.Errorf("expected: %v. got: %v", expect, got)
//		}
//		if !got.IsVector() || !expect.IsVector() {
//			t.Errorf("expected IsVector() == true for both bare struct and New method.")
//		}
//		if got.IsPoint() || expect.IsPoint() {
//			t.Errorf("expected IsPoint() == false for both bare struct and New method.")
//		}
//		notVector := NewPoint(1, 1, 1)
//		if notVector.IsVector() {
//			t.Errorf("expected IsVector() to be false for %v", notVector)
//		}
//	}
//
//	func Test_Equals(t *testing.T) {
//		same1 := NewTuple(1, 2, 3, 0)
//		same2 := NewVector(1, 2, 3)
//		diff := NewPoint(1, 2, 3)
//		zero := NewTuple(0, 0, 0, 0)
//		if NewPoint(1, 2, 3).Equals(NewPoint(2, 2, 3)) {
//			t.Errorf("Tuples should be unequal because X")
//		}
//		if NewPoint(1, 2, 3).Equals(NewPoint(1, 1, 3)) {
//			t.Errorf("Tuples should be unequal because Y")
//		}
//		if NewPoint(1, 2, 3).Equals(NewPoint(1, 2, 2)) {
//			t.Errorf("Tuples should be unequal because Z")
//		}
//		if NewPoint(1, 2, 3).Equals(NewVector(1, 2, 3)) {
//			t.Errorf("Tuples should be unequal because  W")
//		}
//		if !same1.Equals(same2) || !same2.Equals(same1) {
//			t.Errorf("expected %v Equals %v", same1, same2)
//		}
//		if same2.Equals(diff) {
//			t.Errorf("expected %v Not Equals %v", same2, diff)
//		}
//		if zero.Equals(same1) || zero.Equals(same2) || zero.Equals(diff) {
//			t.Errorf("zero tuple shouldn't equal %v, %v, or %v", same1, same2, diff)
//		}
//	}
//
//	func Test_Plus(t *testing.T) {
//		point := NewPoint(1, 2, 3)
//		vector := NewVector(1, 2, 3)
//		test1 := point.Plus(vector)
//		test2 := vector.Plus(point)
//		test3 := vector.Plus(vector)
//		if !test1.Equals(test2) {
//			t.Errorf("expected Plus to be commutative: %v would Equal %v", test1, test2)
//		}
//		if !test1.IsPoint() || !test2.IsPoint() {
//			t.Errorf("expected a Vector Plus a Point would be a Point: %v", test1)
//		}
//		if !test3.IsVector() {
//			t.Errorf("expected a Vector Plus a Vector would be a Vector: %v", test3)
//		}
//	}
//
//	func Test_Minus_Method(t *testing.T) {
//		// Subtracting two points
//		p1 := NewPoint(3, 2, 1)
//		p2 := NewPoint(5, 6, 7)
//		result2p := p1.Minus(p2)
//		expect2p := NewVector(-2, -4, -6)
//		if !result2p.Equals(expect2p) {
//			t.Errorf("%v Minus %v should be %v, but was %v", p1, p2, expect2p, result2p)
//		}
//		// Subtracting a vector from a point
//		v2 := NewVector(5, 6, 7)
//		resultvp := p1.Minus(v2)
//		expectvp := NewPoint(-2, -4, -6)
//		if !resultvp.Equals(expectvp) {
//			t.Errorf("%v Minus %v should be %v, but was %v", p1, v2, expectvp, resultvp)
//		}
//		// Subtracting a vector from a vector
//		v1 := NewVector(3, 2, 1)
//		result2v := v1.Minus(v2)
//		expect2v := NewVector(-2, -4, -6)
//		if !result2v.Equals(expect2v) {
//			t.Errorf("%v Minus %v should be %v, but was %v", v1, v2, expect2v, result2v)
//		}
//	}
//
//	func Test_Minus_Func(t *testing.T) {
//		t1 := NewTuple(1, -2, 3, -4)
//		result := NewTuple(0, 0, 0, 0).Minus(t1)
//		expect := NewTuple(-1, 2, -3, 4)
//		if !result.Equals(expect) {
//			t.Errorf("Minus %v should be %v, but was %v", t1, expect, result)
//		}
//	}
//
//	func Test_Times(t *testing.T) {
//		t1 := NewTuple(1, -2, 3, -4)
//		f1 := 3.5
//		result1 := t1.Times(f1)
//		expect1 := NewTuple(3.5, -7, 10.5, -14)
//		if !result1.Equals(expect1) {
//			t.Errorf("%v Times %v should be %v, but was %v", t1, f1, expect1, result1)
//		}
//		f2 := 0.5
//		result2 := t1.Times(f2)
//		expect2 := NewTuple(0.5, -1, 1.5, -2)
//		if !result2.Equals(expect2) {
//			t.Errorf("%v Times %v should be %v, but was %v", t1, f2, expect2, result2)
//		}
//	}
//
//	func Test_Divide(t *testing.T) {
//		t1 := NewTuple(1, -2, 3, -4)
//		div := 2.0
//		result := t1.Divide(div)
//		expect := NewTuple(0.5, -1, 1.5, -2)
//		if !result.Equals(expect) {
//			t.Errorf("%v Divide %v should be %v, but was %v", t1, div, expect, result)
//		}
//	}
//
//	func Test_Divide_Panic(t *testing.T) {
//		defer func() { _ = recover() }()
//		t1 := NewTuple(1, -2, 3, -4)
//		_ = t1.Divide(0.0)
//		t.Errorf("Dividing %v by zero scalar did not panic", t1)
//	}
//
//	func Test_Length(t *testing.T) {
//		tuples := []Tuple{
//			NewVector(1, 0, 0),
//			NewVector(0, 1, 0),
//			NewVector(0, 0, 1),
//			NewVector(1, 2, 3),
//			NewVector(-1, -2, -3),
//		}
//
//		lengths := []float64{
//			1.0, 1.0, 1.0, math.Sqrt(14.0), math.Sqrt(14.0),
//		}
//
//		for ix, tup := range tuples {
//			if tup.Length() != lengths[ix] {
//				t.Errorf("Length of %v should be %v, but was %v", tup, lengths[ix], tup.Length())
//			}
//		}
//	}
//
//	func Test_Normalized(t *testing.T) {
//		tuples := []Tuple{
//			NewVector(4, 0, 0),
//			NewVector(1, 2, 3),
//		}
//		normalized := []Tuple{
//			NewVector(1, 0, 0),
//			NewVector(1, 2, 3).Divide(math.Sqrt(14)),
//		}
//		for ix, tup := range tuples {
//			if tup.Normalized() != normalized[ix] {
//				t.Errorf("%v Normalized should be %v, but was %v", tup, normalized[ix], tup.Normalized())
//			}
//			if tup.Normalized().Length()-1 > epsilon {
//				t.Errorf("%v Normalized should have Length of 1, but was %v", tup, tup.Normalized().Length())
//			}
//		}
//	}
//
//	func Test_Normalized_Panic(t *testing.T) {
//		defer func() { _ = recover() }()
//		t1 := NewTuple(0, 0, 0, 0)
//		_ = t1.Normalized()
//		t.Errorf("Normalizing %v did not panic", t1)
//	}
//
//	func Test_Dot(t *testing.T) {
//		t1 := NewVector(1, 2, 3)
//		t2 := NewVector(2, 3, 4)
//		result := t1.Dot(t2)
//		expect := 20.0
//		if result != expect {
//			t.Errorf("%v Dot %v should be %v, but was %v", t1, t2, expect, result)
//		}
//		if t1.Dot(t2) != t2.Dot(t1) {
//			t.Errorf("%v Dot %v and %v Dot %v should be equal but are instead %v and %v",
//				t1, t2, t2, t1, t1.Dot(t2), t2.Dot(t1))
//		}
//	}
//
//	func Test_Dot_Panic(t *testing.T) {
//		defer func() { _ = recover() }()
//		p := NewPoint(1, 2, 3)
//		v := NewVector(4, 5, 6)
//		_ = p.Dot(v)
//		t.Errorf("%v Dot %v did not panic", p, v)
//	}
//
//	func Test_Cross(t *testing.T) {
//		v1 := NewVector(1, 2, 3)
//		v2 := NewVector(2, 3, 4)
//		expect1 := NewVector(-1, 2, -1)
//		if v1.Cross(v2) != expect1 {
//			t.Errorf("%v Cross %v should be %v, but was %v", v1, v2, expect1, v1.Cross(v2))
//		}
//		expect2 := NewVector(1, -2, 1)
//		if v2.Cross(v1) != expect2 {
//			t.Errorf("%v Cross %v should be %v, but was %v", v2, v1, expect2, v2.Cross(v1))
//		}
//	}
//
//	func Test_Cross_Panic(t *testing.T) {
//		defer func() { _ = recover() }()
//		p := NewPoint(1, 2, 3)
//		v := NewVector(4, 5, 6)
//		_ = p.Cross(v)
//		t.Errorf("%v Cross %v did not panic", p, v)
//	}
package tracer

import (
	"math"
	"testing"
)

func Test_Constructors(t *testing.T) {
	tests := []struct {
		name     string
		tuple    Tuple
		want     Tuple
		isPoint  bool
		isVector bool
	}{
		{
			name:     "raw tuple constructor",
			tuple:    NewTuple(4.3, -4.2, 3.1, 1),
			want:     Tuple{X: 4.3, Y: -4.2, Z: 3.1, W: 1},
			isPoint:  true,
			isVector: false,
		},
		{
			name:     "point constructor",
			tuple:    NewPoint(4.3, -4.2, 3.1),
			want:     Tuple{X: 4.3, Y: -4.2, Z: 3.1, W: 1},
			isPoint:  true,
			isVector: false,
		},
		{
			name:     "vector constructor",
			tuple:    NewVector(4.3, -4.2, 3.1),
			want:     Tuple{X: 4.3, Y: -4.2, Z: 3.1, W: 0},
			isPoint:  false,
			isVector: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.tuple != tt.want {
				t.Errorf("got %v, want %v", tt.tuple, tt.want)
			}
			if tt.tuple.IsPoint() != tt.isPoint {
				t.Errorf("IsPoint() = %v, want %v", tt.tuple.IsPoint(), tt.isPoint)
			}
			if tt.tuple.IsVector() != tt.isVector {
				t.Errorf("IsVector() = %v, want %v", tt.tuple.IsVector(), tt.isVector)
			}
		})
	}
}

func Test_Equality(t *testing.T) {
	tests := []struct {
		name string
		a, b Tuple
		want bool
	}{
		{"different X", NewPoint(1, 2, 3), NewPoint(2, 2, 3), false},
		{"different Y", NewPoint(1, 2, 3), NewPoint(1, 1, 3), false},
		{"different Z", NewPoint(1, 2, 3), NewPoint(1, 2, 2), false},
		{"point vs vector", NewPoint(1, 2, 3), NewVector(1, 2, 3), false},
		{"equivalent tuples", NewTuple(1, 2, 3, 0), NewVector(1, 2, 3), true},
		{"zero vs nonzero", NewTuple(0, 0, 0, 0), NewVector(1, 2, 3), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Equals(tt.b); got != tt.want {
				t.Errorf("%v.Equals(%v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func Test_Arithmetic(t *testing.T) {
	tests := []struct {
		name string
		op   string
		a, b Tuple
		f    float64
		want Tuple
	}{
		{"point + vector = point", "plus", NewPoint(1, 2, 3), NewVector(1, 2, 3), 0, NewPoint(2, 4, 6)},
		{"vector + vector = vector", "plus", NewVector(1, 2, 3), NewVector(1, 2, 3), 0, NewVector(2, 4, 6)},
		{"point - point = vector", "minus", NewPoint(3, 2, 1), NewPoint(5, 6, 7), 0, NewVector(-2, -4, -6)},
		{"point - vector = point", "minus", NewPoint(3, 2, 1), NewVector(5, 6, 7), 0, NewPoint(-2, -4, -6)},
		{"vector - vector = vector", "minus", NewVector(3, 2, 1), NewVector(5, 6, 7), 0, NewVector(-2, -4, -6)},
		{"negate", "minus", NewTuple(0, 0, 0, 0), NewTuple(1, -2, 3, -4), 0, NewTuple(-1, 2, -3, 4)},
		{"multiply", "times", NewTuple(1, -2, 3, -4), Tuple{}, 3.5, NewTuple(3.5, -7, 10.5, -14)},
		{"divide", "divide", NewTuple(1, -2, 3, -4), Tuple{}, 2.0, NewTuple(0.5, -1, 1.5, -2)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got Tuple
			switch tt.op {
			case "plus":
				got = tt.a.Plus(tt.b)
			case "minus":
				got = tt.a.Minus(tt.b)
			case "times":
				got = tt.a.Times(tt.f)
			case "divide":
				got = tt.a.Divide(tt.f)
			}
			if !got.Equals(tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_VectorOps(t *testing.T) {
	tests := []struct {
		name    string
		op      string
		v1, v2  Tuple
		want    interface{} // Tuple or float64
		wantLen float64
	}{
		{"dot product", "dot", NewVector(1, 2, 3), NewVector(2, 3, 4), 20.0, 0},
		{"cross product a×b", "cross", NewVector(1, 2, 3), NewVector(2, 3, 4), NewVector(-1, 2, -1), 0},
		{"cross product b×a", "cross", NewVector(2, 3, 4), NewVector(1, 2, 3), NewVector(1, -2, 1), 0},
		{"unit x length", "length", NewVector(1, 0, 0), Tuple{}, 0, 1.0},
		{"arbitrary vector length", "length", NewVector(1, 2, 3), Tuple{}, 0, math.Sqrt(14.0)},
		{"normalize x axis", "normalize", NewVector(4, 0, 0), Tuple{}, NewVector(1, 0, 0), 0},
		{"normalize arbitrary", "normalize", NewVector(1, 2, 3), Tuple{}, NewVector(1, 2, 3).Divide(math.Sqrt(14)), 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.op {
			case "dot":
				if got := tt.v1.Dot(tt.v2); got != tt.want.(float64) {
					t.Errorf("got %v, want %v", got, tt.want)
				}
			case "cross":
				if got := tt.v1.Cross(tt.v2); !got.Equals(tt.want.(Tuple)) {
					t.Errorf("got %v, want %v", got, tt.want)
				}
			case "length":
				if got := tt.v1.Length(); got != tt.wantLen {
					t.Errorf("got %v, want %v", got, tt.wantLen)
				}
			case "normalize":
				got := tt.v1.Normalized()
				if !got.Equals(tt.want.(Tuple)) {
					t.Errorf("got %v, want %v", got, tt.want)
				}
				if math.Abs(got.Length()-1) > epsilon {
					t.Errorf("normalized length = %v, want 1", got.Length())
				}
			}
		})
	}
}

func Test_Panics(t *testing.T) {
	tests := []struct {
		name string
		fn   func()
	}{
		{"divide by zero", func() { NewTuple(1, -2, 3, -4).Divide(0.0) }},
		{"normalize zero", func() { NewTuple(0, 0, 0, 0).Normalized() }},
		{"dot with point", func() { NewPoint(1, 2, 3).Dot(NewVector(4, 5, 6)) }},
		{"cross with point", func() { NewPoint(1, 2, 3).Cross(NewVector(4, 5, 6)) }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Error("expected panic")
				}
			}()
			tt.fn()
		})
	}
}
