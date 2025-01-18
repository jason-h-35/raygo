package main

import (
	"fmt"
	"github.com/jason-h-35/raygo/tracer"
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
	return Projectile{
		position: proj.position.Plus(proj.velocity.Divide(5)),
		velocity: proj.velocity.Plus(env.gravity.Plus(env.wind).Divide(5)),
	}
}

func main() {
	p := Projectile{
		position: tracer.NewPoint(0, 1, 0),
		velocity: tracer.NewVector(1, 1.8, 0).Normalized().Times(11.25),
	}
	e := Environment{
		gravity: tracer.NewVector(0, -0.2, 0),
		wind:    tracer.NewVector(-0.01, 0, 0),
	}
	c := tracer.NewCanvas(900, 450)
	for i := 0; i <= 1500; i += 3 {
		p = tick(e, p)
		c.WritePixel(int(p.position.X), int(p.position.Y), tracer.Red)
		p = tick(e, p)
		c.WritePixel(int(p.position.X), int(p.position.Y), tracer.Green)
		p = tick(e, p)
		c.WritePixel(int(p.position.X), int(p.position.Y), tracer.Blue)
		p = tick(e, p)
		c.WritePixel(int(p.position.X), int(p.position.Y), tracer.Yellow)
		p = tick(e, p)
		c.WritePixel(int(p.position.X), int(p.position.Y), tracer.Magenta)
		p = tick(e, p)
		c.WritePixel(int(p.position.X), int(p.position.Y), tracer.Cyan)
	}
	bytes, err := c.PPMFile(255, "/home/jason/out.ppm")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bytes)
}

// STILL BUGGED ON NON-SQUARE CANVAS, EVEN WITH BOOK TESTS
func main2() {
	c := tracer.NewCanvas(100, 100)
	count := 0
	for count != 4 {
		c.WritePixel(0, count, tracer.Magenta)
		c.WritePixel(count, 0, tracer.Green)
		c.WritePixel(count, count, tracer.Blue)
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
