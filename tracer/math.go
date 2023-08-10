package tracer

//  import (
// 	"fmt"
// 	"image/color"
//
// 	rl "github.com/gen2brain/raylib-go/raylib"
// )

const eps = 1.e-5

type Tuple struct {
	X float32
	Y float32
	Z float32
	W float32
}

func NewTuple(x float32, y float32, z float32, w float32) Tuple {
	return Tuple{x, y, z, w}
}

func Point(x, y, z float32) Tuple {
	return NewTuple(x, y, z, 1)
}

func Vector(x, y, z float32) Tuple {
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

func abs(x float32) float32 {
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

func Minus(t Tuple) Tuple {
	return t.Times(-1)
}

func (t1 Tuple) Plus(t2 Tuple) Tuple {
	return NewTuple(
		t1.X+t2.X, t1.Y+t2.Y,
		t1.Z+t2.Z, t1.W+t2.W,
	)
}

func (t1 Tuple) Minus(t2 Tuple) Tuple {
	return t1.Plus(Minus(t2))
}

func (t Tuple) Times(f float32) Tuple {
	return NewTuple(f*t.X, f*t.Y, f*t.Z, f*t.W)
}

func (t Tuple) Divide(f float32) Tuple {
	if abs(f) < eps {
		panic("divisor too close to 0.0")
	}
	return t.Times(1 / f)
}
