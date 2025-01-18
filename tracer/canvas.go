package tracer

import (
	"fmt"
	"image"
	"image/color"
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
	// TODO: test it because it's broken!!!
	bounds := c.Bounds().Max
	width, height := bounds.X, bounds.Y
	ppmHeader := fmt.Sprintf("P3\n%v %v\n%v\n", width, height, maxColorVal)
	lenCount, lenMax := 0, 67
	var b strings.Builder
	b.WriteString(ppmHeader)
	// transform Canvas of Colors into 1-D arrays of ints representing just one Color Value from 0 to maxColorVal
	for _, row := range c.image {
		for _, pix := range row {
			R, G, B := pix.Times(maxColorVal).AsInts()
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
