package tracer

import (
	"math"
)

type LinearColor struct{ R, G, B float32 }

const colorEps = 1e-5

var (
	ColorBlack   = LinearColor{0, 0, 0}
	ColorGrey    = LinearColor{0.0625, 0.0625, 0.0625}
	ColorBlue    = LinearColor{0, 0, 1}
	ColorGreen   = LinearColor{0, 1, 0}
	ColorCyan    = LinearColor{0, 1, 1}
	ColorRed     = LinearColor{1, 0, 0}
	ColorMagenta = LinearColor{1, 0, 1}
	ColorYellow  = LinearColor{1, 1, 0}
	ColorWhite   = LinearColor{1, 1, 1}
)

func NewColorFromFloat64(r, g, b float64) LinearColor {
	return LinearColor{
		R: float32(max(r, 0)),
		G: float32(max(g, 0)),
		B: float32(max(b, 0)),
	}
}

func clamp01(v float32) float32 {
	return min(max(v, 0), 1)
}

func linearToUint16(v float32) uint16 {
	v = clamp01(v)
	return uint16(math.Round(float64(v * 65535)))
}

func (c LinearColor) RGBA() (r, g, b, a uint32) {
	r = uint32(linearToUint16(c.R))
	g = uint32(linearToUint16(c.G))
	b = uint32(linearToUint16(c.B))
	a = 0xffff
	return // r,g,b,a
}

func (c1 LinearColor) Equals(c2 LinearColor) bool {
	return c1.Distance(c2) <= colorEps
}

func (c1 LinearColor) Distance(c2 LinearColor) float32 {
	dr := float32(math.Abs(float64(c1.R - c2.R)))
	dg := float32(math.Abs(float64(c1.G - c2.G)))
	db := float32(math.Abs(float64(c1.B - c2.B)))
	return dr + dg + db
}

func (c1 LinearColor) Plus(c2 LinearColor) LinearColor {
	return LinearColor{
		R: c1.R + c2.R,
		G: c1.G + c2.G,
		B: c1.B + c2.B,
	}
}

func (c1 LinearColor) Minus(c2 LinearColor) LinearColor {
	return LinearColor{
		R: c1.R - c2.R,
		G: c1.G - c2.G,
		B: c1.B - c2.B,
	}
}

func (c LinearColor) Times(f float32) LinearColor {
	return LinearColor{
		R: c.R * f,
		G: c.G * f,
		B: c.B * f,
	}
}

func (c1 LinearColor) Hadamard(c2 LinearColor) LinearColor {
	return LinearColor{
		R: c1.R * c2.R,
		G: c1.G * c2.G,
		B: c1.B * c2.B,
	}
}

func (c LinearColor) ToPPMRange(limit uint64) (uint64, uint64, uint64) {
	scale := float64(limit)
	r := uint64(math.Round(float64(clamp01(c.R)) * scale))
	g := uint64(math.Round(float64(clamp01(c.G)) * scale))
	b := uint64(math.Round(float64(clamp01(c.B)) * scale))
	return r, g, b
}

func (c LinearColor) AsFloats() (float32, float32, float32) {
	return c.R, c.G, c.B
}

func (c LinearColor) AsInts() (int, int, int) {
	r, g, b, _ := c.RGBA()
	return int(r), int(g), int(b)
}
