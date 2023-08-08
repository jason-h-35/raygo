package tracer

//  import (
// 	"fmt"
// 	"image/color"
//
// 	rl "github.com/gen2brain/raylib-go/raylib"
// )

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
