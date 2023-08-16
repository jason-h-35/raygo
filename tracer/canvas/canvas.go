package canvas

import (
	"fmt"
	"math"
	"os"
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

func (c *Canvas) PPMStr(maxColorVal int) string {
	// TODO: test it
	ppmHeader := fmt.Sprintf("\nP3\n%v %v\n%v\n", c.Width(), c.Height(), maxColorVal)
	maxColorValf := float64(maxColorVal)
	colorFunc := func(pixcolor float64) int {
		return int(math.Round(maxColorValf * pixcolor))
	}
	var ppmData map[rune][]int
	// transform Canvas of Colors into 1-D arrays of ints representing just one Color Value from 0 to maxColorVal
	for _, row := range c.image {
		for _, pix := range row {
			ppmData['R'] = append(ppmData['R'], colorFunc(pix.R))
			ppmData['G'] = append(ppmData['G'], colorFunc(pix.G))
			ppmData['B'] = append(ppmData['B'], colorFunc(pix.B))
		}
	}
	// string join into the string to be written for each of R,G,B
	var ppmString map[rune]string
	for k := range ppmData {
		ppmString[k] = strings.Trim(fmt.Sprint(ppmData[k]), "[]") + "\n"
	}
	// concat it all and return
	var b strings.Builder
	b.WriteString(ppmHeader)
	b.WriteString(ppmString['R'])
	b.WriteString(ppmString['G'])
	b.WriteString(ppmString['B'])
	return b.String()
}

func (c *Canvas) PPMFile(maxColorVal int, writePath string) (int, error) {
	ppmStr := c.PPMStr(maxColorVal)
	file, fileErr := os.Create("writepath")
	if fileErr != nil {
		fmt.Println(fileErr)
		return 0, fileErr
	}
	return fmt.Fprintf(file, "%v", ppmStr)
}
