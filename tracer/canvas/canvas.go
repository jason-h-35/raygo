package canvas

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Canvas struct {
	Image [][]Color
}

func NewCanvas(xMax, yMax int) Canvas {
	image := make([][]Color, xMax)
	for i := range image {
		image[i] = make([]Color, yMax)
	}
	return Canvas{image}
}

func (c *Canvas) Width() int {
	return len(c.Image)
}

func (c *Canvas) Height() int {
	return len(c.Image[0]) // how to check if all sub-arrays are the same length?
}

func (c *Canvas) WritePixel(x int, y int, color Color) error {
	if !(0 <= x && x < c.Width()) {
		return errors.New("x out of range")
	}
	if !(0 <= y && y < c.Height()) {
		return errors.New("y out of range")
	}
	c.Image[x][y] = color
	return nil
}

func (c *Canvas) ReadPixel(x, y int) Color {
	return c.Image[x][y]
}

func (c Color) ToPPMRange(maximum int) Color {
	// TODO: handle values outside of 0.0 and 1.0 color min and max
	c = c.Times(float64(maximum))
	return NewColor(
		math.Round(c.R),
		math.Round(c.G),
		math.Round(c.B),
	)
}

func (c Color) asInts() (int, int, int) {
	return int(c.R), int(c.G), int(c.B)
}

func (c *Canvas) PPMStr(maxColorVal int) string {
	// TODO: test it because it's broken!!!
	width, height := c.Width(), c.Height()
	ppmHeader := fmt.Sprintf("P3\n%v %v\n%v\n", width, height, maxColorVal)
	lenCount, lenMax := 0, 67
	var b strings.Builder
	b.WriteString(ppmHeader)
	// transform Canvas of Colors into 1-D arrays of ints representing just one Color Value from 0 to maxColorVal
	for _, row := range c.Image {
		for _, pix := range row {
			R, G, B := pix.ToPPMRange(maxColorVal).asInts()
			Rs, Gs, Bs := strconv.Itoa(R), strconv.Itoa(G), strconv.Itoa(B)
			b.WriteString(Rs)
			b.WriteRune(' ')
			b.WriteString(Gs)
			b.WriteRune(' ')
			b.WriteString(Bs)
			b.WriteRune(' ')
			lenCount += len(Rs) + len(Gs) + len(Bs) + 3
			if lenCount > lenMax {
				b.WriteRune('\n')
				lenCount = 0
			}
		}
	}
	b.WriteRune('\n')
	return b.String()
}

func (c *Canvas) PPMFile(maxColorVal int, writePath string) (int, error) {
	ppmStr := c.PPMStr(maxColorVal)
	file, fileErr := os.Create(writePath)
	defer file.Close()
	if fileErr != nil {
		fmt.Println(fileErr)
		return 0, fileErr
	}
	return fmt.Fprintf(file, "%v", ppmStr)
}
