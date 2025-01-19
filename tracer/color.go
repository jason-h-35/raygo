package tracer

import (
	"math"
)

type HDRColor struct{ R, G, B uint64 }

const colorEps = 3

var (
	ColorBlack   = HDRColor{0x0000, 0x0000, 0x0000}
	ColorBlue    = HDRColor{0x0000, 0x0000, 0xffff}
	ColorGreen   = HDRColor{0x0000, 0xffff, 0x0000}
	ColorCyan    = HDRColor{0x0000, 0xffff, 0xffff}
	ColorRed     = HDRColor{0xffff, 0x0000, 0x0000}
	ColorMagenta = HDRColor{0xffff, 0x0000, 0xffff}
	ColorYellow  = HDRColor{0xffff, 0xffff, 0x0000}
	ColorWhite   = HDRColor{0xffff, 0xffff, 0xffff}
)

func NewColorFromFloat64(r, g, b float64) HDRColor {
	if r < 0 || g < 0 || b < 0 {
		panic("tracer.NewColorFromFloat64: negative color components not allowed")
	}
	return HDRColor{
		uint64(r * 0xffff),
		uint64(g * 0xffff),
		uint64(b * 0xffff),
	}
}

func (c HDRColor) RGBA() (r, g, b, a uint32) {
	// Convert from uint64 to uint32 [0,65535]
	r = uint32(min(c.R, 0xffff))
	g = uint32(min(c.G, 0xffff))
	b = uint32(min(c.B, 0xffff))
	a = 0xffff
	return // r,g,b,a
}

func (c1 HDRColor) Equals(c2 HDRColor) bool {
	return c1.Distance(c2) == 0
}

func (c1 HDRColor) Distance(c2 HDRColor) uint64 {
	var dr, dg, db uint64
	if c1.R > c2.R {
		dr = c1.R - c2.R
	} else {
		dr = c2.R - c1.R
	}
	if c1.G > c2.G {
		dg = c1.G - c2.G
	} else {
		dg = c2.G - c1.G
	}
	if c1.B > c2.B {
		db = c1.B - c2.B
	} else {
		db = c2.B - c1.B
	}
	return dr + dg + db
}

func (c1 HDRColor) Plus(c2 HDRColor) HDRColor {
	var r, g, b uint64
	if math.MaxUint64-c1.R < c2.R {
		r = math.MaxUint64
	} else {
		r = c1.R + c2.R
	}
	if math.MaxUint64-c1.G < c2.G {
		g = math.MaxUint64
	} else {
		g = c1.G + c2.G
	}
	if math.MaxUint64-c1.B < c2.B {
		b = math.MaxUint64
	} else {
		b = c1.B + c2.B
	}
	return HDRColor{r, g, b}
}

func (c1 HDRColor) Minus(c2 HDRColor) HDRColor {
	var r, g, b uint64
	if c1.R < c2.R {
		r = 0
	} else {
		r = c1.R - c2.R
	}
	if c1.G < c2.G {
		g = 0
	} else {
		g = c1.G - c2.G
	}
	if c1.B < c2.B {
		b = 0
	} else {
		b = c1.B - c2.B
	}
	return HDRColor{r, g, b}
}

func (c HDRColor) Times(f uint64) HDRColor {
	var r, g, b uint64
	if f != 0 && c.R > math.MaxUint64/f {
		r = math.MaxUint64
	} else {
		r = c.R * f
	}
	if f != 0 && c.G > math.MaxUint64/f {
		g = math.MaxUint64
	} else {
		g = c.G * f
	}
	if f != 0 && c.B > math.MaxUint64/f {
		b = math.MaxUint64
	} else {
		b = c.B * f
	}
	return HDRColor{r, g, b}
}

// Hadamard requires colors be mapped into the [0-1] floating point color space.
// For large uint64, we may not be able to cast to float64, so check first.
// These are then mapped back
func (c1 HDRColor) Hadamard(c2 HDRColor) HDRColor {
	const maxExactFloat = 1<<53 - 1
	if c1.R > maxExactFloat || c2.R > maxExactFloat ||
		c1.G > maxExactFloat || c2.G > maxExactFloat ||
		c1.B > maxExactFloat || c2.B > maxExactFloat {
		panic("HDRColor.Hadamard: color components too large for precise float64 conversion")
	}
	fr := float64(c1.R) * float64(c2.R) / 0xffff
	fg := float64(c1.G) * float64(c2.G) / 0xffff
	fb := float64(c1.B) * float64(c2.B) / 0xffff
	r := uint64(min(fr, math.MaxUint64))
	g := uint64(min(fg, math.MaxUint64))
	b := uint64(min(fb, math.MaxUint64))
	return HDRColor{r, g, b}
}

func (c HDRColor) ToPPMRange(limit uint64) HDRColor {
	const fullColor = uint64(0xFFFF)
	// Clamp each component to 0xFFFF, mapping into RGBA space
	r := min(c.R, fullColor)
	g := min(c.G, fullColor)
	b := min(c.B, fullColor)
	// Now scale the clamped values to the target range
	r = (r * limit) / fullColor
	g = (g * limit) / fullColor
	b = (b * limit) / fullColor
	return HDRColor{r, g, b}
}

func (c HDRColor) AsFloats() (float64, float64, float64) {
	return float64(c.R), float64(c.G), float64(c.B)
}

func (c HDRColor) AsInts() (int, int, int) {
	return int(c.R), int(c.G), int(c.B)
}
