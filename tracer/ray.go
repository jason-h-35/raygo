package tracer

import (
	"math"
	"math/rand"
	"slices"
)

type Ray struct {
	Origin    Tuple
	Direction Tuple
}

var (
	sphereIDGen       = rand.New(rand.NewSource(69))
	intersectionIDGen = rand.New(rand.NewSource(420))
)

type Sphere struct {
	id        int
	transform Mat[Size4]
}

type Intersect struct {
	id     int
	object Sphere
	time   float64
}

func NewIntersect(object Sphere, time float64) Intersect {
	return Intersect{intersectionIDGen.Int(), object, time}
}

func NewIntersects(object Sphere, times ...float64) []Intersect {
	xs := make([]Intersect, len(times))
	for i, t := range times {
		xs[i] = NewIntersect(object, t)
	}
	return xs
}

func NewRay(origin, direction Tuple) Ray {
	return Ray{origin, direction}
}

func (r Ray) Position(time float64) Tuple {
	return r.Origin.Plus(r.Direction.Times(time))
}

func NewSphere() Sphere {
	s := Sphere{sphereIDGen.Int(), I4}
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
	// not necessary, but trying to keep sorted t's
	if t1 > t2 {
		t1, t2 = t2, t1
	}
	return NewIntersects(s, t1, t2)
}

func Hit(xs []Intersect) (Intersect, bool) {
	if len(xs) == 0 {
		return NewIntersect(NewSphere(), 0), false
	}
	// needed for slice sort funcs to sort Intersects
	f := func(a, b Intersect) int {
		if abs(a.time-b.time) < eps {
			return 0
		} else if a.time < b.time {
			return -1
		} else {
			return 1
		}
	}
	if !slices.IsSortedFunc(xs, f) {
		slices.SortFunc(xs, f)
	}
	for _, x := range xs {
		if x.time >= 0 {
			return x, true
		}
	}
	return NewIntersect(NewSphere(), 0), false
}

func (r Ray) Transform(m Mat[Size4]) Ray {
	return NewRay(m.TimesTuple(r.Origin), m.TimesTuple(r.Direction))
}
