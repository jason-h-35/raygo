package tracer

import (
	"math"
	"testing"
)

func TestTransform(t *testing.T) {
	data := []struct {
		transform Mat4
		pt        Tuple
		expect    Tuple
	}{
		{I4.Translate(5, -3, 2), Point(-3, 4, 5), Point(2, 1, 7)},
		{I4.Translate(5, -3, 2).Inverse(), Point(-3, 4, 5), Point(-8, 7, 3)},
		{I4.Translate(5, -3, 2), Vector(-3, 4, 5), Vector(-3, 4, 5)},
		{I4.Scale(2, 3, 4), Point(-4, 6, 8), Point(-8, 18, 32)},
		{I4.Scale(2, 3, 4), Vector(-4, 6, 8), Vector(-8, 18, 32)},
		{I4.Scale(2, 3, 4).Inverse(), Vector(-4, 6, 8), Vector(-2, 2, 2)},
		{I4.Scale(-1, 1, 1), Point(2, 3, 4), Point(-2, 3, 4)},
		{I4.RotateX(math.Pi / 4.0), Point(0, 1, 0), Point(0, math.Sqrt(2)/2.0, math.Sqrt(2)/2.0)},
		{I4.RotateX(math.Pi / 2.0), Point(0, 1, 0), Point(0, 0, 1)},
	}
	for _, row := range data {
		result := row.transform.TimesTuple(row.pt)
		if !result.Equals(row.expect) {
			t.Errorf("%v should have transformed to %v but was %v", row.pt, row.expect, result)
		}
	}
}
