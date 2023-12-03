package tracer

import (
	"testing"
)

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
