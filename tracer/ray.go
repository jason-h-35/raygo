package tracer

type Ray struct {
	origin    Tuple
	direction Tuple
}

var id int = 0

type Sphere struct {
	id int
}

func NewRay(origin, direction Tuple) Ray {
	return Ray{origin, direction}
}

func (r Ray) Position(time float64) Tuple {
	return r.origin.Plus(r.direction.Times(time))
}

func NewSphere() Sphere {
	s := Sphere{id}
	id += 1
	return s
}
