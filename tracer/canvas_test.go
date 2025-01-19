package tracer

import (
	"image"
	"image/draw"
	"strings"
	"testing"
)

func Test_NewCanvas(t *testing.T) {
	xMax, yMax := 10, 20
	canvas := NewCanvas(xMax, yMax)
	xLen := len(canvas.image)
	if xLen != xMax {
		t.Errorf("expected len of canvas.Image to be %v but instead was %v", xMax, xLen)
	}
	for i := range canvas.image {
		yLen := len(canvas.image[i])
		if yLen != yMax {
			t.Errorf("expected len of canvas.Image to be %v but instead was %v", yMax, yLen)
		}
	}
	for i := range canvas.image {
		for j := range canvas.image[i] {
			if !canvas.image[i][j].Equals(ColorBlack) {
				t.Errorf("expected NewCanvas Canvas to be %v on every pixel but was %v on (%v, %v)", ColorBlack, canvas.image[i][j], i, j)
			}
		}
	}
}

func Test_WritePixel(t *testing.T) {
	c := NewCanvas(10, 20)
	red := NewColorFromFloat64(1, 0, 0)
	c.SetColor(2, 3, red)
	result := c.image[2][3]
	if result != red {
		t.Errorf("expected WritePixel pixel to be %v but was %v", red, result)
	}
}

func Test_ReadPixel(t *testing.T) {
	c := Canvas{
		image: [][]HDRColor{
			{ColorRed, ColorGreen},
			{ColorBlue, ColorWhite},
		},
	}
	for i := range c.image {
		for j := range c.image[i] {
			if c.AtHDR(i, j) != c.image[i][j] {
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
	// pixel data test
	doubleRed := NewColorFromFloat64(1.5, 0, 0)
	halfGreen := NewColorFromFloat64(0, 0.5, 0)
	underBlue := NewColorFromFloat64(0, 0, 1)
	c.SetColor(0, 0, doubleRed)
	c.SetColor(2, 1, halfGreen)
	c.SetColor(4, 2, underBlue)
	ppm = c.PPMStr(255)
	result := strings.ReplaceAll(strings.TrimPrefix(ppm, header), "\n", " ")
	expect := []string{
		"255 0 0 0 0 0 0 0 0 0 0 0 0 0 0",
		"0 0 0 0 0 0 0 127 0 0 0 0 0 0 0",
		"0 0 0 0 0 0 0 0 0 0 0 0 0 0 255",
	}
	if result == strings.Join(expect, " ") {
		t.Errorf("Data portion of PPM seems incorrect.\nGot:\n%s\nWant:\n%s\n", result, expect)
	}
	// line length test
	ppmLines := strings.Split(ppm, "\n")
	maxLen := 70
	for lineNum, line := range ppmLines {
		if len(line) > maxLen {
			t.Errorf("Line too long for PPM format. expected max of %v but was instead %v at line %v\n%v",
				maxLen, len(line), lineNum, line)
		}
	}
	// last char is newline test
	lastChar := ppm[len(ppm)-1]
	if lastChar != '\n' {
		t.Errorf("Expected last char of PPM to be newline, but was %v instead.", lastChar)
	}
}

// TestImageInterface verifies Canvas implements image.Image
func TestImageInterface(t *testing.T) {
	// Compile-time verification that Canvas implements image.Image
	var _ image.Image = &Canvas{}

	canvas := NewCanvas(100, 100)
	bounds := canvas.Bounds()

	// Test Bounds method
	if bounds.Min.X != 0 || bounds.Min.Y != 0 {
		t.Errorf("Expected bounds min (0,0), got (%d,%d)", bounds.Min.X, bounds.Min.Y)
	}
	if bounds.Max.X != 100 || bounds.Max.Y != 100 {
		t.Errorf("Expected bounds max (100,100), got (%d,%d)", bounds.Max.X, bounds.Max.Y)
	}

	// Test ColorModel method
	if canvas.ColorModel() != canvas.ColorModel() {
		t.Error("ColorModel() should return consistent results")
	}

	// Test At method - should not panic on valid coordinates
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("At() panicked on valid coordinates: %v", r)
		}
	}()
	_ = canvas.At(50, 50)

	// Test At method with out of bounds - should not panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("At() panicked on out of bounds coordinates: %v", r)
		}
	}()
	_ = canvas.At(-1, -1)
	_ = canvas.At(1000, 1000)
}

// TestDrawImageInterface verifies Canvas implements draw.Image
func TestDrawImageInterface(t *testing.T) {
	// Compile-time verification that Canvas implements draw.Image
	var _ draw.Image = &Canvas{}

	canvas := NewCanvas(100, 100)
	testColor := HDRColor{0xffff, 0, 0} // Red

	// Test Set method within bounds
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Set() panicked on valid coordinates: %v", r)
		}
	}()
	canvas.Set(50, 50, testColor)

	// Verify pixel was set
	if got := canvas.At(50, 50); got != testColor {
		t.Errorf("Set() color mismatch: got %v, want %v", got, testColor)
	}

	// Test Set method out of bounds - should not panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Set() panicked on out of bounds coordinates: %v", r)
		}
	}()
	canvas.Set(-1, -1, testColor)
	canvas.Set(1000, 1000, testColor)
}
