package tracer

import (
	"math"
)

const epsilon = 1e-5

type Tuple struct {
	X, Y, Z, W float64
}

func NewTuple(x, y, z, w float64) Tuple {
	return Tuple{x, y, z, w}
}

func (t Tuple) AsArray() [4]float64 {
	return [4]float64{t.X, t.Y, t.Z, t.W}
}

func NewPoint(x, y, z float64) Tuple {
	return NewTuple(x, y, z, 1)
}

func NewVector(x, y, z float64) Tuple {
	return NewTuple(x, y, z, 0)
}

func (t Tuple) IsVector() bool {
	return t.W == 0
}

func (t Tuple) IsPoint() bool {
	return t.W == 1
}

func (t1 Tuple) Equals(t2 Tuple) bool {
	return math.Abs(t1.X-t2.X) <= epsilon &&
		math.Abs(t1.Y-t2.Y) <= epsilon &&
		math.Abs(t1.Z-t2.Z) <= epsilon &&
		t1.W == t2.W
}

func (t1 Tuple) Plus(t2 Tuple) Tuple {
	return NewTuple(
		t1.X+t2.X, t1.Y+t2.Y,
		t1.Z+t2.Z, t1.W+t2.W,
	)
}

func (t1 Tuple) Minus(t2 Tuple) Tuple {
	return NewTuple(
		t1.X-t2.X, t1.Y-t2.Y,
		t1.Z-t2.Z, t1.W-t2.W,
	)
}

func (t Tuple) Times(f float64) Tuple {
	return NewTuple(f*t.X, f*t.Y, f*t.Z, f*t.W)
}

func (t Tuple) Divide(f float64) Tuple {
	if math.Abs(f) < epsilon {
		panic("division by near-zero value")
	}
	return t.Times(1 / f)
}

func (t Tuple) Length() float64 {
	return math.Sqrt(t.X*t.X + t.Y*t.Y + t.Z*t.Z + t.W*t.W)
}

func (t Tuple) Normalized() Tuple {
	length := t.Length()
	if length < epsilon {
		panic("cannot normalize tuple with zero-length")
	}
	return t.Times(1 / length)
}

func (t1 Tuple) Dot(t2 Tuple) float64 {
	if !t1.IsVector() || !t2.IsVector() {
		panic("Dot cannot operate on non-vectors.")
	}
	return t1.X*t2.X + t1.Y*t2.Y + t1.Z*t2.Z
}

func (t1 Tuple) Cross(t2 Tuple) Tuple {
	if !t1.IsVector() || !t2.IsVector() {
		panic("Cross cannot operate on non-vectors.")
	}
	return NewVector(
		t1.Y*t2.Z-t1.Z*t2.Y,
		t1.Z*t2.X-t1.X*t2.Z,
		t1.X*t2.Y-t1.Y*t2.X,
	)
}
