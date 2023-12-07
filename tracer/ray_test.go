package tracer

import (
	"testing"
)

// Creating and querying a ray
func TestNewRay(t *testing.T) {
	origin := NewPointTuple(1, 2, 3)
	direction := NewVectorTuple(4, 5, 6)
	r := NewRay(origin, direction)
	if origin != r.origin {
		t.Errorf("ray origin not initialized properly")
	}
	if direction != r.direction {
		t.Errorf("ray origin not initialized properly")
	}
}

// Computing a point from a distance
func TestPosition(t *testing.T) {
	r := NewRay(
		NewPointTuple(2, 3, 4),
		NewVectorTuple(1, 0, 0),
	)
	data := []struct {
		ray    Ray
		time   float64
		expect Tuple
	}{
		{r, 0, NewPointTuple(2, 3, 4)},
		{r, 1, NewPointTuple(3, 3, 4)},
		{r, -1, NewPointTuple(1, 3, 4)},
		{r, 2.5, NewPointTuple(4.5, 3, 4)},
	}
	for _, row := range data {
		result := row.ray.Position(row.time)
		if !result.Equals(row.expect) {
			t.Errorf("%v at time %v should have transformed to %v but was %v", row.ray, row.time, row.expect, result)
		}
	}
}

func TestSphereTable(t *testing.T) {
	data := []struct {
		r         Ray
		s         Sphere
		lenExpect int
		xsExpect  []float64
	}{
		// A ray intersects a sphere at two points
		{NewRay(NewPointTuple(0, 0, -5), NewVectorTuple(0, 0, 1)), NewSphere(), 2, []float64{4, 6}},
		// A ray intersects a sphere at a tangent
		{NewRay(NewPointTuple(0, 1, -5), NewVectorTuple(0, 0, 1)), NewSphere(), 2, []float64{5, 5}},
		// A ray misses a sphere
		{NewRay(NewPointTuple(0, 2, -5), NewVectorTuple(0, 0, 1)), NewSphere(), 0, []float64{}},
		// A ray originates inside a sphere
		{NewRay(NewPointTuple(0, 0, 0), NewVectorTuple(0, 0, 1)), NewSphere(), 2, []float64{-1, 1}},
		// A sphere is behind a ray
		{NewRay(NewPointTuple(0, 0, 5), NewVectorTuple(0, 0, 1)), NewSphere(), 2, []float64{-6, -4}},
	}
}

// A ray misses a sphere
func TestSphereMiss(t *testing.T) {
	r := NewRay(
		NewPointTuple(0, 2, -5),
		NewVectorTuple(0, 0, 1),
	)
	s := NewSphere()
	xs := s.Intersect(r)
	if len(xs) != 0 {
		t.Error("incorrect amount of intersections. expected 0. got %v", len(xs))
	}
}
