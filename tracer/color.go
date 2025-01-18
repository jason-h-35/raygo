package tracer

import (
	"math"
)

type Color struct {
	R float64
	G float64
	B float64
}

var Black Color = NewColor(0, 0, 0)
var White Color = NewColor(1, 1, 1)

func NewColor(R, G, B float64) Color {
	return Color{R, G, B}
}

func (c Color) RGBA() (r, g, b, a uint32) {
	// Convert from float64 [0,1] to uint32 [0,65535]
	r = uint32(c.R * 65535)
	g = uint32(c.G * 65535)
	b = uint32(c.B * 65535)
	a = 65535
	return
}

func (c1 Color) Equals(c2 Color) bool {
	if math.Abs(c1.R-c2.R) > epsilon {
		return false
	}
	if math.Abs(c1.G-c2.G) > epsilon {
		return false
	}
	if math.Abs(c1.B-c2.B) > epsilon {
		return false
	}
	return true
}

func (c1 Color) Plus(c2 Color) Color {
	return NewColor(c1.R+c2.R, c1.G+c2.G, c1.B+c2.B)
}

func (c1 Color) Minus(c2 Color) Color {
	return NewColor(c1.R-c2.R, c1.G-c2.G, c1.B-c2.B)
}

func (c Color) Times(f float64) Color {
	return NewColor(f*c.R, f*c.G, f*c.B)
}

func (c1 Color) Hadamard(c2 Color) Color {
	return NewColor(c1.R*c2.R, c1.G*c2.G, c1.B*c2.B)
}

func (c Color) ToPPMRange(maximum int) Color {
	c = c.Times(float64(maximum))
	// clamp each of c.R, c.G, c.B into range [0, maximum]
	c.R = math.Min(math.Max(c.R, 0), float64(maximum))
	c.G = math.Min(math.Max(c.G, 0), float64(maximum))
	c.B = math.Min(math.Max(c.B, 0), float64(maximum))
	return c.Round()
}

func (c Color) Round() Color {
	c.R = math.Round(c.R)
	c.G = math.Round(c.G)
	c.B = math.Round(c.B)
	return c
}

func (c Color) asFloats() (float64, float64, float64) {
	return c.R, c.G, c.B
}

func (c Color) asInts() (int, int, int) {
	return int(c.R), int(c.G), int(c.B)
}
