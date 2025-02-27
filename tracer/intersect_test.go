package tracer

import (
	"math"
	"math/rand"
	"testing"
)

// Creating and querying a ray
func TestNewRay(t *testing.T) {
	origin := NewPoint(1, 2, 3)
	direction := NewVector(4, 5, 6)
	r := NewRay(origin, direction)
	if origin != r.origin {
		t.Errorf("ray origin not initialized properly")
	}
	if direction != r.velocity {
		t.Errorf("ray origin not initialized properly")
	}
}

// Creating and querying a sphere
func TestNewSphere(t *testing.T) {
	s := NewSphere(I4)
	// A sphere's default transformation
	if !s.transform.Equals(I4) {
		t.Errorf("sphere transform not initialized properly")
	}
	transform := I4.Translate(2, 3, 4)
	s.transform = transform
	// Changing a sphere's transformation
	if !s.transform.Equals(transform) {
		t.Errorf("sphere transform property not setting properly")
	}
}

// Computing a point from a distance
func TestPosition(t *testing.T) {
	r := NewRay(
		NewPoint(2, 3, 4),
		NewVector(1, 0, 0),
	)
	data := []struct {
		ray    Ray
		time   float64
		expect Tuple
	}{
		{r, 0, NewPoint(2, 3, 4)},
		{r, 1, NewPoint(3, 3, 4)},
		{r, -1, NewPoint(1, 3, 4)},
		{r, 2.5, NewPoint(4.5, 3, 4)},
	}
	for _, row := range data {
		result := row.ray.Position(row.time)
		if !result.Equals(row.expect) {
			t.Errorf("%v at time %v should have transformed to %v but was %v",
				row.ray, row.time, row.expect, result)
		}
	}
}

// Intersecting Rays with Spheres
func TestSphereTable(t *testing.T) {
	vec := NewVector(0, 0, 1)
	data := []struct {
		r      Ray
		s      Sphere
		expect []float64
	}{
		// A ray intersects a sphere at two points
		{NewRay(NewPoint(0, 0, -5), vec), NewSphere(I4), []float64{4, 6}},
		// A ray intersects a sphere at a tangent
		{NewRay(NewPoint(0, 1, -5), vec), NewSphere(I4), []float64{5, 5}},
		// A ray misses a sphere
		{NewRay(NewPoint(0, 2, -5), vec), NewSphere(I4), []float64{}},
		// A ray originates inside a sphere
		{NewRay(NewPoint(0, 0, 0), vec), NewSphere(I4), []float64{-1, 1}},
		// A sphere is behind a ray
		{NewRay(NewPoint(0, 0, 5), vec), NewSphere(I4), []float64{-6, -4}},
		// Intersecting a scaled sphere with a ray
		{NewRay(NewPoint(0, 0, -5), vec), NewSphere(I4.Scale(2, 2, 2)), []float64{3, 7}},
		// Intersecting a translated sphere with a ray
		{NewRay(NewPoint(0, 0, -5), vec), NewSphere(I4.Translate(5, 0, 0)), []float64{}},
	}
	for _, row := range data {
		result := row.s.GetIntersects(row.r)
		if len(result) != len(row.expect) {
			t.Errorf("sphere %v intersecting ray %v should have %v intersects but had %v instead.\nexpect: %v\nresult: %v",
				row.s, row.r, len(row.expect), len(result), row.expect, result)
		}
		if len(result) != 0 {
			if math.Abs(result[0].time-row.expect[0]) > 0 {
				t.Errorf("sphere %v intersecting ray %v should have 1st intersect at %v but was at %v instead",
					row.s, row.r, row.expect[0], result[0])
			}
			if math.Abs(result[1].time-row.expect[1]) > 0 {
				t.Errorf("sphere %v intersecting ray %v should have 2nd intersect at %v but was at %v instead",
					row.s, row.r, row.expect[1], result[1])
			}
		}
	}
}

// An Intersection encapsulates t and object
func TestNewIntersection(t *testing.T) {
	s := NewSphere(I4)
	time := 3.5
	i := NewIntersect(s, time)
	if s.id != i.object.id {
		t.Errorf("expected object id to be %v but was %v", s.id, i.object.id)
	}
	if time != i.time {
		t.Errorf("expected intersect time to be %v but was %v", time, i.time)
	}
}

func TestNewIntersects(t *testing.T) {
	// Aggregating intersections
	s := NewSphere(I4)
	xs := NewIntersects(s, 1, 2)
	if len(xs) != 2 {
		t.Errorf("incorrect intersection slice len. expected 2, got %v", len(xs))
	}
	if xs[0].time != 1 {
		t.Errorf("incorrect time value. expected 1, got %v", xs[0].time)
	}
	if xs[1].time != 2 {
		t.Errorf("incorrect time value. expected 2, got %v", xs[1].time)
	}
}

// Intersect sets the object on the intersection
func TestIntersectSetsObject(t *testing.T) {
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	s1 := NewSphere(I4)
	s2 := NewSphere(I4)
	xs := s1.GetIntersects(r)
	if len(xs) != 2 {
		t.Errorf("incorrect len. got %v", len(xs))
	}
	for _, ix := range xs {
		if ix.object.id != s1.id {
			t.Errorf("intersection object not set. expected %v. got %v", s1, ix.object)
		}
		if ix.object.id == s2.id {
			t.Errorf("ray intersecting with a different sphere. are object IDs working?. expected %v. got %v", s1, ix.object)
		}
	}
}

func TestHitTable(t *testing.T) {
	s := NewSphere(I4)
	data := []struct {
		name   string
		xs     []Intersect
		expect Intersect
		ok     bool
	}{
		// The hit, when all intersections have positive t
		{"all positive t", NewIntersects(s, 2, 1), NewIntersect(s, 1), true},
		// The hit, when some intersections have negative t
		{"some negative t", NewIntersects(s, 1, -1), NewIntersect(s, 1), true},
		// The hit, when all intersections have negative t. ok is checked and expect is not read
		{"all negative t", NewIntersects(s, -1, -2), NewIntersect(NewSphere(I4), 200*rand.Float64()-100), false},
		// The hit is always the lowest nonnegative intersection
		{"sorting many t", NewIntersects(s, 5, 7, -3, 2), NewIntersect(s, 2), true},
	}
	for _, row := range data {
		result, ok := Hit(row.xs)
		if ok != row.ok {
			t.Errorf("%v: expected ok to be %v but was %v instead. row.expect test should also now fail.", row.name, row.expect, result)
		}
		if row.ok && (!row.expect.Same(result)) {
			t.Errorf("%v: expected hit to be %v but was %v instead",
				row.name, row.expect, result)
		}
	}
}

func TestRayTransform(t *testing.T) {
	r := NewRay(NewPoint(1, 2, 3), NewVector(0, 1, 0))
	data := []struct {
		r      Ray
		m      Mat[Size4]
		expect Ray
	}{
		// Translating a ray
		{r, I4.Translate(3, 4, 5), NewRay(NewPoint(4, 6, 8), NewVector(0, 1, 0))},
		// Scaling a ray
		{r, I4.Scale(2, 3, 4), NewRay(NewPoint(2, 6, 12), NewVector(0, 3, 0))},
	}
	for _, row := range data {
		result := row.r.Transform(row.m)
		if !row.expect.origin.Equals(result.origin) {

			t.Errorf("expected ray origin to be %v but was %v instead", row.expect.origin, result.origin)
		}
		if !row.expect.velocity.Equals(result.velocity) {

			t.Errorf("expected ray direction to be %v but was %v instead", row.expect.velocity, result.velocity)
		}
	}
}
