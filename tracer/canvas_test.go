package tracer

import (
	"image"
	"image/color"
	"image/draw"
	"os"
	"path/filepath"
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
	c.SetLinear(2, 3, red)
	result := c.image[2][3]
	if !result.Equals(red) {
		t.Errorf("expected WritePixel pixel to be %v but was %v", red, result)
	}
}

func Test_ReadPixel(t *testing.T) {
	c := Canvas{
		image: [][]LinearColor{
			{ColorRed, ColorGreen},
			{ColorBlue, ColorWhite},
		},
	}
	for i := range c.image {
		for j := range c.image[i] {
			if !c.AtLinear(i, j).Equals(c.image[i][j]) {
				t.Errorf("c.ReadPixel does not match c.image on c.image[%v][%v]", i, j)
			}
		}
	}
}

func Test_PPMStr(t *testing.T) {
	c := NewCanvas(5, 3)
	ppm := c.PPMStr(255)
	header := "P3\n5 3\n255\n"
	if !strings.HasPrefix(ppm, header) {
		t.Errorf("PPM header does not match.\nExpected ppm:\n%v\nto begin with:\n%v\n", ppm, header)
	}

	doubleRed := NewColorFromFloat64(1.5, 0, 0)
	halfGreen := NewColorFromFloat64(0, 0.5, 0)
	underBlue := NewColorFromFloat64(-0.5, 0, 1)
	c.SetLinear(0, 0, doubleRed)
	c.SetLinear(2, 1, halfGreen)
	c.SetLinear(4, 2, underBlue)
	ppm = c.PPMStr(255)
	result := strings.Fields(strings.TrimPrefix(ppm, header))
	expect := []string{
		"255 0 0 0 0 0 0 0 0 0 0 0 0 0 0",
		"0 0 0 0 0 0 0 128 0 0 0 0 0 0 0",
		"0 0 0 0 0 0 0 0 0 0 0 0 0 0 255",
	}
	want := strings.Fields(strings.Join(expect, "\n"))
	if strings.Join(result, " ") != strings.Join(want, " ") {
		t.Errorf("Data portion of PPM seems incorrect.\nGot:\n%v\nWant:\n%v\n", result, want)
	}

	ppmLines := strings.Split(ppm, "\n")
	maxLen := 70
	for lineNum, line := range ppmLines {
		if len(line) > maxLen {
			t.Errorf("Line too long for PPM format. expected max of %v but was instead %v at line %v\n%v",
				maxLen, len(line), lineNum, line)
		}
	}

	lastChar := ppm[len(ppm)-1]
	if lastChar != '\n' {
		t.Errorf("Expected last char of PPM to be newline, but was %v instead.", lastChar)
	}
}

func TestZeroSizeCanvas(t *testing.T) {
	c := NewCanvas(0, 0)
	bounds := c.Bounds()
	if bounds.Max.X != 0 || bounds.Max.Y != 0 {
		t.Errorf("expected zero bounds, got %v", bounds)
	}

	if got := c.At(0, 0); got == nil {
		t.Errorf("expected non-nil color for out-of-bounds read")
	}

	c.Set(0, 0, ColorWhite)
	if ppm := c.PPMStr(255); ppm != "P3\n0 0\n255\n\n" {
		t.Errorf("unexpected zero-sized PPM output: %q", ppm)
	}
}

func TestImageInterface(t *testing.T) {
	var _ image.Image = &Canvas{}

	canvas := NewCanvas(100, 100)
	bounds := canvas.Bounds()

	if bounds.Min.X != 0 || bounds.Min.Y != 0 {
		t.Errorf("Expected bounds min (0,0), got (%d,%d)", bounds.Min.X, bounds.Min.Y)
	}
	if bounds.Max.X != 100 || bounds.Max.Y != 100 {
		t.Errorf("Expected bounds max (100,100), got (%d,%d)", bounds.Max.X, bounds.Max.Y)
	}

	if canvas.ColorModel() != color.RGBA64Model {
		t.Error("ColorModel() should return color.RGBA64Model")
	}

	_ = canvas.At(50, 50)
	_ = canvas.At(-1, -1)
	_ = canvas.At(1000, 1000)
}

func TestDrawImageInterface(t *testing.T) {
	var _ draw.Image = &Canvas{}

	canvas := NewCanvas(100, 100)
	testColor := ColorRed

	canvas.Set(50, 50, testColor)

	if got := canvas.AtLinear(50, 50); !got.Equals(testColor) {
		t.Errorf("Set() color mismatch: got %v, want %v", got, testColor)
	}

	canvas.Set(-1, -1, testColor)
	canvas.Set(1000, 1000, testColor)
}

func TestCanvasImageExports(t *testing.T) {
	canvas := NewCanvas(10, 10)
	canvas.SetLinear(5, 5, ColorRed)
	outDir := t.TempDir()

	if err := canvas.PNGFile(filepath.Join(outDir, "out.png")); err != nil {
		t.Fatalf("PNGFile() failed: %v", err)
	}
	if err := canvas.JPEGFile(filepath.Join(outDir, "out.jpg"), 95); err != nil {
		t.Fatalf("JPEGFile() failed: %v", err)
	}
	if _, err := canvas.PPMFile(255, filepath.Join(outDir, "out.ppm")); err != nil {
		t.Fatalf("PPMFile() failed: %v", err)
	}

	for _, name := range []string{"out.png", "out.jpg", "out.ppm"} {
		path := filepath.Join(outDir, name)
		info, err := os.Stat(path)
		if err != nil {
			t.Fatalf("expected %s to exist: %v", name, err)
		}
		if info.Size() == 0 {
			t.Fatalf("expected %s to be non-empty", name)
		}
	}
}
