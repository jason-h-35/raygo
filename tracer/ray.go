package tracer

import (
	"math"
	"math/rand"
)

type Ray struct {
	Origin    Tuple
	Direction Tuple
}

var idGen = rand.New(rand.NewSource(69))

type Sphere struct {
	id int
}

type Intersect struct {
	time   float64
	object Sphere
}

func NewIntersect(time float64, object Sphere) Intersect {
	return Intersect{time, object}
}

func NewRay(origin, direction Tuple) Ray {
	return Ray{origin, direction}
}

func (r Ray) Position(time float64) Tuple {
	return r.Origin.Plus(r.Direction.Times(time))
}

func NewSphere() Sphere {
	s := Sphere{idGen.Int()}
	return s
}

func (s Sphere) GetIntersects(r Ray) []Intersect {
	sphereToRay := r.Origin.Minus(NewPointTuple(0, 0, 0))
	a := r.Direction.Dot(r.Direction)
	b := 2 * r.Direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return []Intersect{}
	}
	sqrtDiscriminant := math.Sqrt(discriminant)
	t1 := (-b - sqrtDiscriminant) / (2 * a)
	t2 := (-b + sqrtDiscriminant) / (2 * a)
	return []Intersect{
		{t1, s},
		{t2, s},
	}
}

func Hit(xs []Intersect) (Intersect, bool) {
	if len(xs) == 0 {
	}
	return xs[0], true
}
