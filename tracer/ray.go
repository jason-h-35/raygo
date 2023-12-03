package tracer

type Ray struct {
	origin    Tuple
	direction Tuple
}

func NewRay(origin, direction Tuple) Ray {
	return Ray{origin, direction}
}

func (r Ray) Position(time float64) Tuple {
	return r.origin.Plus(r.direction.Times(time))
}
