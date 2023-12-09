package tracer

import (
	"math"
)

type Ray struct {
	Origin    Tuple
	Direction Tuple
}

var id int = 0

type Sphere struct {
	id int
}

type Intersection struct {
	time   float64
	object Sphere
}

func NewRay(origin, direction Tuple) Ray {
	return Ray{origin, direction}
}

func (r Ray) Position(time float64) Tuple {
	return r.Origin.Plus(r.Direction.Times(time))
}

func NewSphere() Sphere {
	s := Sphere{id}
	id += 1
	return s
}

func (s Sphere) Intersect(r Ray) [2]float64 {
	sphereToRay := r.Origin.Minus(NewPointTuple(0, 0, 0))
	a := r.Direction.Dot(r.Direction)
	b := 2 * r.Direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return [2]float64{}
	}
	sqrt := math.Sqrt(discriminant)
	t1 := (-b - sqrt) / (2 * a)
	t2 := (-b + sqrt) / (2 * a)
	return [2]float64{t1, t2}
}
