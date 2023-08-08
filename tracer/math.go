package tracer

//  import (
// 	"fmt"
// 	"image/color"
//
// 	rl "github.com/gen2brain/raylib-go/raylib"
// )

var eps float32 = 1.e-5

type Tuple struct {
	X float32
	Y float32
	Z float32
	W int
}

func NewTuple(x float32, y float32, z float32, w int) Tuple {
	return Tuple{x, y, z, w}
}

func Point(x, y, z float32) Tuple {
	return Tuple{x, y, z, 1}
}

func Vector(x, y, z float32) Tuple {
	return Tuple{x, y, z, 0}
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
	if abs(t1.X-t2.X) > eps {
		return false
	}
	if abs(t1.X-t2.X) > eps {
		return false
	}
	if t1.W != t2.W {
		return false
	}
	return true
}
