package canvas

type Canvas struct {
	image [][]Color
}

func NewCanvas(width, height int) Canvas {
	image := make([][]Color, width)
	for i := range image {
		image[i] = make([]Color, height)
	}
	return Canvas{image}
}

func (c *Canvas) WritePixel(x int, y int, color Color) {
	canvas := *c
	canvas.image[x][y] = color
}

func (c *Canvas) ReadPixel(x, y int) Color {
	canvas := *c
	return canvas.image[x][y]
}
