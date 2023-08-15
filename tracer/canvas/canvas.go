package canvas

type Canvas struct {
	data [][]Color
}

func NewCanvas(width, height int) Canvas {
	c := Canvas{}
	return c
}

func (c *Canvas) WritePixel(x int, y int, color Color) {
	canvas := *c
	canvas.data[x][y] = color
}

func (c *Canvas) ReadPixel(x, y int) Color {
	canvas := *c
	return canvas.data[x][y]
}
