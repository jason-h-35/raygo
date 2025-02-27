package main

import (
	"fmt"
	"github.com/jason-h-35/raygo/tracer"
)

func main() {
	canvas := tracer.NewCanvas(100, 100)
	sphere := tracer.NewSphere(tracer.I4.Scale(50, 50, 1).Translate(50, 50, 0))
	bounds := canvas.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			canvas.Set(x, y, tracer.ColorGrey)
			r := tracer.NewRay(tracer.NewPoint(float64(x), float64(y), 0), tracer.NewVector(0, 0, -1))
			intersects := sphere.GetIntersects(r)
			for _, time := range tracer.GetIntersectTimes(intersects) {
				if time > 0 {
					canvas.Set(x, y, tracer.ColorRed)
					break
				}
			}
			fmt.Println(x, y)
		}
	}
	bytes, err := canvas.PPMFile(255, "/home/jason/out.ppm")
	canvas.PNGFile("/home/jason/out.png")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bytes)
}
