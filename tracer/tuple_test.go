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
