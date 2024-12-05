package tracer

import (
	"math/rand"
	"testing"
)

// Creating and querying a ray
func TestNewRay(t *testing.T) {
	origin := NewPointTuple(1, 2, 3)
	direction := NewVectorTuple(4, 5, 6)
	r := NewRay(origin, direction)
	if origin != r.Origin {
		t.Errorf("ray origin not initialized properly")
	}
	if direction != r.Direction {
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
			t.Errorf("%v at time %v should have transformed to %v but was %v",
				row.ray, row.time, row.expect, result)
		}
	}
}

func TestSphereTable(t *testing.T) {
	vec := NewVectorTuple(0, 0, 1)
	sphere := NewSphere()
	data := []struct {
		r      Ray
		s      Sphere
		expect []float64
	}{
		// A ray intersects a sphere at two points
		{NewRay(NewPointTuple(0, 0, -5), vec), sphere, []float64{4, 6}},
		// A ray intersects a sphere at a tangent
		{NewRay(NewPointTuple(0, 1, -5), vec), sphere, []float64{5, 5}},
		// A ray misses a sphere
		{NewRay(NewPointTuple(0, 2, -5), vec), sphere, []float64{}},
		// A ray originates inside a sphere
		{NewRay(NewPointTuple(0, 0, 0), vec), sphere, []float64{-1, 1}},
		// A sphere is behind a ray
		{NewRay(NewPointTuple(0, 0, 5), vec), sphere, []float64{-6, -4}},
	}
	for _, row := range data {
		result := row.s.GetIntersects(row.r)
		if len(result) != len(row.expect) {
			t.Errorf("sphere %v intersecting ray %v should have %v intersects but had %v instead.\nexpect: %v\nresult: %v",
				row.s, row.r, len(row.expect), len(result), row.expect, result)
		}
		if len(result) != 0 {
			if abs(result[0].time-row.expect[0]) > 0 {
				t.Errorf("sphere %v intersecting ray %v should have 1st intersect at %v but was at %v instead",
					row.s, row.r, row.expect[0], result[0])
			}
			if abs(result[1].time-row.expect[1]) > 0 {
				t.Errorf("sphere %v intersecting ray %v should have 2nd intersect at %v but was at %v instead",
					row.s, row.r, row.expect[1], result[1])
			}
		}
	}
}

func TestNewIntersection(t *testing.T) {
	// An intersection encapsulates t and object
	s := NewSphere()
	time := 3.5
	i := NewIntersect(s, time)
	if i.time != time {
		t.Errorf("expected intersect time to be %v but was %v", time, i.time)
	}
}

func TestIntersectionSlice(t *testing.T) {
	// Aggregating intersections
	s := NewSphere()
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

func TestIntersectSetsObject(t *testing.T) {
	r := NewRay(NewPointTuple(0, 0, -5), NewVectorTuple(0, 0, 1))
	s1 := NewSphere()
	s2 := NewSphere()
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
	s := NewSphere()
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
		{"all negative t", NewIntersects(s, -1, -2), NewIntersect(NewSphere(), 200*rand.Float64()-100), false},
		// The hit is always the lowest nonnegative intersection
		{"sorting many t", NewIntersects(s, 5, 7, -3, 2), NewIntersect(s, 2), true},
	}
	for _, row := range data {
		result, ok := Hit(row.xs)
		if ok != row.ok {
			t.Errorf("%v: expected ok to be %v but was %v instead. row.expect test should also now fail.", row.name, row.expect, result)
		}
		if row.ok && (row.expect.id != result.id) {
			t.Errorf("%v: expected hit to be %v but was %v instead",
				row.name, row.expect, result)
		}
	}
}

func TestRayTransform(t *testing.T) {
	r := NewRay(NewPointTuple(1, 2, 3), NewVectorTuple(0, 1, 0))
	data := []struct {
		r      Ray
		m      Mat[Size4]
		expect Ray
	}{
		{r, I4.Translate(3, 4, 5), NewRay(NewPointTuple(4, 6, 8), NewVectorTuple(0, 1, 0))},
		{r, I4.Scale(2, 3, 4), NewRay(NewPointTuple(2, 6, 12), NewVectorTuple(0, 3, 0))},
	}
	for _, row := range data {
		result := row.r.Transform(row.m)
		if !row.expect.Origin.Equals(result.Origin) {

			t.Errorf("expected ray origin to be %v but was %v instead", row.expect.Origin, result.Origin)
		}
		if !row.expect.Direction.Equals(result.Direction) {

			t.Errorf("expected ray direction to be %v but was %v instead", row.expect.Direction, result.Direction)
		}
	}
}
