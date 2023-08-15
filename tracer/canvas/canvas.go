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

func (c *Canvas) ToPPM(maxColorVal int) string {
	// TODO: test it
	// TODO: actually write out the PPM file
	// TODO: feels repetitive. factor into bigger for-loop or maybe anon func()?
	// TODO: feels like combining calculation and an action. split it?
	ppmHeader := fmt.Sprintf("\nP3\n%v %v\n%v\n", c.Width(), c.Height(), maxColorVal)
	maxColorValf := float64(maxColorVal)
	var b strings.Builder
	b.WriteString(ppmHeader)
	for _, row := range c.image {
		for _, pix := range row {
			R := int(math.Round(maxColorValf * pix.R)) // multiply by maxColorVal and round to nearest int
			b.WriteString(fmt.Sprint(R))               // put into builder
			b.WriteRune('\n')
		}
	}
	for _, row := range c.image {
		for _, pix := range row {
			G := int(math.Round(maxColorValf * pix.G))
			b.WriteString(fmt.Sprint(G))
			b.WriteRune('\n')
		}
	}
	for _, row := range c.image {
		for _, pix := range row {
			B := int(math.Round(maxColorValf * pix.B)) // multiply by maxColorVal and round to nearest int
			b.WriteString(fmt.Sprint(B))               // put into builder
			b.WriteRune('\n')
		}
	}
	b.WriteRune('\n')
	return b.String()
}
