package tracer

import (
	"math"
	"math/rand"
	"slices"
)

// Ray represents a ray in 3D space with an origin point and a direction vector.
type Ray struct {
	origin   Tuple // The starting point of the ray
	velocity Tuple // The direction and magnitude the ray travels
}

// Random number generators for creating unique IDs
var (
	sphereIDGen       = rand.New(rand.NewSource(69))
	intersectionIDGen = rand.New(rand.NewSource(420))
)

// Sphere represents a sphere in 3D space.
type Sphere struct {
	id        int        // Unique identifier for the sphere
	transform Mat[Size4] // Transformation matrix for the sphere
}

// Intersect represents the intersection of a ray with an object.
type Intersect struct {
	id     int     // Unique identifier for the intersection
	object Sphere  // The object that was intersected
	time   float64 // The time/distance along the ray where intersection occurred
}

// NewIntersect creates a new intersection with a unique ID.
func NewIntersect(object Sphere, time float64) Intersect {
	return Intersect{intersectionIDGen.Int(), object, time}
}

// NewIntersects creates multiple intersections for a single object.
// Useful for objects like spheres that can have multiple intersection points.
func NewIntersects(object Sphere, times ...float64) []Intersect {
	xs := make([]Intersect, len(times))
	for i, t := range times {
		xs[i] = NewIntersect(object, t)
	}
	return xs
}

func (ix Intersect) GetObject() Sphere {
	return ix.object
}

func (ix Intersect) GetTime() float64 {
	return ix.time
}

// NewRay creates a new ray with the given origin point and direction vector.
func NewRay(origin, direction Tuple) Ray {
	return Ray{origin, direction}
}

// Position calculates a point along the ray at the given time/distance.
func (r Ray) Position(time float64) Tuple {
	return r.origin.Plus(r.velocity.Times(time))
}

// NewSphere creates a new sphere with a unique ID and identity transformation.
func NewSphere(transform Mat[Size4]) Sphere {
	s := Sphere{sphereIDGen.Int(), transform}
	return s
}

// GetIntersects calculates all intersections between a ray and a sphere.
// Returns an empty slice if no intersections exist.
// For a sphere centered at the origin with radius 1, uses the quadratic formula:
// t^2(d⋅d) + 2t(d⋅(o-c)) + (o-c)⋅(o-c) - r^2 = 0
func (s Sphere) GetIntersects(r2 Ray) []Intersect {
	r := r2.Transform(s.transform.Inverse())
	sphereToRay := r.origin.Minus(NewPoint(0, 0, 0))
	a := r.velocity.Dot(r.velocity)
	b := 2 * r.velocity.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1 // radius is 1
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

func GetIntersectTimes(intersects []Intersect) []float64 {
	times := make([]float64, len(intersects), len(intersects))
	for i := 0; i < len(intersects); i++ {
		times[i] = intersects[i].time
	}
	return times
}

// Equals checks if two intersections are exactly the same,
// comparing IDs, object IDs, and intersection times.
func (x Intersect) Equals(y Intersect) bool {
	return x.id == y.id &&
		x.object.id == y.object.id &&
		math.Abs(x.time-y.time) < epsilon
}

// Same checks if two intersections represent the same physical intersection point,
// ignoring the intersection ID but comparing object IDs and times.
func (x Intersect) Same(y Intersect) bool {
	return x.object.id == y.object.id &&
		math.Abs(x.time-y.time) < epsilon
}

// compareIntersectTime compares two intersections by their time values.
// Returns:
//
//	-1 if a occurs before b
//	 0 if times are approximately equal
//	 1 if a occurs after b
func compareIntersectTime(a, b Intersect) int {
	if math.Abs(a.time-b.time) < epsilon {
		return 0
	} else if a.time < b.time {
		return -1
	} else {
		return 1
	}
}

// Hit finds the first non-negative intersection from a slice of intersections.
// Returns the intersection and true if found, or a zero intersection and false if not found.
// Ensures intersections are sorted by time before searching.
func Hit(xs []Intersect) (Intersect, bool) {
	if len(xs) == 0 {
		return NewIntersect(NewSphere(I4), 0), false
	}
	// TODO: Eventually it won't be viable to sort for every Hit calculation
	if !slices.IsSortedFunc(xs, compareIntersectTime) {
		slices.SortFunc(xs, compareIntersectTime)
	}
	for _, x := range xs {
		if x.time >= 0 {
			return x, true
		}
	}
	return NewIntersect(NewSphere(I4), 0), false
}

// Transform applies a transformation matrix to both the origin and direction of a ray.
func (r Ray) Transform(m Mat[Size4]) Ray {
	return NewRay(m.TimesTuple(r.origin), m.TimesTuple(r.velocity))
}
