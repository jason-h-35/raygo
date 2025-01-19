package tracer

import (
	"image/color"
	"testing"
)

func Test_Equals(t *testing.T) {
	var x uint64 = 0x5555
	var y uint64 = 21845 // same as 0x5555
	c := HDRColor{x, x, x}
	cEq := HDRColor{y, y, y}
	cX := HDRColor{0, x, x}
	cY := HDRColor{x, 0, x}
	cZ := HDRColor{x, x, 0}
	if !c.Equals(cEq) {
		t.Errorf("%v should Equal %v", c, cEq)
	}
	if c.Equals(cX) {
		t.Errorf("%v should not Equal %v", c, cX)
	}
	if c.Equals(cY) {
		t.Errorf("%v should not Equal %v", c, cY)
	}
	if c.Equals(cZ) {
		t.Errorf("%v should not Equal %v", c, cZ)
	}
}

func Test_Plus(t *testing.T) {
	c1 := NewColorFromFloat64(0.9, 0.6, 0.75)
	c2 := NewColorFromFloat64(0.7, 0.1, 0.25)
	expect := NewColorFromFloat64(1.6, 0.7, 1.0)
	result := c1.Plus(c2)
	if d := expect.Distance(result); d > colorEps {
		t.Errorf("%v Plus %v should be %v, but was %v. A distance of %v", c1, c2, expect, result, d)
	}
}

func Test_Minus(t *testing.T) {
	c1 := NewColorFromFloat64(0.9, 0.6, 0.75)
	c2 := NewColorFromFloat64(0.7, 0.1, 0.25)
	expect := NewColorFromFloat64(0.2, 0.5, 0.5)
	result := c1.Minus(c2)
	if d := expect.Distance(result); d > colorEps {
		t.Errorf("%v Minus %v should be %v, but was %v. A distance of %v", c1, c2, expect, result, d)
	}
}

func Test_Times(t *testing.T) {
	c := NewColorFromFloat64(0.2, 0.3, 0.4)
	var f uint64 = 2
	expect := NewColorFromFloat64(0.4, 0.6, 0.8)
	result := c.Times(f)
	if d := expect.Distance(result); d > colorEps {
		t.Errorf("%v Times %v should be %v, but was %v. A distance of %v", c, f, expect, result, d)
	}
}

func Test_Hadamard(t *testing.T) {
	c1 := NewColorFromFloat64(1, 0.2, 0.4)
	c2 := NewColorFromFloat64(0.9, 1, 0.1)
	expect := NewColorFromFloat64(0.9, 0.2, 0.04)
	result := c1.Hadamard(c2)
	if d := expect.Distance(result); d > colorEps {
		t.Errorf("%v Hadamard %v should be %v, but was %v. A distance of %v", c1, c2, expect, result, d)
	}
}

func TestColorImplementsColorInterface(t *testing.T) {
	// Static type assertion at compile time
	var _ color.Color = HDRColor{} // Will fail to compile if Color doesn't implement color.Color

	// Runtime behavior test
	c := NewColorFromFloat64(0.5, 0.25, 0.75)
	r, g, b, a := c.RGBA()

	// Expected values: 0.5 * 65535 ≈ 32767, 0.25 * 65535 ≈ 16383, 0.75 * 65535 ≈ 49151
	expectedR := uint32(32767)
	expectedG := uint32(16383)
	expectedB := uint32(49151)
	expectedA := uint32(65535) // Full opacity

	if r != expectedR {
		t.Errorf("Red channel incorrect. Got %d, want %d", r, expectedR)
	}
	if g != expectedG {
		t.Errorf("Green channel incorrect. Got %d, want %d", g, expectedG)
	}
	if b != expectedB {
		t.Errorf("Blue channel incorrect. Got %d, want %d", b, expectedB)
	}
	if a != expectedA {
		t.Errorf("Alpha channel incorrect. Got %d, want %d", a, expectedA)
	}
}
