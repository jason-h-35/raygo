package main

import "github.com/jason-h-35/raygo/tracer"

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

}
