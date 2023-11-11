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
		{I4.RotateX(math.Pi / 4.), Point(0, 1, 0), Point(0, math.Sqrt(2)/2., math.Sqrt(2)/2.)},
		{I4.RotateX(math.Pi / 2.), Point(0, 1, 0), Point(0, 0, 1)},
		{I4.RotateX(math.Pi / 4.).Inverse(), Point(0, 1, 0), Point(0, math.Sqrt(2)/2., -math.Sqrt(2)/2.)},
		{I4.RotateY(math.Pi / 4.), Point(0, 0, 1), Point(math.Sqrt(2)/2., 0, math.Sqrt(2)/2.)},
		{I4.RotateY(math.Pi / 2.), Point(0, 0, 1), Point(1, 0, 0)},
		{I4.RotateZ(math.Pi / 4.), Point(0, 1, 0), Point(-math.Sqrt(2)/2., math.Sqrt(2)/2., 0)},
		{I4.RotateZ(math.Pi / 2.), Point(0, 1, 0), Point(-1, 0, 0)},
		{I4.Shear(1, 0, 0, 0, 0, 0), Point(2, 3, 4), Point(5, 3, 4)},
		{I4.Shear(0, 1, 0, 0, 0, 0), Point(2, 3, 4), Point(6, 3, 4)},
		{I4.Shear(0, 0, 1, 0, 0, 0), Point(2, 3, 4), Point(2, 5, 4)},
		{I4.Shear(0, 0, 0, 1, 0, 0), Point(2, 3, 4), Point(2, 7, 4)},
		{I4.Shear(0, 0, 0, 0, 1, 0), Point(2, 3, 4), Point(2, 3, 6)},
		{I4.Shear(0, 0, 0, 0, 0, 1), Point(2, 3, 4), Point(2, 3, 7)},
		{I4.RotateX(math.Pi / 2), Point(1, 0, 1), Point(1, -1, 0)},
		{I4.Scale(5, 5, 5), Point(1, -1, 0), Point(5, -5, 0)},
		{I4.Translate(10, 5, 7), Point(5, -5, 0), Point(15, 0, 7)},
		{I4.RotateX(math.Pi/2).Scale(5, 5, 5).Translate(10, 5, 7), Point(1, 0, 1), Point(15, 0, 7)},
		{I4.Translate(10, 5, 7).Times(I4.Scale(5, 5, 5)).Times(I4.RotateX(math.Pi / 2)), Point(1, 0, 1), Point(15, 0, 7)},
	}
	for _, row := range data {
		result := row.transform.TimesTuple(row.pt)
		if !result.Equals(row.expect) {
			t.Errorf("%v should have transformed to %v but was %v", row.pt, row.expect, result)
		}
	}
}
