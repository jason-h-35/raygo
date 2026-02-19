package tracer

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
	"strconv"
	"strings"
)

type Canvas struct {
	image [][]LinearColor
}

func NewCanvas(xMax, yMax int) Canvas {
	image := make([][]LinearColor, xMax)
	for i := range image {
		image[i] = make([]LinearColor, yMax)
	}
	return Canvas{image}
}

func (c *Canvas) At(x, y int) color.Color {
	pixel := c.AtLinear(x, y)
	r, g, b, a := pixel.RGBA()
	return color.RGBA64{
		R: uint16(r),
		G: uint16(g),
		B: uint16(b),
		A: uint16(a),
	}
}

func (c *Canvas) AtLinear(x, y int) LinearColor {
	if !(image.Point{x, y}.In(c.Bounds())) {
		return LinearColor{}
	}
	return c.image[x][y]
}

func (c *Canvas) ColorModel() color.Model {
	return color.RGBA64Model
}

func (c *Canvas) Bounds() image.Rectangle {
	if len(c.image) == 0 {
		return image.Rect(0, 0, 0, 0)
	}
	return image.Rect(0, 0, len(c.image), len(c.image[0]))
}

func (c *Canvas) Set(x, y int, clr color.Color) {
	bounds := c.Bounds().Max
	if x < 0 || x >= bounds.X || y < 0 || y >= bounds.Y {
		return
	}
	r, g, b, _ := clr.RGBA()
	c.image[x][y] = LinearColor{
		R: float32(r) / 65535,
		G: float32(g) / 65535,
		B: float32(b) / 65535,
	}
}

func (c *Canvas) SetLinear(x, y int, clr LinearColor) {
	bounds := c.Bounds().Max
	if x < 0 || x >= bounds.X || y < 0 || y >= bounds.Y {
		return
	}
	c.image[x][y] = clr
}

func (c *Canvas) PPMStr(maxColorVal uint64) string {
	const (
		maxLineLen = 70 // per PPM standard
		pixelSep   = " "
		lineSep    = "\n"
	)
	bounds := c.Bounds().Max
	var builder strings.Builder
	// Write the header
	fmt.Fprintf(&builder, "P3\n%d %d\n%d\n", bounds.X, bounds.Y, maxColorVal)
	lineLen := 0
	for y := 0; y < bounds.Y; y++ {
		for x := 0; x < bounds.X; x++ {
			r, g, b := c.image[x][y].ToPPMRange(maxColorVal)
			for _, v := range []uint64{r, g, b} {
				s := strconv.FormatUint(v, 10)
				// Add newline if this component would exceed line length
				if lineLen > 0 && lineLen+len(s) >= maxLineLen {
					builder.WriteString(lineSep)
					lineLen = 0
				}
				// Add separator between components if not at start of line
				if lineLen > 0 {
					builder.WriteString(pixelSep)
					lineLen++
				}
				// Write the value
				builder.WriteString(s)
				lineLen += len(s)
			}
		}
	}
	builder.WriteString(lineSep)
	return builder.String()
}

func (c *Canvas) PPMFile(maxColorVal uint64, writePath string) (int, error) {
	ppmStr := c.PPMStr(maxColorVal)
	file, err := os.Create(writePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	return file.WriteString(ppmStr)
}

func (c *Canvas) PNGFile(writePath string) error {
	file, err := os.Create(writePath)
	if err != nil {
		return fmt.Errorf("failed to create PNG file: %w", err)
	}
	defer file.Close()

	if err := png.Encode(file, c); err != nil {
		return fmt.Errorf("failed to encode PNG: %w", err)
	}
	return nil
}

func (c *Canvas) JPEGFile(writePath string, quality int) error {
	file, err := os.Create(writePath)
	if err != nil {
		return fmt.Errorf("failed to create JPEG file: %w", err)
	}
	defer file.Close()

	if quality < 1 || quality > 100 {
		quality = 95
	}

	if err := jpeg.Encode(file, c, &jpeg.Options{Quality: quality}); err != nil {
		return fmt.Errorf("failed to encode JPEG: %w", err)
	}
	return nil
}
