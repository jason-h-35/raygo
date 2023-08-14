package canvas

import (
	"testing"
)

func Test_Equals(t *testing.T) {
	c1 := NewColor(1.0/3, 1.0/3, 1.0/3)
	c2 := NewColor(0.333333, 0.333333, 0.333333)
	if !c1.Equals(c2) {
		t.Errorf("%v Equal %v should be true, but was false", c1, c2)
	}
}

func Test_Plus(t *testing.T) {
	c1 := NewColor(0.9, 0.6, 0.75)
	c2 := NewColor(0.7, 0.1, 0.25)
	expect := NewColor(1.6, 0.7, 1.0)
	result := c1.Plus(c2)
	if !expect.Equals(result) {
		t.Errorf("%v Plus %v should be %v, but was %v", c1, c2, expect, result)
	}
}

func Test_Minus(t *testing.T) {
	c1 := NewColor(0.9, 0.6, 0.75)
	c2 := NewColor(0.7, 0.1, 0.25)
	expect := NewColor(0.2, 0.5, 0.5)
	result := c1.Minus(c2)
	if !expect.Equals(result) {
		t.Errorf("%v Minus %v should be %v, but was %v", c1, c2, expect, result)
	}
}

func Test_Times(t *testing.T) {
	c := NewColor(0.2, 0.3, 0.4)
	f := 2.0
	expect := NewColor(0.4, 0.6, 0.8)
	result := c.Times(f)
	if !expect.Equals(result) {
		t.Errorf("%v Times %v should be %v, but was %v", c, f, expect, result)
	}
}

func Test_Hadamard(t *testing.T) {
	c1 := NewColor(1, 0.2, 0.4)
	c2 := NewColor(0.9, 1, 0.1)
	expect := NewColor(0.9, 0.2, 0.04)
	result := c1.Hadamard(c2)
	if !expect.Equals(result) {
		t.Errorf("%v Times %v should be %v, but was %v", c1, c2, expect, result)
	}
}
