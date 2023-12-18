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

var idGen = rand.New(rand.NewSource(69))

type Sphere struct {
	id        int
	transform Mat4
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
	s := Sphere{idGen.Int(), I4}
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
	return []Intersect{
		{t1, s},
		{t2, s},
	}
}

func Hit(xs []Intersect) (Intersect, bool) {
	if len(xs) == 0 {
		return NewIntersect(0, NewSphere()), false
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
	return NewIntersect(0, NewSphere()), false
}

func (r Ray) Transform(m Mat4) Ray {
	return NewRay(m.TimesTuple(r.Origin), m.TimesTuple(r.Direction))
}
