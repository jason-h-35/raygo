package canvas

import (
// "image"
// "image/color"
)

type Color struct {
	R float64
	G float64
	B float64
}

var Black Color = NewColor(0, 0, 0)
var White Color = NewColor(1, 1, 1)

var eps float64 = 1.e-5

func NewColor(R, G, B float64) Color {
	return Color{R, G, B}
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func (c1 Color) Equals(c2 Color) bool {
	if abs(c1.R-c2.R) > eps {
		return false
	}
	if abs(c1.G-c2.G) > eps {
		return false
	}
	if abs(c1.B-c2.B) > eps {
		return false
	}
	return true
}

func (c1 Color) Plus(c2 Color) Color {
	return NewColor(c1.R+c1.R, c1.G+c2.G, c1.B+c2.B)
}

func (c1 Color) Minus(c2 Color) Color {
	return NewColor(c1.R-c1.R, c1.G-c2.G, c1.B-c2.B)
}

func (c Color) Times(f float64) Color {
	return NewColor(f*c.R, f*c.G, f*c.B)
}

func (c1 Color) Hadamard(c2 Color) Color {
	return NewColor(c1.R*c1.R, c1.G*c2.G, c1.B*c2.B)
}
