package canvas

import (
	"fmt"
	"math"
	"strings"
)

type Canvas struct {
	image [][]Color
}

func NewCanvas(width, height int) Canvas {
	image := make([][]Color, width, width)
	for i := range image {
		image[i] = make([]Color, height, height)
	}
	return Canvas{image}
}

func (c *Canvas) Width() int {
	return len(c.image)
}

func (c *Canvas) Height() int {
	return len(c.image[0]) // how to check if all sub-arrays are the same length?
}

func (c *Canvas) WritePixel(x int, y int, color Color) {
	canvas := *c
	canvas.image[x][y] = color
}

func (c *Canvas) ReadPixel(x, y int) Color {
	canvas := *c
	return canvas.image[x][y]
}

func (c *Canvas) ToPPM() string {
	colorRepMax := 255.0
	ppmHeader := fmt.Sprintf("\nP3\n%v %v\n%v\n", c.Width, c.Height, colorMax)
	var b strings.Builder
	for _, row := range c.image {
		for _, color := range row {
			colorRep := string(int(math.Round(color.Times(colorRepMax))))
			b.WriteString(colorRep)
		}
	}
	b.WriteRune('\n')
	return b.String()
}
