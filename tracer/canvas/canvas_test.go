package canvas

import (
	"testing"
)

func Test_NewCanvas(t *testing.T) {
	canvas := NewCanvas(10, 5)
	result := canvas.Image[9][4]
	expect := NewColor(0, 0, 0)
	if !result.Equals(expect) {
		t.Errorf("expected pixel on new Canvas to be %v but was %v", expect, result)
	}
}

func Test_WritePixel(t *testing.T) {
	c := NewCanvas(20, 40)
	white := NewColor(1, 1, 1)
	c.WritePixel(10, 20, white)
	result := c.Image[10][20]
	if result != white {
		t.Errorf("expected WritePixel pixel to be %v but was %v", white, result)
	}
}

func Test_ReadPixel(t *testing.T) {
	c := Canvas{
		Image: [][]Color{
			{NewColor(1, 0, 0), NewColor(0, 1, 0)},
			{NewColor(0, 0, 1), NewColor(1, 1, 1)},
		},
	}
	for i := range c.Image {
		for j := range c.Image[i] {
			if !c.ReadPixel(i, j).Equals(c.Image[i][j]) {
				t.Errorf("c.ReadPixel does not match c.image on c.image[%v][%v]", i, j)
			}
		}
	}
}
