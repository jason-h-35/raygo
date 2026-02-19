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
			canvas.SetLinear(x, y, tracer.ColorGrey)
			r := tracer.NewRay(tracer.NewPoint(float64(x), float64(y), 0), tracer.NewVector(0, 0, -1))
			intersects := sphere.GetIntersects(r)
			for _, time := range tracer.GetIntersectTimes(intersects) {
				if time > 0 {
					canvas.SetLinear(x, y, tracer.ColorRed)
					break
				}
			}
		}
	}

	bytes, err := canvas.PPMFile(255, "out.ppm")
	if err != nil {
		fmt.Printf("failed to write PPM: %v\n", err)
	} else {
		fmt.Printf("wrote %d bytes to out.ppm\n", bytes)
	}

	if err := canvas.PNGFile("out.png"); err != nil {
		fmt.Printf("failed to write PNG: %v\n", err)
	}
	if err := canvas.JPEGFile("out.jpg", 95); err != nil {
		fmt.Printf("failed to write JPEG: %v\n", err)
	}
}
