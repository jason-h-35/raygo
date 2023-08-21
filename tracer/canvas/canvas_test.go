package canvas

import (
	"strings"
	"testing"
)

func Test_NewCanvas(t *testing.T) {
	xMax, yMax := 10, 20
	canvas := NewCanvas(xMax, yMax)
	xLen := len(canvas.Image)
	if xLen != xMax {
		t.Errorf("expected len of canvas.Image to be %v but instead was %v", xMax, xLen)
	}
	for i := range canvas.Image {
		yLen := len(canvas.Image[i])
		if yLen != yMax {
			t.Errorf("expected len of canvas.Image to be %v but instead was %v", yMax, yLen)
		}
	}
	for i := range canvas.Image {
		for j := range canvas.Image[i] {
			if !canvas.Image[i][j].Equals(Black) {
				t.Errorf("expected NewCanvas Canvas to be %v on every pixel but was %v on (%v, %v)", Black, canvas.Image[i][j], i, j)
			}
		}
	}
}

func Test_WritePixel(t *testing.T) {
	c := NewCanvas(10, 20)
	red := NewColor(1, 0, 0)
	c.WritePixel(2, 3, red)
	result := c.Image[2][3]
	if result != red {
		t.Errorf("expected WritePixel pixel to be %v but was %v", red, result)
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

func Test_PPMStr(t *testing.T) {
	c := NewCanvas(5, 3)
	ppm := c.PPMStr(255)
	// header test
	header := "P3\n5 3\n255\n"
	if !strings.HasPrefix(ppm, header) {
		t.Errorf("PPM header does not match.\nExpected ppm:\n%v\nto begin with:\n%v\n", ppm, header)
	}
	// data test
	overRed := NewColor(1.5, 0, 0)
	halfGreen := NewColor(0, 0.5, 0)
	underBlue := NewColor(-0.5, 0, 1)
	c.WritePixel(0, 0, overRed)
	c.WritePixel(2, 1, halfGreen)
	c.WritePixel(4, 2, underBlue)
	ppm = c.PPMStr(255)
	result := strings.ReplaceAll(ppm, "\n", "")
	expect := []string{
		"255 0 0 0 0 0 0 0 0 0 0 0 0 0 0 ",
		"0 0 0 0 0 0 0 128 0 0 0 0 0 0 0 ",
		"0 0 0 0 0 0 0 0 0 0 0 0 0 0 255",
	}
	if !strings.Contains(result, strings.Join(expect, "")) {
		t.Errorf("Data portion of PPM seems incorrect. Expected \n%v\nto contain %v", ppm, expect)
	}
}
