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
	position := proj.position.Plus(proj.velocity)
	velocity := proj.velocity.Plus(env.gravity).Plus(env.wind)
	return Projectile{position, velocity}
}

func blah() {
	p := Projectile{
		position: tracer.Point(0, 1, 0),
		velocity: tracer.Vector(1, 1.8, 0).Normalized().Times(11.25),
	}
	e := Environment{
		gravity: tracer.Vector(0, -1.0, 0),
		wind:    tracer.Vector(-0.01, 0, 0),
	}
	c := canvas.NewCanvas(900, 550)
	count := 0
	// for p.position.Y >= 0 {
	for count != 1000 {
		count += 1
		p = tick(e, p)
		// c.WritePixel(int(math.Round(p.position.X)), int(math.Round(p.position.Y)), canvas.White)
		c.WritePixel(0, count, canvas.White)
		c.WritePixel(count, 0, canvas.White)
		c.WritePixel(count, count, canvas.White)
		fmt.Printf("%v\n", p.position)
	}
	c.WritePixel(0, 0, canvas.White)
	fmt.Printf("%v\n", c.PPMStr(255))
	bytesWritten, err := c.PPMFile(255, "/home/jason/out.ppm")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bytesWritten)
	fmt.Println(count)
}

func main() {
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
