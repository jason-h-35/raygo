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
		position: tracer.Point(0, 200, 0),
		velocity: tracer.Vector(0, 0, 0),
	}
	e := Environment{
		gravity: tracer.Vector(0, -1, 0),
		wind:    tracer.Vector(-0.01, 0, 0),
	}
	c := canvas.NewCanvas(900, 550)
	count := 0
	red, green, blue := canvas.NewColor(1, 0, 0), canvas.NewColor(0, 1, 0), canvas.NewColor(0, 0, 1)
	for count <= 10 {
		p = tick(e, p)
		c.WritePixel(int(p.position.X), int(p.position.Y), red)
		p = tick(e, p)
		c.WritePixel(int(p.position.X), int(p.position.Y), green)
		p = tick(e, p)
		c.WritePixel(int(p.position.X), int(p.position.Y), blue)
		count += 3
	}
	bytes, err := c.PPMFile(255, "/home/jason/out.ppm")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bytes)
}

func canvastest() {
	c := canvas.NewCanvas(1000, 1000)
	count := 0
	for count != 1000 {
		c.WritePixel(0, count, canvas.White)
		c.WritePixel(count, 0, canvas.White)
		c.WritePixel(count, count, canvas.White)
		count++
	}
	bytes, err := c.PPMFile(255, "/home/jason/out.ppm")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bytes)
}
