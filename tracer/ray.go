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

type Intersection struct {
	time   float64
	object Sphere
}

func NewIntersection(time float64, object Sphere) Intersection {
	return Intersection{time, object}
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

func (s Sphere) Intersect(r Ray) []Intersection {
	sphereToRay := r.Origin.Minus(NewPointTuple(0, 0, 0))
	a := r.Direction.Dot(r.Direction)
	b := 2 * r.Direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return []Intersection{}
	}
	sqrtDiscriminant := math.Sqrt(discriminant)
	t1 := (-b - sqrtDiscriminant) / (2 * a)
	t2 := (-b + sqrtDiscriminant) / (2 * a)
	return []Intersection{
		{t1, s},
		{t2, s},
	}
}

func Hit(xs []Intersection) (Intersection, bool) {
	nonneg_t := xs[:0]
	for _, x := range xs {
		if x.time > 0
	}
	return xs[0], true
}
