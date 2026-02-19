package tracer

import (
	"image/color"
	"testing"
)

func Test_Equals(t *testing.T) {
	c := NewColorFromFloat64(0.333333, 0.2, 0.8)
	cEq := NewColorFromFloat64(0.333334, 0.2, 0.8)
	cX := NewColorFromFloat64(0.4, 0.2, 0.8)
	cY := NewColorFromFloat64(0.333333, 0.4, 0.8)
	cZ := NewColorFromFloat64(0.333333, 0.2, 0.4)

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
	f := float32(2)
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
	var _ color.Color = LinearColor{}

	c := NewColorFromFloat64(0.5, 0.25, 0.75)
	r, g, b, a := c.RGBA()

	expectedR := uint32(32768)
	expectedG := uint32(16384)
	expectedB := uint32(49151)
	expectedA := uint32(65535)

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

func TestColorClampOnExport(t *testing.T) {
	c := LinearColor{R: 1.5, G: -0.1, B: 0.5}
	r, g, b, _ := c.RGBA()

	if r != 65535 {
		t.Errorf("expected red to clamp high, got %d", r)
	}
	if g != 0 {
		t.Errorf("expected green to clamp low, got %d", g)
	}
	if b != 32768 {
		t.Errorf("expected blue to map midpoint, got %d", b)
	}
}
