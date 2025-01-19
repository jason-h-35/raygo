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
		c.Set(int(p.position.X), int(p.position.Y), tracer.ColorRed)
		p = tick(e, p)
		c.Set(int(p.position.X), int(p.position.Y), tracer.ColorGreen)
		p = tick(e, p)
		c.Set(int(p.position.X), int(p.position.Y), tracer.ColorBlue)
		p = tick(e, p)
		c.Set(int(p.position.X), int(p.position.Y), tracer.ColorYellow)
		p = tick(e, p)
		c.Set(int(p.position.X), int(p.position.Y), tracer.ColorMagenta)
		p = tick(e, p)
		c.Set(int(p.position.X), int(p.position.Y), tracer.ColorCyan)
	}
	bytes, err := c.PPMFile(255, "/home/jason/out.ppm")
	c.PNGFile("/home/jason/out.png")
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
		c.Set(0, count, tracer.ColorMagenta)
		c.Set(count, 0, tracer.ColorGreen)
		c.Set(count, count, tracer.ColorBlue)
		count++
	}
	// fmt.Println(c.Image)
	fmt.Println(c.PPMStr(1))
	bytes, err := c.PPMFile(1, "/home/jason/out.ppm")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bytes)
}
