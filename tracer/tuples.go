package tracer

import (
	"math"
)

const eps = 1.e-5

type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64
}

func NewTuple(x float64, y float64, z float64, w float64) Tuple {
	return Tuple{x, y, z, w}
}

func Point(x, y, z float64) Tuple {
	return NewTuple(x, y, z, 1)
}

func Vector(x, y, z float64) Tuple {
	return NewTuple(x, y, z, 0)
}

func (t Tuple) IsVector() bool {
	if t.W == 0 {
		return true
	}
	return false
}

func (t Tuple) IsPoint() bool {
	if t.W == 1 {
		return true
	}
	return false
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func (t1 Tuple) Equals(t2 Tuple) bool {
	if abs(t1.X-t2.X) > eps {
		return false
	}
	if abs(t1.Y-t2.Y) > eps {
		return false
	}
	if abs(t1.Z-t2.Z) > eps {
		return false
	}
	if t1.W != t2.W {
		return false
	}
	return true
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
	if abs(f) < eps {
		panic("divisor too close to 0.0")
	}
	return t.Times(1 / f)
}

func (t Tuple) Length() float64 {
	return math.Sqrt(t.X*t.X + t.Y*t.Y + t.Z*t.Z + t.W*t.W)
}

func (t Tuple) Normalized() Tuple {
	if t.Length() < eps {
		panic("divisor too close to 0.0")
	}
	return t.Times(1 / t.Length())
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
	return Vector(
		t1.Y*t2.Z-t1.Z*t2.Y,
		t1.Z*t2.X-t1.X*t2.Z,
		t1.X*t2.Y-t1.Y*t2.X,
	)
}
