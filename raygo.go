package main

import (
	"fmt"
	// "math"

	"github.com/jason-h-35/raygo/tracer"
	"github.com/jason-h-35/raygo/tracer/canvas"
)

type Projectile struct {
	position tracer.Tuple
	velocity tracer.Tuple
}

type Environment struct {
	gravity tracer.Tuple
	wind    tracer.Tuple
}

func tick(env Environment, proj Projectile) Projectile {
	position := proj.position.Plus(proj.velocity.Divide(5))
	velocity := proj.velocity.Plus(env.gravity.Plus(env.wind).Divide(5))
	return Projectile{position, velocity}
}

func main() {
	p := Projectile{
		position: tracer.NewPointTuple(0, 1, 0),
		velocity: tracer.NewVectorTuple(1, 1.8, 0).Normalized().Times(11.25),
	}
	e := Environment{
		gravity: tracer.NewVectorTuple(0, -0.2, 0),
		wind:    tracer.NewVectorTuple(-0.01, 0, 0),
	}
	c := canvas.NewCanvas(900, 450)
	count := 0
	red, green, blue := canvas.NewColor(1, 0, 0), canvas.NewColor(0, 1, 0), canvas.NewColor(0, 0, 1)
	yellow, magenta, cyan := canvas.NewColor(1, 1, 0), canvas.NewColor(1, 0, 1), canvas.NewColor(0, 1, 1)
	for count <= 1500 {
		p = tick(e, p)
		c.WritePixel(int(p.position.X), int(p.position.Y), red)
		p = tick(e, p)
		c.WritePixel(int(p.position.X), int(p.position.Y), green)
		p = tick(e, p)
		c.WritePixel(int(p.position.X), int(p.position.Y), blue)
		p = tick(e, p)
		c.WritePixel(int(p.position.X), int(p.position.Y), yellow)
		p = tick(e, p)
		c.WritePixel(int(p.position.X), int(p.position.Y), magenta)
		p = tick(e, p)
		c.WritePixel(int(p.position.X), int(p.position.Y), cyan)
		count += 3
	}
	bytes, err := c.PPMFile(255, "/home/jason/out.ppm")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bytes)
}

// STILL BUGGED ON NON-SQUARE CANVAS, EVEN WITH BOOK TESTS
func main2() {
	c := canvas.NewCanvas(10, 5)
	count := 0
	for count != 4 {
		c.WritePixel(0, count, canvas.NewColor(1, 0, 1))
		c.WritePixel(count, 0, canvas.NewColor(0, 1, 0))
		c.WritePixel(count, count, canvas.NewColor(0, 0, 1))
		count++
	}
	fmt.Println(c.Image)
	fmt.Println(c.PPMStr(1))
	bytes, err := c.PPMFile(1, "/home/jason/out.ppm")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bytes)
}
