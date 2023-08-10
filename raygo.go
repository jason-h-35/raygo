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
	position := proj.position.Plus(proj.velocity)
	velocity := proj.velocity.Plus(env.gravity).Plus(env.wind)
	return Projectile{position, velocity}
}

func main() {
	p := Projectile{
		position: tracer.Point(0, 1, 0),
		velocity: tracer.Vector(1, 1, 0).Normalized(),
	}
	e := Environment{
		gravity: tracer.Vector(0, -0.1, 0),
		wind:    tracer.Vector(-0.01, 0, 0),
	}
	for p.position.Y <= 0 {
		p = tick(e, p)
		fmt.Printf("%v", p.position)
	}
}
