package tracer

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
	"strings"
)

type Canvas struct {
	image [][]HDRColor
}

func NewCanvas(xMax, yMax int) Canvas {
	image := make([][]HDRColor, xMax)
	for i := range image {
		image[i] = make([]HDRColor, yMax)
	}
	return Canvas{image}
}

func (c *Canvas) At(x, y int) color.Color {
	if !(image.Point{x, y}.In(c.Bounds())) {
		return HDRColor{}
	}
	return c.image[x][y]
}

func (c *Canvas) AtHDR(x, y int) HDRColor {
	if !(image.Point{x, y}.In(c.Bounds())) {
		return HDRColor{}
	}
	return c.image[x][y]
}

func (c *Canvas) ColorModel() color.Model {
	return color.RGBAModel
}

func (c *Canvas) Bounds() image.Rectangle {
	return image.Rect(0, 0, len(c.image), len(c.image[0]))
}

func (c *Canvas) Set(x, y int, clr color.Color) {
	bounds := c.Bounds().Max
	if x < 0 || x >= bounds.X || y < 0 || y >= bounds.Y {
		return
	}
	r, g, b, _ := clr.RGBA()
	// Convert from uint32 [0-65535] to float64 [0-1]
	c.image[x][y] = HDRColor{uint64(r), uint64(g), uint64(b)}
}

func (c *Canvas) SetColor(x, y int, clr HDRColor) {
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
			pixel := c.image[x][y].ToPPMRange(maxColorVal)
			r, g, b := int(pixel.R), int(pixel.G), int(pixel.B)
			for _, v := range []int{r, g, b} {
				s := strconv.Itoa(v)
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
	file, fileErr := os.Create(writePath)
	defer file.Close()
	if fileErr != nil {
		fmt.Println(fileErr)
		return 0, fileErr
	}
	return fmt.Fprintf(file, "%v", ppmStr)
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
