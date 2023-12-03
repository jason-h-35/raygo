package tracer

type Ray struct {
	origin    Tuple
	direction Tuple
}

func NewRay(origin, direction Tuple) Ray {
	return Ray{origin, direction}
}
