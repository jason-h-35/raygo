package canvas

import (
	"errors"
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

func (c *Canvas) WritePixel(x int, y int, color Color) error {
	if !(0 <= x && x < c.Width()) {
		return errors.New("x out of range")
	}
	if !(0 <= y && y < c.Height()) {
		return errors.New("y out of range")
	}
	c.image[x][y] = color
	return nil
}

func (c *Canvas) ReadPixel(x, y int) Color {
	return c.image[x][y]
}

func ppmRange(intensity float64, maximum int) int {
	return int(math.Round(float64(maximum) * intensity))
}

func (c *Canvas) PPMStr(maxColorVal int) string {
	// TODO: test it because it's broken!!!
	width, height := c.Width(), c.Height()
	ppmHeader := fmt.Sprintf("\nP3\n%v %v\n%v\n", width, height, maxColorVal)
	ppmData := map[rune][]int{}
	// transform Canvas of Colors into 1-D arrays of ints representing just one Color Value from 0 to maxColorVal
	for _, row := range c.image {
		for _, pix := range row {
			ppmData['R'] = append(ppmData['R'], ppmRange(pix.R, maxColorVal))
			ppmData['G'] = append(ppmData['G'], ppmRange(pix.G, maxColorVal))
			ppmData['B'] = append(ppmData['B'], ppmRange(pix.B, maxColorVal))
		}
	}
	// string join into the string to be written for each of R,G,B
	ppmString := map[rune]string{}
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
	file, fileErr := os.Create(writePath)
	defer file.Close()
	if fileErr != nil {
		fmt.Println(fileErr)
		return 0, fileErr
	}
	return fmt.Fprintf(file, "%v", ppmStr)
}
